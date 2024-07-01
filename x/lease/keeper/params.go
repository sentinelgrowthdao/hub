package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) MaxLeaseHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v1.KeyMaxLeaseHours, &v)
	return
}

func (k *Keeper) MinLeaseHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v1.KeyMinLeaseHours, &v)
	return
}

func (k *Keeper) StakingShare(ctx sdk.Context) (v sdkmath.LegacyDec) {
	k.params.Get(ctx, v1.KeyStakingShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params v1.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) v1.Params {
	return v1.NewParams(
		k.MaxLeaseHours(ctx),
		k.MinLeaseHours(ctx),
		k.StakingShare(ctx),
	)
}

func (k *Keeper) IsValidLeaseHours(ctx sdk.Context, hours int64) bool {
	if hours < k.MinLeaseHours(ctx) {
		return false
	}
	if hours > k.MaxLeaseHours(ctx) {
		return false
	}

	return true
}
