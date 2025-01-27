package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) HandleMsgEndLease(ctx sdk.Context, msg *v1.MsgEndLeaseRequest) (*v1.MsgEndLeaseResponse, error) {
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

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return nil, err
	}

	amount := lease.RefundAmount()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&v1.EventRefund{
			ID:          lease.ID,
			ProvAddress: lease.ProvAddress,
			Amount:      amount.String(),
		},
	)

	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return nil, err
	}

	k.DeleteLease(ctx, lease.ID)
	k.DeleteLeaseForNode(ctx, nodeAddr, lease.ID)
	k.DeleteLeaseForProvider(ctx, provAddr, lease.ID)
	k.DeleteLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)
	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	ctx.EventManager().EmitTypedEvent(
		&v1.EventEnd{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
		},
	)

	return &v1.MsgEndLeaseResponse{}, nil
}

func (k *Keeper) HandleMsgRenewLease(ctx sdk.Context, msg *v1.MsgRenewLeaseRequest) (*v1.MsgRenewLeaseResponse, error) {
	if !k.IsValidLeaseHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	provAddr, err := base.ProvAddressFromBech32(lease.ProvAddress)
	if err != nil {
		return nil, err
	}

	provider, found := k.provider.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}
	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	nodeAddr, err := base.NodeAddressFromBech32(lease.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.node.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	amount := lease.RefundAmount()
	if err := k.SubtractDeposit(ctx, provAddr.Bytes(), amount); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&v1.EventRefund{
			ID:          lease.ID,
			ProvAddress: lease.ProvAddress,
			Amount:      amount.String(),
		},
	)

	k.DeleteLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	k.DeleteLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

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
		Renewable:  lease.Renewable,
		InactiveAt: ctx.BlockTime().Add(msg.GetHours()),
		PayoutAt:   ctx.BlockTime(),
	}

	if err := k.AddDeposit(ctx, provAddr.Bytes(), lease.Deposit); err != nil {
		return nil, err
	}

	k.SetLease(ctx, lease)
	k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	ctx.EventManager().EmitTypedEvent(
		&v1.EventRenew{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
			MaxHours:    lease.MaxHours,
			Price:       lease.Price.String(),
			Deposit:     lease.Deposit.String(),
		},
	)

	return &v1.MsgRenewLeaseResponse{}, nil
}

func (k *Keeper) HandleMsgStartLease(ctx sdk.Context, msg *v1.MsgStartLeaseRequest) (*v1.MsgStartLeaseResponse, error) {
	if !k.IsValidLeaseHours(ctx, msg.Hours) {
		return nil, types.NewErrorInvalidHours(msg.Hours)
	}

	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.provider.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}
	if !provider.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidProviderStatus(provAddr, provider.Status)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.node.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	if _, found := k.GetLatestLeaseForProviderByNode(ctx, provAddr, nodeAddr); found {
		return nil, types.NewErrorDuplicateLease(provAddr, nodeAddr)
	}

	count := k.GetCount(ctx)
	lease := v1.Lease{
		ID:          count + 1,
		ProvAddress: provAddr.String(),
		NodeAddress: nodeAddr.String(),
		Price:       price,
		Deposit: sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Hours),
		),
		Hours:      0,
		MaxHours:   msg.Hours,
		Renewable:  msg.Renewable,
		InactiveAt: ctx.BlockTime().Add(msg.GetHours()),
		PayoutAt:   ctx.BlockTime(),
	}

	if err := k.AddDeposit(ctx, provAddr.Bytes(), lease.Deposit); err != nil {
		return nil, err
	}

	k.SetCount(ctx, count+1)
	k.SetLease(ctx, lease)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)
	k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	ctx.EventManager().EmitTypedEvent(
		&v1.EventCreate{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
			MaxHours:    lease.MaxHours,
			Price:       lease.Price.String(),
			Deposit:     lease.Deposit.String(),
		},
	)

	return &v1.MsgStartLeaseResponse{
		ID: lease.ID,
	}, nil
}

func (k *Keeper) HandleMsgUpdateLease(ctx sdk.Context, msg *v1.MsgUpdateLeaseRequest) (*v1.MsgUpdateLeaseResponse, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}
	if msg.From != lease.ProvAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.DeleteLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	lease.Renewable = msg.Renewable

	k.SetLease(ctx, lease)
	k.SetLeaseForRenewalAt(ctx, lease.RenewalAt(), lease.ID)

	ctx.EventManager().EmitTypedEvent(
		&v1.EventUpdate{
			ID:          lease.ID,
			NodeAddress: lease.NodeAddress,
			ProvAddress: lease.ProvAddress,
			Renewable:   lease.Renewable,
		},
	)

	return &v1.MsgUpdateLeaseResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetParams(ctx, msg.Params)
	return &v1.MsgUpdateParamsResponse{}, nil
}
