package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	return nil
}

func (k *Keeper) SubscriptionInactivePendingPreHook(ctx sdk.Context, id uint64) error {
	statusChangeDelay := k.StatusChangeDelay(ctx)
	k.IterateSessionsForSubscription(ctx, id, func(_ int, item v3.Session) (stop bool) {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		item.SetStatus(v1base.StatusInactivePending)
		item.SetInactiveAt(ctx.BlockTime().Add(statusChangeDelay))
		item.SetStatusAt(ctx.BlockTime())

		k.SetSession(ctx, item)
		k.SetSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		return false
	})

	return nil
}
