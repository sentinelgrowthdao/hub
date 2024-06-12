package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v2.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, item := range state.Providers {
		k.SetProvider(ctx, item)
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v2.GenesisState {
	return v2.NewGenesisState(
		k.GetProviders(ctx),
		k.GetParams(ctx),
	)
}
