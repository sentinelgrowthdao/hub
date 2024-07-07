package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) SetSession(ctx sdk.Context, session v3.Session) {
	var (
		store = k.Store(ctx)
		key   = types.SessionKey(session.GetID())
	)

	value, err := k.cdc.MarshalInterface(session)
	if err != nil {
		panic(err)
	}

	store.Set(key, value)
}

func (k *Keeper) GetSession(ctx sdk.Context, id uint64) (session v3.Session, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.SessionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return session, false
	}
	if err := k.cdc.UnmarshalInterface(value, &session); err != nil {
		panic(err)
	}

	return session, true
}

func (k *Keeper) DeleteSession(ctx sdk.Context, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionKey(id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessions(ctx sdk.Context) (items []v3.Session) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Session
		if err := k.cdc.UnmarshalInterface(iterator.Value(), &item); err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateSessions(ctx sdk.Context, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.SessionKeyPrefix)
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v3.Session
		if err := k.cdc.UnmarshalInterface(iterator.Value(), &item); err != nil {
			panic(err)
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items []v3.Session) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.GetSessionForAccountKeyPrefix(addr))
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForAccountKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for account key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForNodeKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForNodeKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessionsForNode(ctx sdk.Context, addr base.NodeAddress) (items []v3.Session) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.GetSessionForNodeKeyPrefix(addr))
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForNodeKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for node key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForSubscriptionKey(subscriptionID, sessionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForSubscription(ctx sdk.Context, subscriptionID, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForSubscriptionKey(subscriptionID, sessionID)
	)

	store.Delete(key)
}

func (k *Keeper) GetSessionsForSubscription(ctx sdk.Context, id uint64) (items []v3.Session) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(id))
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateSessionsForSubscription(ctx sdk.Context, id uint64, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(id))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		session, found := k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForAllocationKey(subscriptionID, addr, sessionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress, sessionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SessionForAllocationKey(subscriptionID, addr, sessionID)
	)

	store.Delete(key)
}

func (k *Keeper) IterateSessionsForAllocation(ctx sdk.Context, id uint64, addr sdk.AccAddress, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForAllocationKeyPrefix(id, addr))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		session, found := k.GetSession(ctx, types.IDFromSessionForAllocationKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription allocation key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SessionForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SessionForInactiveAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateSessionsForInactiveAt(ctx sdk.Context, end time.Time, fn func(index int, item v3.Session) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.SessionForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetSessionForInactiveAtKeyPrefix(end)))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		session, found := k.GetSession(ctx, types.IDFromSessionForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, session); stop {
			break
		}
		i++
	}
}

func (k *Keeper) GetLatestSessionForSubscription(ctx sdk.Context, subscriptionID uint64) (session v3.Session, found bool) {
	store := k.Store(ctx)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForSubscriptionKeyPrefix(subscriptionID))
	defer iterator.Close()

	if iterator.Valid() {
		session, found = k.GetSession(ctx, types.IDFromSessionForSubscriptionKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription key %X does not exist", iterator.Key()))
		}
	}

	return session, false
}

func (k *Keeper) GetLatestSessionForAllocation(ctx sdk.Context, subscriptionID uint64, addr sdk.AccAddress) (session v3.Session, found bool) {
	store := k.Store(ctx)

	iterator := sdk.KVStoreReversePrefixIterator(store, types.GetSessionForAllocationKeyPrefix(subscriptionID, addr))
	defer iterator.Close()

	if iterator.Valid() {
		session, found = k.GetSession(ctx, types.IDFromSessionForAllocationKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("session for subscription allocation key %X does not exist", iterator.Key()))
		}
	}

	return session, found
}
