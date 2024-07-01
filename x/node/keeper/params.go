package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) Deposit(ctx sdk.Context) (v sdk.Coin) {
	k.params.Get(ctx, v3.KeyDeposit, &v)
	return
}

func (k *Keeper) ActiveDuration(ctx sdk.Context) (v time.Duration) {
	k.params.Get(ctx, v3.KeyActiveDuration, &v)
	return
}

func (k *Keeper) MinGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, v3.KeyMinGigabytePrices, &v)
	return
}

func (k *Keeper) MinHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, v3.KeyMinHourlyPrices, &v)
	return
}

func (k *Keeper) MaxSessionGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v3.KeyMaxSessionGigabytes, &v)
	return
}

func (k *Keeper) MinSessionGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v3.KeyMinSessionGigabytes, &v)
	return
}

func (k *Keeper) MaxSessionHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v3.KeyMaxSessionHours, &v)
	return
}

func (k *Keeper) MinSessionHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v3.KeyMinSessionHours, &v)
	return
}

func (k *Keeper) StakingShare(ctx sdk.Context) (v sdkmath.LegacyDec) {
	k.params.Get(ctx, v3.KeyStakingShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params v3.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) v3.Params {
	return v3.NewParams(
		k.Deposit(ctx),
		k.ActiveDuration(ctx),
		k.MinGigabytePrices(ctx),
		k.MinHourlyPrices(ctx),
		k.MaxSessionGigabytes(ctx),
		k.MinSessionGigabytes(ctx),
		k.MaxSessionHours(ctx),
		k.MinSessionHours(ctx),
		k.StakingShare(ctx),
	)
}

func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.Coins) bool {
	minPrices := k.MinGigabytePrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidHourlyPrices(ctx sdk.Context, prices sdk.Coins) bool {
	minPrices := k.MinHourlyPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidSessionGigabytes(ctx sdk.Context, gigabytes int64) bool {
	if gigabytes < k.MinSessionGigabytes(ctx) {
		return false
	}
	if gigabytes > k.MaxSessionGigabytes(ctx) {
		return false
	}

	return true
}

func (k *Keeper) IsValidSessionHours(ctx sdk.Context, hours int64) bool {
	if hours < k.MinSessionHours(ctx) {
		return false
	}
	if hours > k.MaxSessionHours(ctx) {
		return false
	}

	return true
}

func (k *Keeper) IsMinGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, v3.KeyMinGigabytePrices)
}

func (k *Keeper) IsMinHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, v3.KeyMinHourlyPrices)
}
