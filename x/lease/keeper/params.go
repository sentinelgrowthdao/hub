package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

// SetParams stores the lease module parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the lease module parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// MaxLeaseHours retrieves the maximum lease hours parameter from the module's parameters.
func (k *Keeper) MaxLeaseHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MaxLeaseHours
}

// MinLeaseHours retrieves the minimum lease hours parameter from the module's parameters.
func (k *Keeper) MinLeaseHours(ctx sdk.Context) int64 {
	return k.GetParams(ctx).MinLeaseHours
}

// StakingShare retrieves the staking share parameter from the module's parameters.
func (k *Keeper) StakingShare(ctx sdk.Context) sdkmath.LegacyDec {
	return k.GetParams(ctx).StakingShare
}

// IsValidLeaseHours checks if the provided lease hours are within the valid range defined by the module's parameters.
func (k *Keeper) IsValidLeaseHours(ctx sdk.Context, hours int64) bool {
	if hours < k.MinLeaseHours(ctx) {
		return false
	}
	if hours > k.MaxLeaseHours(ctx) {
		return false
	}

	return true
}
