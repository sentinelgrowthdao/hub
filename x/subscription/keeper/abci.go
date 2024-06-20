package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

// EndBlock is a function that gets called at the end of every block.
// It processes the subscriptions that have become inactive and performs the necessary actions accordingly.
func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	// Get the status change delay from the store.
	statusChangeDelay := k.StatusChangeDelay(ctx)

	// Iterate over all subscriptions that have become inactive at the current block time.
	k.IterateSubscriptionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Subscription) bool {
		k.Logger(ctx).Info("Found an inactive subscription", "id", item.ID, "status", item.Status)

		// Delete the subscription from the InactiveAt index before updating the InactiveAt value.
		k.DeleteSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

		// If the subscription status is 'Active', update its InactiveAt value and set it to 'InactivePending'.
		if item.Status.Equal(v1base.StatusActive) {
			// Run the SubscriptionInactivePendingHook to perform custom actions before setting the subscription to inactive pending state.
			if err := k.SubscriptionInactivePendingHook(ctx, item.ID); err != nil {
				panic(err)
			}

			item.Status = v1base.StatusInactivePending
			item.StatusAt = ctx.BlockTime()
			item.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)

			// Save the updated subscription to the store and update the InactiveAt index.
			k.SetSubscription(ctx, item)
			k.SetSubscriptionForInactiveAt(ctx, item.InactiveAt, item.ID)

			// Emit an event to notify that the subscription status has been updated.
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

		// For plan-level subscriptions, delete the subscription from the PlanID index.
		k.DeleteSubscriptionForPlan(ctx, item.PlanID, item.ID)

		// Iterate over all allocations associated with the plan-level subscription and delete them from the store.
		k.IterateAllocationsForSubscription(ctx, item.ID, func(_ int, item v2.Allocation) bool {
			accAddr := item.GetAddress()

			// Delete the allocation associated with the plan-level subscription.
			k.DeleteAllocation(ctx, item.ID, accAddr)

			// Delete the plan-level subscription from the Account index.
			k.DeleteSubscriptionForAccount(ctx, accAddr, item.ID)

			return false
		})

		// Finally, delete the subscription from the store and emit an event to notify its status change to 'Inactive'.
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

	// Return an empty ValidatorUpdate slice as no validator updates are needed for the end block.
	return nil
}
