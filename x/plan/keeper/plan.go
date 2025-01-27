package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

// SetActivePlan stores an active plan in the module's KVStore.
func (k *Keeper) SetActivePlan(ctx sdk.Context, plan v3.Plan) {
	store := k.Store(ctx)
	key := types.ActivePlanKey(plan.ID)
	value := k.cdc.MustMarshal(&plan)

	store.Set(key, value)
}

// HasActivePlan checks if an active plan exists in the module's KVStore.
func (k *Keeper) HasActivePlan(ctx sdk.Context, id uint64) bool {
	store := k.Store(ctx)
	key := types.ActivePlanKey(id)

	return store.Has(key)
}

// GetActivePlan retrieves an active plan from the module's KVStore.
func (k *Keeper) GetActivePlan(ctx sdk.Context, id uint64) (plan v3.Plan, found bool) {
	store := k.Store(ctx)
	key := types.ActivePlanKey(id)
	value := store.Get(key)

	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(value, &plan)
	return plan, true
}

// DeleteActivePlan removes an active plan from the module's KVStore.
func (k *Keeper) DeleteActivePlan(ctx sdk.Context, id uint64) {
	store := k.Store(ctx)
	key := types.ActivePlanKey(id)

	store.Delete(key)
}

// SetInactivePlan stores an inactive plan in the module's KVStore.
func (k *Keeper) SetInactivePlan(ctx sdk.Context, plan v3.Plan) {
	store := k.Store(ctx)
	key := types.InactivePlanKey(plan.ID)
	value := k.cdc.MustMarshal(&plan)

	store.Set(key, value)
}

// HasInactivePlan checks if an inactive plan exists in the module's KVStore.
func (k *Keeper) HasInactivePlan(ctx sdk.Context, id uint64) bool {
	store := k.Store(ctx)
	key := types.InactivePlanKey(id)

	return store.Has(key)
}

// GetInactivePlan retrieves an inactive plan from the module's KVStore.
func (k *Keeper) GetInactivePlan(ctx sdk.Context, id uint64) (plan v3.Plan, found bool) {
	store := k.Store(ctx)
	key := types.InactivePlanKey(id)
	value := store.Get(key)

	if value == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(value, &plan)
	return plan, true
}

// DeleteInactivePlan removes an inactive plan from the module's KVStore.
func (k *Keeper) DeleteInactivePlan(ctx sdk.Context, id uint64) {
	store := k.Store(ctx)
	key := types.InactivePlanKey(id)

	store.Delete(key)
}

// SetPlan stores a plan in the KVStore based on its status (Active or Inactive).
func (k *Keeper) SetPlan(ctx sdk.Context, plan v3.Plan) {
	switch plan.Status {
	case v1base.StatusActive:
		k.SetActivePlan(ctx, plan)
	case v1base.StatusInactive:
		k.SetInactivePlan(ctx, plan)
	default:
		panic(fmt.Errorf("failed to set the plan %v", plan))
	}
}

// HasPlan checks if a plan exists in either active or inactive state.
func (k *Keeper) HasPlan(ctx sdk.Context, id uint64) bool {
	return k.HasActivePlan(ctx, id) || k.HasInactivePlan(ctx, id)
}

// GetPlan retrieves a plan from either the active or inactive state.
func (k *Keeper) GetPlan(ctx sdk.Context, id uint64) (plan v3.Plan, found bool) {
	plan, found = k.GetActivePlan(ctx, id)
	if found {
		return plan, true
	}

	plan, found = k.GetInactivePlan(ctx, id)
	if found {
		return plan, true
	}

	return plan, false
}

// GetPlans retrieves all plans (both active and inactive) from the module's KVStore.
func (k *Keeper) GetPlans(ctx sdk.Context) (items []v3.Plan) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.PlanKeyPrefix)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v3.Plan
		k.cdc.MustUnmarshal(iterator.Value(), &item)

		items = append(items, item)
	}

	return items
}

// SetPlanForProvider associates a plan ID with a provider address.
func (k *Keeper) SetPlanForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	store := k.Store(ctx)
	key := types.PlanForProviderKey(addr, id)
	value := k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})

	store.Set(key, value)
}

// DeletePlanForProvider removes the association between a plan ID and a provider address.
func (k *Keeper) DeletePlanForProvider(ctx sdk.Context, addr base.ProvAddress, id uint64) {
	store := k.Store(ctx)
	key := types.PlanForProviderKey(addr, id)

	store.Delete(key)
}

// GetPlansForProvider retrieves all plans associated with a specific provider address.
func (k *Keeper) GetPlansForProvider(ctx sdk.Context, addr base.ProvAddress) (items []v3.Plan) {
	store := k.Store(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.GetPlanForProviderKeyPrefix(addr))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetPlan(ctx, types.IDFromPlanForProviderKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("plan for provider key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}
