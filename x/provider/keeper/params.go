package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

// SetParams stores the parameters for the module in the KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v2.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v2.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// Deposit returns the deposit parameter from the module's parameters.
func (k *Keeper) Deposit(ctx sdk.Context) sdk.Coin {
	return k.GetParams(ctx).Deposit
}

// StakingShare returns the staking share parameter from the module's parameters.
func (k *Keeper) StakingShare(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).StakingShare
}
