package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/types"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Nodes {
		k.SetNode(ctx, item)
		if item.Status.Equal(base.StatusActive) {
			addr := item.GetAddress()
			k.SetNodeForInactiveAt(ctx, item.InactiveAt, addr)
		}
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetNodes(ctx),
		k.GetParams(ctx),
	)
}
