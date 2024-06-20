package keeper

import (
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) SetSubscription(ctx sdk.Context, subscription v3.Subscription) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(subscription.ID)
		value = k.cdc.MustMarshal(&subscription)
	)

	store.Set(key, value)
}

func (k *Keeper) GetSubscription(ctx sdk.Context, id uint64) (subscription v3.Subscription, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionKey(id)
		value = store.Get(key)
	)

	if value == nil {
		return subscription, false
	}

	k.cdc.MustUnmarshal(value, &subscription)
	return subscription, true
}

func (k *Keeper) DeleteSubscription(ctx sdk.Context, id uint64) {
	key := types.SubscriptionKey(id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) GetSubscriptions(ctx sdk.Context) (items []v3.Subscription) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.SubscriptionKeyPrefix)
	)

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
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForAccountKey(addr, id)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForAccount(ctx sdk.Context, addr sdk.AccAddress) (items []v3.Subscription) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForAccountKeyPrefix(addr))
	)

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
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HashSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) bool {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteSubscriptionForPlan(ctx sdk.Context, planID, subscriptionID uint64) {
	var (
		store = k.Store(ctx)
		key   = types.SubscriptionForPlanKey(planID, subscriptionID)
	)

	store.Delete(key)
}

func (k *Keeper) GetSubscriptionsForPlan(ctx sdk.Context, id uint64) (items []v3.Subscription) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.GetSubscriptionForPlanKeyPrefix(id))
	)

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
	key := types.SubscriptionForInactiveAtKey(at, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store := k.Store(ctx)
	store.Set(key, value)
}

func (k *Keeper) DeleteSubscriptionForInactiveAt(ctx sdk.Context, at time.Time, id uint64) {
	key := types.SubscriptionForInactiveAtKey(at, id)

	store := k.Store(ctx)
	store.Delete(key)
}

func (k *Keeper) IterateSubscriptionsForInactiveAt(ctx sdk.Context, endTime time.Time, fn func(index int, item v3.Subscription) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.SubscriptionForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetSubscriptionForInactiveAtKeyPrefix(endTime)))
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

// CreateSubscriptionForPlan creates a new PlanSubscription for a specific plan and account.
func (k *Keeper) CreateSubscriptionForPlan(ctx sdk.Context, accAddr sdk.AccAddress, id uint64, denom string) (*v3.Subscription, error) {
	// Check if the plan exists and is in an active status.
	plan, found := k.GetPlan(ctx, id)
	if !found {
		return nil, types.NewErrorPlanNotFound(id)
	}
	if !plan.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidPlanStatus(plan.ID, plan.Status)
	}

	// Get the price of the plan in the specified denomination.
	price, found := plan.Price(denom)
	if !found {
		return nil, types.NewErrorPriceNotFound(denom)
	}

	// Calculate the staking reward based on the plan price and staking share.
	var (
		stakingShare  = k.provider.StakingShare(ctx)
		stakingReward = baseutils.GetProportionOfCoin(price, stakingShare)
	)

	// Move the staking reward from the account to the fee collector module account.
	if err := k.SendCoinFromAccountToModule(ctx, accAddr, k.feeCollectorName, stakingReward); err != nil {
		return nil, err
	}

	// Calculate the payment amount after deducting the staking reward.
	var (
		provAddr = plan.GetProviderAddress()
		payment  = price.Sub(stakingReward)
	)

	// Send the payment amount from the account to the plan provider address.
	if err := k.SendCoin(ctx, accAddr, provAddr.Bytes(), payment); err != nil {
		return nil, err
	}

	// Emit an event for the plan payment.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventPayForPlan{
			Address:         accAddr.String(),
			Payment:         payment.String(),
			ProviderAddress: plan.ProviderAddress,
			StakingReward:   stakingReward.String(),
			ID:              plan.ID,
		},
	)

	// Retrieve the current count and create a new PlanSubscription.
	count := k.GetSubscriptionCount(ctx)
	subscription := v3.Subscription{
		ID:         count + 1,
		AccAddress: accAddr.String(),
		PlanID:     plan.ID,
		Price:      price,
		Status:     v1base.StatusActive,
		StatusAt:   ctx.BlockTime(),
		InactiveAt: ctx.BlockTime().Add(plan.Duration),
	}

	// Save the new PlanSubscription to the store and update the count.
	k.SetSubscriptionCount(ctx, count+1)
	k.SetSubscription(ctx, subscription)
	k.SetSubscriptionForAccount(ctx, accAddr, subscription.ID)
	k.SetSubscriptionForPlan(ctx, plan.ID, subscription.ID)
	k.SetSubscriptionForInactiveAt(ctx, subscription.InactiveAt, subscription.ID)

	// Create an allocation for the plan subscription and emit an event.
	alloc := v2.Allocation{
		ID:            subscription.ID,
		Address:       accAddr.String(),
		GrantedBytes:  base.Gigabyte.MulRaw(plan.Gigabytes),
		UtilisedBytes: sdkmath.ZeroInt(),
	}

	k.SetAllocation(ctx, alloc)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventAllocate{
			Address:       alloc.Address,
			GrantedBytes:  alloc.GrantedBytes,
			UtilisedBytes: alloc.UtilisedBytes,
			ID:            alloc.ID,
		},
	)

	return &subscription, nil
}
