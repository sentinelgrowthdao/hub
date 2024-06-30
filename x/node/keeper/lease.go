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
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Lease
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateLeases(ctx sdk.Context, fn func(index int, item v3.Lease) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.LeaseKeyPrefix)
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v3.Lease
		k.cdc.MustUnmarshal(iterator.Value(), &item)

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

func (k *Keeper) GetLatestLeaseForProviderByNode(ctx sdk.Context, provAddr base.ProvAddress, nodeAddr base.NodeAddress) (lease v3.Lease, found bool) {
	store := k.Store(ctx)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetLeaseForProviderByNodeKeyPrefix(provAddr, nodeAddr))
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
	key := types.LeaseForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteLeaseForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.LeaseForInactiveAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateLeasesForInactiveAt(ctx sdk.Context, endTime time.Time, fn func(index int, item v3.Lease) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.LeaseForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForInactiveAtKeyPrefix(endTime)))
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

func (k *Keeper) IterateLeasesForPayoutAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Lease) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.LeaseForPayoutAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForPayoutAtKeyPrefix(at)))
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

func (k *Keeper) SetLeaseForRenewAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForRenewAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteLeaseForRenewAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.LeaseForRenewAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IterateLeasesForRenewAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Lease) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.LeaseForRenewAtKeyPrefix, sdk.PrefixEndBytes(types.GetLeaseForRenewAtKeyPrefix(at)))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		lease, found := k.GetLease(ctx, types.IDFromLeaseForRenewAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("lease for renew key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, lease); stop {
			break
		}
		i++
	}
}

func (k *Keeper) RenewLease(ctx sdk.Context, msg *v3.MsgRenewLeaseRequest) (*v3.Lease, error) {
	lease, found := k.GetLease(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorLeaseNotFound(msg.ID)
	}

	var (
		nodeAddr = lease.GetNodeAddress()
		provAddr = lease.GetProvAddress()
	)

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	price, found := node.HourlyPrice(msg.Denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(msg.Denom)
	}

	lease = v3.Lease{
		ID:          lease.ID,
		ProvAddress: lease.ProvAddress,
		NodeAddress: lease.NodeAddress,
		Price:       price,
		Deposit: sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Hours),
		),
		Hours:     0,
		MaxHours:  msg.Hours,
		CreatedAt: ctx.BlockTime(),
		PayoutAt:  ctx.BlockTime(),
	}

	if err := k.AddDeposit(ctx, provAddr.Bytes(), lease.Deposit); err != nil {
		return nil, err
	}

	duration := time.Duration(lease.MaxHours) * time.Hour
	if msg.Renewable {
		lease.InactiveAt = time.Time{}
		lease.RenewAt = lease.CreatedAt.Add(duration)
	} else {
		lease.RenewAt = time.Time{}
		lease.InactiveAt = lease.CreatedAt.Add(duration)
	}

	k.SetLease(ctx, lease)
	k.SetLeaseForNode(ctx, nodeAddr, lease.ID)
	k.SetLeaseForPayoutAt(ctx, lease.PayoutAt, lease.ID)
	k.SetLeaseForProvider(ctx, provAddr, lease.ID)
	k.SetLeaseForProviderByNode(ctx, provAddr, nodeAddr, lease.ID)

	if msg.Renewable {
		k.SetLeaseForRenewAt(ctx, lease.RenewAt, lease.ID)
	} else {
		k.SetLeaseForInactiveAt(ctx, lease.InactiveAt, lease.ID)
	}

	return &lease, nil
}
