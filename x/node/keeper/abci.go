package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func (k *Keeper) handleInactiveNodes(ctx sdk.Context) {
	k.IterateNodesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v2.Node) bool {
		nodeAddr, err := base.NodeAddressFromBech32(item.Address)
		if err != nil {
			panic(err)
		}

		k.DeleteNodeForInactiveAt(ctx, item.InactiveAt, nodeAddr)

		if err := k.NodeInactivePreHook(ctx, nodeAddr); err != nil {
			panic(err)
		}

		k.DeleteActiveNode(ctx, nodeAddr)

		item.InactiveAt = time.Time{}
		item.Status = v1base.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetNode(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&v2.EventUpdateStatus{
				Status:  v1base.StatusInactive,
				Address: item.Address,
			},
		)

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactiveNodes(ctx)
}
