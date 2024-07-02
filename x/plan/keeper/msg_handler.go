package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

func (k *Keeper) HandleMsgCreate(ctx sdk.Context, msg *v2.MsgCreateRequest) (*v2.MsgCreateResponse, error) {
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	if !k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	count := k.GetPlanCount(ctx)
	plan := v2.Plan{
		ID:              count + 1,
		ProviderAddress: provAddr.String(),
		Duration:        msg.Duration,
		Gigabytes:       msg.Gigabytes,
		Prices:          msg.Prices,
		Status:          v1base.StatusInactive,
		StatusAt:        ctx.BlockTime(),
	}

	k.SetPlanCount(ctx, count+1)
	k.SetPlan(ctx, plan)
	k.SetPlanForProvider(ctx, provAddr, plan.ID)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventCreate{
			Address: plan.ProviderAddress,
			ID:      plan.ID,
		},
	)

	return &v2.MsgCreateResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateStatus(ctx sdk.Context, msg *v2.MsgUpdateStatusRequest) (*v2.MsgUpdateStatusResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.ProviderAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if plan.Status.Equal(v1base.StatusActive) {
		if msg.Status.Equal(v1base.StatusInactive) {
			k.DeleteActivePlan(ctx, plan.ID)
		}
	}

	if plan.Status.Equal(v1base.StatusInactive) {
		if msg.Status.Equal(v1base.StatusActive) {
			k.DeleteInactivePlan(ctx, plan.ID)
		}
	}

	plan.Status = msg.Status
	plan.StatusAt = ctx.BlockTime()

	k.SetPlan(ctx, plan)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateStatus{
			Status:  plan.Status,
			Address: plan.ProviderAddress,
			ID:      plan.ID,
		},
	)

	return &v2.MsgUpdateStatusResponse{}, nil
}

func (k *Keeper) HandleMsgLinkNode(ctx sdk.Context, msg *v2.MsgLinkNodeRequest) (*v2.MsgLinkNodeResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.ProviderAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	if !k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	k.SetNodeForPlan(ctx, plan.ID, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventLinkNode{
			Address:     plan.ProviderAddress,
			NodeAddress: msg.NodeAddress,
			ID:          plan.ID,
		},
	)

	return &v2.MsgLinkNodeResponse{}, nil
}

func (k *Keeper) HandleMsgUnlinkNode(ctx sdk.Context, msg *v2.MsgUnlinkNodeRequest) (*v2.MsgUnlinkNodeResponse, error) {
	plan, found := k.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if msg.From != plan.ProviderAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	k.DeleteNodeForPlan(ctx, plan.ID, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUnlinkNode{
			Address:     plan.ProviderAddress,
			NodeAddress: msg.NodeAddress,
			ID:          plan.ID,
		},
	)

	return &v2.MsgUnlinkNodeResponse{}, nil
}

func (k *Keeper) HandleMsgSubscribe(ctx sdk.Context, msg *v2.MsgSubscribeRequest) (*v2.MsgSubscribeResponse, error) {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	subscription, err := k.CreateSubscriptionForPlan(ctx, accAddr, msg.ID, msg.Denom)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitTypedEvent(
		&v2.EventCreateSubscription{
			Address:         subscription.Address,
			ProviderAddress: "",
			ID:              subscription.ID,
			PlanID:          subscription.PlanID,
		},
	)

	return &v2.MsgSubscribeResponse{}, nil
}
