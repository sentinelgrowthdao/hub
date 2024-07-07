package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	return nil
}
