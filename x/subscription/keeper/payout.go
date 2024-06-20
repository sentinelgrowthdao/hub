package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func (k *Keeper) SetPayout(ctx sdk.Context, payout v2.Payout) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(payout.ID)
		value = k.cdc.MustMarshal(&payout)
	)

	store.Set(key, value)
}

func (k *Keeper) HasPayout(ctx sdk.Context, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(id)
	)

	return store.Has(key)
}

func (k *Keeper) GetPayout(ctx sdk.Context, id uint64) (payout v2.Payout, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return payout, false
	}

	k.cdc.MustUnmarshal(value, &payout)
	return payout, true
}

func (k *Keeper) DeletePayout(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetPayouts(ctx sdk.Context) (items v2.Payouts) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.PayoutKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v2.Payout
		k.cdc.MustUnmarshal(iterator.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IteratePayouts(ctx sdk.Context, fn func(index int, item v2.Payout) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.PayoutKeyPrefix)
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v2.Payout
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetPayoutForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashPayoutForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeletePayoutForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetPayoutForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashPayoutForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNodeKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeletePayoutForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) SetPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountByNodeKey(accAddr, nodeAddr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr base.NodeAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountByNodeKey(accAddr, nodeAddr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeletePayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForAccountByNodeKey(accAddr, nodeAddr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetLatestPayoutForAccountByNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr base.NodeAddress) (payout v2.Payout, found bool) {
	store := k.Store(ctx)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetPayoutForAccountByNodeKeyPrefix(accAddr, nodeAddr))
	defer iterator.Close()

	if iterator.Valid() {
		payout, found = k.GetPayout(ctx, types.IDFromPayoutForAccountByNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("payout for account by node key %X does not exist", iterator.Key()))
		}
	}

	return payout, found
}

func (k *Keeper) SetPayoutForNextAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNextAtKey(at, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeletePayoutForNextAt(ctx sdk.Context, at time.Time, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PayoutForNextAtKey(at, id)
	)

	store.Delete(key)
}

func (k *Keeper) IteratePayoutsForNextAt(ctx sdk.Context, at time.Time, fn func(index int, item v2.Payout) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.PayoutForNextAtKeyPrefix, sdk.PrefixEndBytes(types.GetPayoutForNextAtKeyPrefix(at)))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		payout, found := k.GetPayout(ctx, types.IDFromPayoutForNextAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("payout for next_at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, payout); stop {
			break
		}
		i++
	}
}
