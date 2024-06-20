package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/mint/types"
	"github.com/sentinel-official/hub/v12/x/mint/types/v1"
)

func (k *Keeper) SetInflation(ctx sdk.Context, inflation v1.Inflation) {
	var (
		store = k.Store(ctx)
		key   = types.InflationKey(inflation.Timestamp)
		value = k.cdc.MustMarshal(&inflation)
	)

	store.Set(key, value)
}

func (k *Keeper) GetInflation(ctx sdk.Context, t time.Time) (inflation v1.Inflation, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InflationKey(t)
		value = store.Get(key)
	)

	if value == nil {
		return inflation, false
	}

	k.cdc.MustUnmarshal(value, &inflation)
	return inflation, true
}

func (k *Keeper) DeleteInflation(ctx sdk.Context, t time.Time) {
	var (
		store = k.Store(ctx)
		key   = types.InflationKey(t)
	)

	store.Delete(key)
}

func (k *Keeper) GetInflations(ctx sdk.Context) (items []v1.Inflation) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.InflationKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v1.Inflation
		k.cdc.MustUnmarshal(iterator.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateInflations(ctx sdk.Context, fn func(index int, item v1.Inflation) (stop bool)) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.InflationKeyPrefix)
	)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v1.Inflation
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}
