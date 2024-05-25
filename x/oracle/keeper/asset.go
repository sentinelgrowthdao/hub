package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

func (k *Keeper) SetAsset(ctx sdk.Context, asset types.Asset) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(asset.Denom)
		value = k.cdc.MustMarshal(&asset)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAsset(ctx sdk.Context, denom string) (v types.Asset, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(denom)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteAsset(ctx sdk.Context, denom string) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(denom)
	)

	store.Delete(key)
}

func (k *Keeper) SetAssetPrice(ctx sdk.Context, price sdk.Coin) {
	var (
		store = k.Store(ctx)
		key   = types.AssetPriceKey(price.Denom)
		value = k.cdc.MustMarshal(&price)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAssetPrice(ctx sdk.Context, denom string) (v sdk.Coin, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AssetPriceKey(denom)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteAssetPrice(ctx sdk.Context, denom string) {
	var (
		store = k.Store(ctx)
		key   = types.AssetPriceKey(denom)
	)

	store.Delete(key)
}

func (k *Keeper) SetDenomForRequest(ctx sdk.Context, sequence uint64, index int, denom string) {
	var (
		store = k.Store(ctx)
		key   = types.DenomForRequestKey(sequence, index)
		value = k.cdc.MustMarshal(&protobuf.StringValue{Value: denom})
	)

	store.Set(key, value)
}

func (k *Keeper) GetDenomForRequest(ctx sdk.Context, sequence uint64, index int) (v string, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.DenomForRequestKey(sequence, index)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	var denom protobuf.StringValue
	k.cdc.MustUnmarshal(value, &denom)

	return denom.GetValue(), true
}

func (k *Keeper) DeleteDenomForRequest(ctx sdk.Context, sequence uint64, index int) {
	var (
		store = k.Store(ctx)
		key   = types.DenomForRequestKey(sequence, index)
	)

	store.Delete(key)
}
