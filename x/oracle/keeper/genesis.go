package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v1.GenesisState) {
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v1.GenesisState {
	return &v1.GenesisState{}
}
