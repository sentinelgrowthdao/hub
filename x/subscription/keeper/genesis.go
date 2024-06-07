package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	var (
		subscriptions = k.GetSubscriptions(ctx)
		items         = make(types.GenesisSubscriptions, 0, len(subscriptions))
	)

	return types.NewGenesisState(items, k.GetParams(ctx))
}
