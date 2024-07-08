package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) SetSubscription(ctx sdk.Context, subscription v3.Subscription) {
	store := k.Store(ctx)
	key := types.SubscriptionKey(subscription.ID)
	value := k.cdc.MustMarshal(&subscription)

	store.Set(key, value)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription v3.Subscription, found bool) {
	store := k.Store(ctx)
	key := types.SubscriptionKey(id)
	value := store.Get(key)

	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshal(value, &subscription)
	return subscription, true
}

func (k *Keeper) DeleteSubscription(ctx sdk.Context, id uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionKey(id)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptions(ctx sdk.Context) (items []v3.Subscription) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Subscription
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateSubscriptions(ctx sdk.Context, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var item v3.Subscription
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForAccountKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

func (k *Keeper) HasSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	store := k.Store(ctx)
	key := types.SubscriptionForAccountKey(addr, id)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForAccountKey(addr, id)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items []v3.Subscription) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetSubscriptionForAccountKeyPrefix(addr))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForAccountKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for account key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForPlanKey(planID, subscriptionID)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

func (k *Keeper) HasSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) bool {
	store := k.Store(ctx)
	key := types.SubscriptionForPlanKey(planID, subscriptionID)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	store := k.Store(ctx)
	key := types.SubscriptionForPlanKey(planID, subscriptionID)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForPlan(ctx sdk.Context, id uint64) (items []v3.Subscription) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetSubscriptionForPlanKeyPrefix(id))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForPlanKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for plan key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}

func (k *Keeper) SetSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

func (k *Keeper) DeleteSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForInactiveAtKey(at, id)

	store.Delete(key)
}

func (k *Keeper) IterateSubscriptionsForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.SubscriptionForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetSubscriptionForInactiveAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetSubscriptionForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForRenewalAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

func (k *Keeper) DeleteSubscriptionForRenewalAt(ctx sdk.Context, at time.Time, id uint64) {
	if at.IsZero() {
		return
	}

	store := k.Store(ctx)
	key := types.SubscriptionForRenewalAtKey(at, id)

	store.Delete(key)
}

func (k *Keeper) IterateSubscriptionsForRenewalAt(ctx sdk.Context, at time.Time, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)
	iterator := store.Iterator(types.SubscriptionForRenewalAtKeyPrefix, sdk.PrefixEndBytes(types.GetSubscriptionForRenewalAtKeyPrefix(at)))

	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		item, found := k.GetSubscription(ctx, types.IDFromSubscriptionForRenewalAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("subscription for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, item); stop {
			break
		}
		i++
	}
}
