package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v2.GenesisState) {
	k.SetParams(ctx, state.Params)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v2.GenesisState {
	var (
		subscriptions = k.GetSubscriptions(ctx)
		items         = make(v2.GenesisSubscriptions, 0, len(subscriptions))
	)

	return v2.NewGenesisState(items, k.GetParams(ctx))
}
