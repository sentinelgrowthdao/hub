package keeper

import (
	"time"

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

			k.DeleteSubscriptionForRenewalAt(ctx, item.RenewalAt, item.ID)

			item.Status = v1base.StatusInactivePending
			item.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)
			item.RenewalAt = time.Time{}
			item.StatusAt = ctx.BlockTime()

			k.SetSubscription(ctx, item)
			k.SetSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

			return false
		}

		k.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
			accAddr := item.GetAddress()
			k.DeleteAllocation(ctx, item.ID, accAddr)
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.ID)

			return false
		})

		k.DeleteSubscription(ctx, item.ID)
		k.DeleteSubscriptionForPlan(ctx, item.PlanID, item.ID)

		return false
	})
}
