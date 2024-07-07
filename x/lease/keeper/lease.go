package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) SetCount(ctx sdk.Context, count uint64) {
	var (
		key   = types.CountKey
		value = k.cdc.MustMarshal(&protobuf.UInt64Value{Value: count})
		store = k.Store(ctx)
	)

	store.Set(key, value)
}

func (k *Keeper) GetCount(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.CountKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var count protobuf.UInt64Value
	k.cdc.MustUnmarshal(value, &count)

	return count.GetValue()
}

func (k *Keeper) SetLease(ctx sdk.Context, lease v1.Lease) {
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

func (k *Keeper) GetLease(ctx sdk.Context, id uint64) (lease v1.Lease, found bool) {
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

func (k *Keeper) GetLeases(ctx sdk.Context) (items []v1.Lease) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v1.Lease
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateLeases(ctx sdk.Context, fn func(index int, item v1.Lease) (stop bool)) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v1.Lease
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasLeaseForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) bool {
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

func (k *Keeper) SetLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasLeaseForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) bool {
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

func (k *Keeper) SetLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForProviderByNodeKey(provAddr, nodeAddr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) bool {
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

func (k *Keeper) GetLatestLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress) (lease v1.Lease, found bool) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStoreReversePrefixIterator(store, types.GetLeaseForProviderByNodeKeyPrefix(provAddr, nodeAddr))
	)

	defer iterator.Close()

	if iterator.Valid() {
		lease, found = k.GetLease(ctx, types.IDFromLeaseForProviderByNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for provider by node key %X does not exist", iterator.Key()))
		}
	}

	return lease, found
}

func (k *Keeper) SetLeaseForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForInactiveAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteLeaseForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForInactiveAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IterateLeasesForInactiveAt(ctx sdk.Context, endTime time.Time, fn func(index int, item v1.Lease) (stop bool)) {
	var (
		store    = k.Store(ctx)
		iterator = store.Iterator(types.LeaseForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForInactiveAtKeyPrefix(endTime)))
	)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetLease(ctx, types.IDFromLeaseForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetLeaseForPayoutAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForPayoutAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteLeaseForPayoutAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForPayoutAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IterateLeasesForPayoutAt(ctx sdk.Context, at time.Time, fn func(index int, item v1.Lease) (stop bool)) {
	var (
		store    = k.Store(ctx)
		iterator = store.Iterator(types.LeaseForPayoutAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForPayoutAtKeyPrefix(at)))
	)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		lease, found := k.GetLease(ctx, types.IDFromLeaseForPayoutAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for payout_at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, lease); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetLeaseForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForRenewalAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteLeaseForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForRenewalAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IterateLeasesForRenewalAt(ctx sdk.Context, at time.Time, fn func(index int, item v1.Lease) (stop bool)) {
	var (
		store    = k.Store(ctx)
		iterator = store.Iterator(types.LeaseForRenewalAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForRenewalAtKeyPrefix(at)))
	)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		lease, found := k.GetLease(ctx, types.IDFromLeaseForRenewalAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for renew key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, lease); stop {
			break
		}
		i++
	}
}
