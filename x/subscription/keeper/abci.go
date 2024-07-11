package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) handleInactivePendingSubscriptions(ctx sdk.Context) {
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		if !item.Status.Equal(v1base.StatusActive) {
			return false
		}

		msg := item.MsgCancelRequest()
		if _, err := k.HandleMsgCancel(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})
}

func (k *Keeper) handleInactiveSubscriptions(ctx sdk.Context) {
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		if !item.Status.Equal(v1base.StatusInactivePending) {
			return false
		}

		k.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
			accAddr, err := sdk.AccAddressFromBech32(item.Address)
			if err != nil {
				panic(err)
			}

			k.DeleteAllocation(ctx, item.ID, accAddr)
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.ID)

			return false
		})

		k.DeleteSubscription(ctx, item.ID)
		k.DeleteSubscriptionForPlan(ctx, item.PlanID, item.ID)
		k.DeleteSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

		ctx.EventManager().EmitTypedEvent(
			&v3.EventUpdate{
				ID:         item.ID,
				PlanID:     item.PlanID,
				AccAddress: item.AccAddress,
				Status:     v1base.StatusInactive,
			},
		)

		return false
	})
}

func (k *Keeper) handleSubscriptionRenewals(ctx sdk.Context) {
	k.IterateSubscriptionsForRenewalAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		msg := item.MsgRenewRequest()
		if _, err := k.HandleMsgRenew(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactivePendingSubscriptions(ctx)
	k.handleInactiveSubscriptions(ctx)
	k.handleSubscriptionRenewals(ctx)
}
