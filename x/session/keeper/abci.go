package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) handleInactiveSessions(ctx sdk.Context) {
	statusChangeDelay := k.StatusChangeDelay(ctx)

	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		if item.GetStatus().Equal(v1base.StatusActive) {
			k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

			item.SetStatus(v1base.StatusInactivePending)
			item.SetInactiveAt(ctx.BlockTime().Add(statusChangeDelay))
			item.SetStatusAt(ctx.BlockTime())

			k.SetSession(ctx, item)
			k.SetSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

			return false
		}

		if err := k.SessionInactivePreHook(ctx, item.GetID()); err != nil {
			panic(err)
		}

		var (
			accAddr  = item.GetAccAddress()
			nodeAddr = item.GetNodeAddress()
		)

		k.DeleteSession(ctx, item.GetID())
		k.DeleteSessionForAccount(ctx, accAddr, item.GetID())
		k.DeleteSessionForNode(ctx, nodeAddr, item.GetID())
		k.DeleteSessionForSubscription(ctx, 0, item.GetID())
		k.DeleteSessionForAllocation(ctx, 0, accAddr, item.GetID())
		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactiveSessions(ctx)
}
