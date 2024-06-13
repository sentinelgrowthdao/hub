package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/vpn/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
	k.Deposit.InitGenesis(ctx, state.Deposits)
	k.Node.InitGenesis(ctx, state.Nodes)
	k.Plan.InitGenesis(ctx, state.Plans)
	k.Provider.InitGenesis(ctx, state.Providers)
	k.Session.InitGenesis(ctx, state.Sessions)
	k.Subscription.InitGenesis(ctx, state.Subscriptions)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return &v1.GenesisState{
		Deposits:      k.Deposit.ExportGenesis(ctx),
		Nodes:         k.Node.ExportGenesis(ctx),
		Plans:         k.Plan.ExportGenesis(ctx),
		Providers:     k.Provider.ExportGenesis(ctx),
		Sessions:      k.Session.ExportGenesis(ctx),
		Subscriptions: k.Subscription.ExportGenesis(ctx),
	}
}
