package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func (k *Keeper) Deposit(ctx sdk.Context) (v sdk.Coin) {
	k.params.Get(ctx, v2.KeyDeposit, &v)
	return
}

func (k *Keeper) ActiveDuration(ctx sdk.Context) (v time.Duration) {
	k.params.Get(ctx, v2.KeyActiveDuration, &v)
	return
}

func (k *Keeper) MaxGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, v2.KeyMaxGigabytePrices, &v)
	return
}

func (k *Keeper) MinGigabytePrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, v2.KeyMinGigabytePrices, &v)
	return
}

func (k *Keeper) MaxHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, v2.KeyMaxHourlyPrices, &v)
	return
}

func (k *Keeper) MinHourlyPrices(ctx sdk.Context) (v sdk.Coins) {
	k.params.Get(ctx, v2.KeyMinHourlyPrices, &v)
	return
}

func (k *Keeper) MaxSubscriptionGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v2.KeyMaxSubscriptionGigabytes, &v)
	return
}

func (k *Keeper) MinSubscriptionGigabytes(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v2.KeyMinSubscriptionGigabytes, &v)
	return
}

func (k *Keeper) MaxSubscriptionHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v2.KeyMaxSubscriptionHours, &v)
	return
}

func (k *Keeper) MinSubscriptionHours(ctx sdk.Context) (v int64) {
	k.params.Get(ctx, v2.KeyMinSubscriptionHours, &v)
	return
}

func (k *Keeper) StakingShare(ctx sdk.Context) (v sdkmath.LegacyDec) {
	k.params.Get(ctx, v2.KeyStakingShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params v2.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) v2.Params {
	return v2.NewParams(
		k.Deposit(ctx),
		k.ActiveDuration(ctx),
		k.MaxGigabytePrices(ctx),
		k.MinGigabytePrices(ctx),
		k.MaxHourlyPrices(ctx),
		k.MinHourlyPrices(ctx),
		k.MaxSubscriptionGigabytes(ctx),
		k.MinSubscriptionGigabytes(ctx),
		k.MaxSubscriptionHours(ctx),
		k.MinSubscriptionHours(ctx),
		k.StakingShare(ctx),
	)
}

func (k *Keeper) IsValidGigabytePrices(ctx sdk.Context, prices sdk.Coins) bool {
	maxPrices := k.MaxGigabytePrices(ctx)
	for _, coin := range maxPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

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
	maxPrices := k.MaxHourlyPrices(ctx)
	for _, coin := range maxPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.GT(coin.Amount) {
			return false
		}
	}

	minPrices := k.MinHourlyPrices(ctx)
	for _, coin := range minPrices {
		amount := prices.AmountOf(coin.Denom)
		if amount.LT(coin.Amount) {
			return false
		}
	}

	return true
}

func (k *Keeper) IsValidSubscriptionGigabytes(ctx sdk.Context, gigabytes int64) bool {
	if gigabytes < k.MinSubscriptionGigabytes(ctx) {
		return false
	}
	if gigabytes > k.MaxSubscriptionGigabytes(ctx) {
		return false
	}

	return true
}

func (k *Keeper) IsValidSubscriptionHours(ctx sdk.Context, hours int64) bool {
	if hours < k.MinSubscriptionHours(ctx) {
		return false
	}
	if hours > k.MaxSubscriptionHours(ctx) {
		return false
	}

	return true
}

func (k *Keeper) IsMaxGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, v2.KeyMaxGigabytePrices)
}

func (k *Keeper) IsMinGigabytePricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, v2.KeyMinGigabytePrices)
}

func (k *Keeper) IsMaxHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, v2.KeyMaxHourlyPrices)
}

func (k *Keeper) IsMinHourlyPricesModified(ctx sdk.Context) bool {
	return k.params.Modified(ctx, v2.KeyMinHourlyPrices)
}
