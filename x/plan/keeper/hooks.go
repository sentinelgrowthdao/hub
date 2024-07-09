package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	return nil
}
