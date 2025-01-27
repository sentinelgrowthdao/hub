package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) handleInactivePendingSessions(ctx sdk.Context) {
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		if !item.GetStatus().Equal(v1base.StatusActive) {
			return false
		}

		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		msg := &v3.MsgCancelSessionRequest{
			From: item.GetAccAddress(),
			ID:   item.GetID(),
		}

		handler := k.router.Handler(msg)
		if _, err := handler(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})
}

func (k *Keeper) handleInactiveSessions(ctx sdk.Context) {
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Session) bool {
		if !item.GetStatus().Equal(v1base.StatusInactivePending) {
			return false
		}

		accAddr, err := sdk.AccAddressFromBech32(item.GetAccAddress())
		if err != nil {
			panic(err)
		}

		nodeAddr, err := base.NodeAddressFromBech32(item.GetNodeAddress())
		if err != nil {
			panic(err)
		}

		k.DeleteSessionForInactiveAt(ctx, item.GetInactiveAt(), item.GetID())

		if err := k.SessionInactivePreHook(ctx, item.GetID()); err != nil {
			panic(err)
		}

		k.DeleteSession(ctx, item.GetID())
		k.DeleteSessionForAccount(ctx, accAddr, item.GetID())
		k.DeleteSessionForAllocation(ctx, 0, accAddr, item.GetID())
		k.DeleteSessionForNode(ctx, nodeAddr, item.GetID())
		k.DeleteSessionForSubscription(ctx, 0, item.GetID())

		ctx.EventManager().EmitTypedEvent(
			&v3.EventUpdateStatus{
				ID:          item.GetID(),
				AccAddress:  item.GetAccAddress(),
				NodeAddress: item.GetNodeAddress(),
				Status:      v1base.StatusInactive,
				StatusAt:    ctx.BlockTime().String(),
			},
		)

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactivePendingSessions(ctx)
	k.handleInactiveSessions(ctx)
}
