package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v3.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Nodes {
		k.SetNode(ctx, item)
		if item.Status.Equal(v1base.StatusActive) {
			addr := item.GetAddress()
			k.SetNodeForInactiveAt(ctx, item.InactiveAt, addr)
		}
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v3.GenesisState {
	return v3.NewGenesisState(
		k.GetNodes(ctx),
		k.GetLeases(ctx),
		k.GetParams(ctx),
	)
}
