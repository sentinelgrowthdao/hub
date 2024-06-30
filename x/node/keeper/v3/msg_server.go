package v3

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

var (
	_ v3.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v3.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgStartLease(c context.Context, msg *v3.MsgStartLeaseRequest) (*v3.MsgStartLeaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if !k.IsValidLeaseHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	lease, found := k.GetLatestLeaseForProviderByNode(ctx, provAddr, nodeAddr)
	if found {
		return nil, types.NewErrorDuplicateLease(provAddr, nodeAddr)
	}

	if found := k.HasProvider(ctx, provAddr); !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	count := k.GetLeaseCount(ctx)
	lease = v3.Lease{
		ID:          count + 1,
		ProvAddress: provAddr.String(),
		NodeAddress: nodeAddr.String(),
		Price:       price,
		Deposit: sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Hours),
		),
		Hours:     0,
		MaxHours:  msg.Hours,
		CreatedAt: ctx.BlockTime(),
		PayoutAt:  ctx.BlockTime(),
	}

	if err := k.AddDeposit(ctx, provAddr.Bytes(), lease.Deposit); err != nil {
		return nil, err
	}

	duration := time.Duration(lease.MaxHours) * time.Hour
	if msg.Renewable {
		lease.InactiveAt = time.Time{}
		lease.RenewAt = lease.CreatedAt.Add(duration)
	} else {
		lease.RenewAt = time.Time{}
		lease.InactiveAt = lease.CreatedAt.Add(duration)
	}

	k.SetLeaseCount(ctx, count+1)
	k.SetLease(ctx, lease)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)

	if msg.Renewable {
		k.SetLeaseForRenewAt(ctx, lease.RenewAt, lease.ID)
	} else {
		k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	}

	return &v3.MsgStartLeaseResponse{}, nil
}

func (k *msgServer) MsgUpdateLeaseDetails(c context.Context, msg *v3.MsgUpdateLeaseDetailsRequest) (*v3.MsgUpdateLeaseDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorised(msg.From)
	}

	if msg.Renewable {
		k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	} else {
		k.DeleteLeaseForRenewAt(ctx, lease.RenewAt, lease.ID)
	}

	duration := time.Duration(lease.MaxHours) * time.Hour
	if msg.Renewable {
		lease.InactiveAt = time.Time{}
		lease.RenewAt = lease.CreatedAt.Add(duration)
	} else {
		lease.RenewAt = time.Time{}
		lease.InactiveAt = lease.CreatedAt.Add(duration)
	}

	if msg.Renewable {
		k.SetLeaseForRenewAt(ctx, lease.RenewAt, lease.ID)
	} else {
		k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	}

	return &v3.MsgUpdateLeaseDetailsResponse{}, nil
}

func (k *msgServer) MsgRenewLease(_ context.Context, _ *v3.MsgRenewLeaseRequest) (*v3.MsgRenewLeaseResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgEndLease(c context.Context, msg *v3.MsgEndLeaseRequest) (*v3.MsgEndLeaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorised(msg.From)
	}

	if err := k.LeaseInactivePreHook(ctx, lease.ID); err != nil {
		return nil, err
	}

	var (
		nodeAddr = lease.GetNodeAddress()
		provAddr = lease.GetProvAddress()
	)

	refund := lease.Refund()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), refund); err != nil {
		return nil, err
	}

	k.DeleteLease(ctx, lease.ID)
	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	k.DeleteLeaseForNode(ctx, nodeAddr, lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.DeleteLeaseForProvider(ctx, provAddr, lease.ID)
	k.DeleteLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)
	k.DeleteLeaseForRenewAt(ctx, lease.RenewAt, lease.ID)

	return &v3.MsgEndLeaseResponse{}, nil
}

func (k *msgServer) MsgStartSession(c context.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
