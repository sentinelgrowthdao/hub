package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func (k *Keeper) SetAllocation(ctx sdk.Context, alloc v2.Allocation) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(alloc.ID, alloc.GetAddress())
		value = k.cdc.MustMarshal(&alloc)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) (alloc v2.Allocation, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
		value = store.Get(key)
	)

	if value == nil {
		return alloc, false
	}

	k.cdc.MustUnmarshal(value, &alloc)
	return alloc, true
}

func (k *Keeper) HasAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress) {
	var (
		store = k.Store(ctx)
		key   = types.AllocationKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) GetAllocationsForSubscription(ctx sdk.Context, id uint64) (items v2.Allocations) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.GetAllocationForSubscriptionKeyPrefix(id))
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item v2.Allocation
		k.cdc.MustUnmarshal(iter.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateAllocationsForSubscription(ctx sdk.Context, id uint64, fn func(index int, item v2.Allocation) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.GetAllocationForSubscriptionKeyPrefix(id))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var alloc v2.Allocation
		k.cdc.MustUnmarshal(iter.Value(), &alloc)

		if stop := fn(i, alloc); stop {
			break
		}
		i++
	}
}
