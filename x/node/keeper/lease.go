package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) SetLeaseCount(ctx sdk.Context, count uint64) {
	var (
		key   = types.LeaseCountKey
		value = k.cdc.MustMarshal(&protobuf.UInt64Value{Value: count})
		store = k.Store(ctx)
	)

	store.Set(key, value)
}

func (k *Keeper) GetLeaseCount(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.LeaseCountKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var count protobuf.UInt64Value
	k.cdc.MustUnmarshal(value, &count)

	return count.GetValue()
}

func (k *Keeper) SetLease(ctx sdk.Context, lease v3.Lease) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(lease.ID)
		value = k.cdc.MustMarshal(&lease)
	)

	store.Set(key, value)
}

func (k *Keeper) HasLease(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetLease(ctx sdk.Context, id uint64) (lease v3.Lease, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return lease, false
	}

	k.cdc.MustUnmarshal(value, &lease)
	return lease, true
}

func (k *Keeper) DeleteLease(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetLeases(ctx sdk.Context) (items []v3.Lease) {
	var (
		store = k.Store(ctx)
		iter  = sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	)

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var item v3.Lease
		k.cdc.MustUnmarshal(iter.Value(), &item)

		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateLeases(ctx sdk.Context, fn func(index int, item v3.Lease) (stop bool)) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item v3.Lease
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetLatestLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress) (lease v3.Lease, found bool) {
	store := k.Store(ctx)

	iter := sdk.KVStoreReversePrefixIterator(store, types.GetLeaseForProviderByNodeKeyPrefix(provAddr, nodeAddr))
	defer iter.Close()

	if iter.Valid() {
		lease, found = k.GetLease(ctx, types.IDFromLeaseForProviderByNodeKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("lease for provider by node key %X does not exist", iter.Key()))
		}
	}

	return lease, found
}

func (k *Keeper) SetLeasePayoutForNextAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeasePayoutForNextAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteLeasePayoutForNextAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeasePayoutForNextAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IterateLeasesForNextAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Lease) (stop bool)) {
	store := k.Store(ctx)

	iter := store.Iterator(types.LeasePayoutForNextAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeasePayoutForNextAtKeyPrefix(at)))
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		lease, found := k.GetLease(ctx, types.IDFromLeasePayoutForNextAtKey(iter.Key()))
		if !found {
			panic(fmt.Errorf("lease for next_at key %X does not exist", iter.Key()))
		}

		if stop := fn(i, lease); stop {
			break
		}
		i++
	}
}
