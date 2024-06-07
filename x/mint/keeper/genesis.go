package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/mint/types"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	for _, item := range state.Inflations {
		k.SetInflation(ctx, item)
	}
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetInflations(ctx),
	)
}
