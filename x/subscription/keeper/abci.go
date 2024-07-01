package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	statusChangeDelay := k.StatusChangeDelay(ctx)

	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		k.DeleteSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

		if item.Status.Equal(v1base.StatusActive) {
			if err := k.SubscriptionInactivePendingPreHook(ctx, item.ID); err != nil {
				panic(err)
			}

			item.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)
			item.Status = v1base.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			k.SetSubscription(ctx, item)
			k.SetSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

			ctx.EventManager().EmitTypedEvent(
				&v2.EventUpdateStatus{
					Status:  v1base.StatusInactivePending,
					Address: item.AccAddress,
					ID:      item.ID,
					PlanID:  0,
				},
			)

			return false
		}

		k.DeleteSubscriptionForPlan(ctx, item.PlanID, item.ID)
		k.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
			accAddr := item.GetAddress()
			k.DeleteAllocation(ctx, item.ID, accAddr)
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.ID)

			return false
		})

		k.DeleteSubscription(ctx, item.ID)
		ctx.EventManager().EmitTypedEvent(
			&v2.EventUpdateStatus{
				Status:  v1base.StatusInactive,
				Address: item.AccAddress,
				ID:      item.ID,
				PlanID:  0,
			},
		)

		return false
	})
}
