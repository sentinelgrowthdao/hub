package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) HandleMsgStart(ctx sdk.Context, msg *v1.MsgStartRequest) (*v1.MsgStartResponse, error) {
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
	lease = v1.Lease{
		ID:          count + 1,
		ProvAddress: provAddr.String(),
		NodeAddress: nodeAddr.String(),
		Price:       price,
		Deposit: sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Hours),
		),
		Hours:    0,
		MaxHours: msg.Hours,
		PayoutAt: ctx.BlockTime(),
	}

	duration := time.Duration(msg.Hours) * time.Hour
	if msg.Renewable {
		lease.InactiveAt = time.Time{}
		lease.RenewalAt = ctx.BlockTime().Add(duration)
	} else {
		lease.InactiveAt = ctx.BlockTime().Add(duration)
		lease.RenewalAt = time.Time{}
	}

	if err := k.AddDeposit(ctx, provAddr.Bytes(), lease.Deposit); err != nil {
		return nil, err
	}

	k.SetLeaseCount(ctx, count+1)
	k.SetLease(ctx, lease)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)

	if msg.Renewable {
		k.SetLeaseForRenewalAt(ctx, lease.RenewalAt, lease.ID)
	} else {
		k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	}

	return &v1.MsgStartResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateDetails(ctx sdk.Context, msg *v1.MsgUpdateDetailsRequest) (*v1.MsgUpdateDetailsResponse, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if msg.Renewable {
		k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	} else {
		k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt, lease.ID)
	}

	if msg.Renewable {
		if lease.RenewalAt.IsZero() {
			lease.InactiveAt, lease.RenewalAt = time.Time{}, lease.InactiveAt
		}
	} else {
		if lease.InactiveAt.IsZero() {
			lease.InactiveAt, lease.RenewalAt = lease.RenewalAt, time.Time{}
		}
	}

	if msg.Renewable {
		k.SetLeaseForRenewalAt(ctx, lease.RenewalAt, lease.ID)
	} else {
		k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	}

	return &v1.MsgUpdateDetailsResponse{}, nil
}

func (k *Keeper) HandleMsgRenewLease(ctx sdk.Context, msg *v1.MsgRenewRequest) (*v1.MsgRenewResponse, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	var (
		nodeAddr = lease.GetNodeAddress()
		provAddr = lease.GetProvAddress()
	)

	refund := lease.Refund()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), refund); err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	duration := time.Duration(msg.Hours) * time.Hour
	lease = v1.Lease{
		ID:          lease.ID,
		ProvAddress: lease.ProvAddress,
		NodeAddress: lease.NodeAddress,
		Price:       price,
		Deposit: sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Hours),
		),
		Hours:      0,
		MaxHours:   msg.Hours,
		InactiveAt: time.Time{},
		PayoutAt:   ctx.BlockTime(),
		RenewalAt:  ctx.BlockTime().Add(duration),
	}

	if err := k.AddDeposit(ctx, provAddr.Bytes(), lease.Deposit); err != nil {
		return nil, err
	}

	k.SetLease(ctx, lease)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt, lease.ID)

	return &v1.MsgRenewResponse{}, nil
}

func (k *Keeper) HandleMsgEnd(ctx sdk.Context, msg *v1.MsgEndRequest) (*v1.MsgEndResponse, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
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
	k.DeleteLeaseForNode(ctx, nodeAddr, lease.ID)
	k.DeleteLeaseForProvider(ctx, provAddr, lease.ID)
	k.DeleteLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)
	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt, lease.ID)

	return &v1.MsgEndResponse{}, nil
}
