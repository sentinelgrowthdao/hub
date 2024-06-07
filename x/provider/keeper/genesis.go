package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Providers {
		k.SetProvider(ctx, item)
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetProviders(ctx),
		k.GetParams(ctx),
	)
}
