package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

func (k *Keeper) GetPortID(ctx sdk.Context) string {
	store := k.Store(ctx)
	return string(store.Get(types.PortIDKey))
}
