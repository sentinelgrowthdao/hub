package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) HandleMsgAllocate(ctx sdk.Context, msg *v2.MsgAllocateRequest) (*v2.MsgAllocateResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	fromAlloc, found := k.GetAllocation(ctx, subscription.ID, fromAddr)
	if !found {
		return nil, types.NewErrorAllocationNotFound(subscription.ID, fromAddr)
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	toAlloc, found := k.GetAllocation(ctx, subscription.ID, toAddr)
	if !found {
		toAlloc = v2.Allocation{
			ID:            subscription.ID,
			Address:       toAddr.String(),
			GrantedBytes:  sdk.ZeroInt(),
			UtilisedBytes: sdk.ZeroInt(),
		}

		k.SetSubscriptionForAccount(ctx, toAddr, subscription.ID)
	}

	grantedBytes := fromAlloc.GrantedBytes.Add(toAlloc.GrantedBytes)
	utilisedBytes := fromAlloc.UtilisedBytes.Add(toAlloc.UtilisedBytes)
	availableBytes := grantedBytes.Sub(utilisedBytes)

	if msg.Bytes.GT(availableBytes) {
		return nil, types.NewErrorInsufficientBytes(subscription.ID, msg.Bytes)
	}

	fromAlloc.GrantedBytes = availableBytes.Sub(msg.Bytes)
	if fromAlloc.GrantedBytes.LT(fromAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.ID, fromAddr)
	}

	k.SetAllocation(ctx, fromAlloc)

	toAlloc.GrantedBytes = msg.Bytes
	if toAlloc.GrantedBytes.LT(toAlloc.UtilisedBytes) {
		return nil, types.NewErrorInvalidAllocation(subscription.ID, toAddr)
	}

	k.SetAllocation(ctx, toAlloc)

	return &v2.MsgAllocateResponse{}, nil
}

func (k *Keeper) HandleMsgCancel(ctx sdk.Context, msg *v2.MsgCancelRequest) (*v2.MsgCancelResponse, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}
	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if err := k.SubscriptionInactivePendingPreHook(ctx, subscription.ID); err != nil {
		return nil, err
	}

	k.DeleteSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt, subscription.ID)

	statusChangeDelay := k.StatusChangeDelay(ctx)
	subscription.Status = v1base.StatusInactivePending
	subscription.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)
	subscription.RenewalAt = time.Time{}
	subscription.StatusAt = ctx.BlockTime()

	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)

	return &v2.MsgCancelResponse{}, nil
}

func (k *Keeper) HandleMsgStart(ctx sdk.Context, msg *v3.MsgStartRequest) (*v3.MsgStartResponse, error) {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	plan, found := k.plan.GetPlan(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorPlanNotFound(msg.ID)
	}
	if !plan.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	// TODO: check duplicate account address for plan?

	price, found := plan.Price(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	var (
		stakingShare  = k.provider.StakingShare(ctx)
		stakingReward = baseutils.GetProportionOfCoin(price, stakingShare)
	)

	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
		return nil, err
	}

	var (
		provAddr = plan.GetProviderAddress()
		payment  = price.Sub(stakingReward)
	)

	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	count := k.GetCount(ctx)
	subscription := v3.Subscription{
		ID:         count + 1,
		AccAddress: accAddr.String(),
		PlanID:     plan.ID,
		Price:      price,
		Status:     v1base.StatusActive,
		InactiveAt: time.Time{},
		RenewalAt:  time.Time{},
		StatusAt:   ctx.BlockTime(),
	}

	if msg.Renewable {
		subscription.InactiveAt = time.Time{}
		subscription.RenewalAt = ctx.BlockTime().Add(plan.Duration)
	} else {
		subscription.InactiveAt = ctx.BlockTime().Add(plan.Duration)
		subscription.RenewalAt = time.Time{}
	}

	k.SetCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.ID)
	k.SetSubscriptionForPlan(ctx, subscription.PlanID, subscription.ID)

	if msg.Renewable {
		k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt, subscription.ID)
	} else {
		k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	}

	alloc := v2.Allocation{
		ID:            subscription.ID,
		Address:       subscription.AccAddress,
		GrantedBytes:  base.Gigabyte.MulRaw(plan.Gigabytes),
		UtilisedBytes: sdkmath.ZeroInt(),
	}

	k.SetAllocation(ctx, alloc)

	return &v3.MsgStartResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateDetails(ctx sdk.Context, msg *v3.MsgUpdateDetailsRequest) (*v3.MsgUpdateDetailsResponse, error) {
	subscription, found := k.GetSubscription(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSubscriptionNotFound(msg.ID)
	}
	if !subscription.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSubscriptionStatus(subscription.ID, subscription.Status)
	}
	if msg.From != subscription.AccAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if msg.Renewable {
		k.DeleteSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	} else {
		k.DeleteSubscriptionForRenewalAt(ctx, subscription.RenewalAt, subscription.ID)
	}

	if msg.Renewable {
		if subscription.RenewalAt.IsZero() {
			subscription.InactiveAt, subscription.RenewalAt = time.Time{}, subscription.InactiveAt
		}
	} else {
		if subscription.InactiveAt.IsZero() {
			subscription.InactiveAt, subscription.RenewalAt = subscription.RenewalAt, time.Time{}
		}
	}

	if msg.Renewable {
		k.SetSubscriptionForRenewalAt(ctx, subscription.RenewalAt, subscription.ID)
	} else {
		k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)
	}

	return &v3.MsgUpdateDetailsResponse{}, nil
}

func (k *Keeper) HandleMsgRenew(ctx sdk.Context, msg *v3.MsgRenewRequest) (*v3.MsgRenewResponse, error) {
	return &v3.MsgRenewResponse{}, nil
}

func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	return &v3.MsgStartSessionResponse{}, nil
}
