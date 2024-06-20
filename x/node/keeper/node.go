package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func (k *Keeper) SetActiveNode(ctx sdk.Context, node v2.Node) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(node.GetAddress())
		value = k.cdc.MustMarshal(&node)
	)

	store.Set(key, value)
}

func (k *Keeper) HasActiveNode(ctx sdk.Context, addr base.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) GetActiveNode(ctx sdk.Context, addr base.NodeAddress) (v v2.Node, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteActiveNode(ctx sdk.Context, addr base.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.ActiveNodeKey(addr)
	)

	store.Delete(key)
}

func (k *Keeper) SetInactiveNode(ctx sdk.Context, node v2.Node) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(node.GetAddress())
		value = k.cdc.MustMarshal(&node)
	)

	store.Set(key, value)
}

func (k *Keeper) HasInactiveNode(ctx sdk.Context, addr base.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(addr)
	)

	return store.Has(key)
}

func (k *Keeper) GetInactiveNode(ctx sdk.Context, addr base.NodeAddress) (v v2.Node, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(addr)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteInactiveNode(ctx sdk.Context, addr base.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.InactiveNodeKey(addr)
	)

	store.Delete(key)
}

func (k *Keeper) SetNode(ctx sdk.Context, node v2.Node) {
	switch node.Status {
	case v1base.StatusActive:
		k.SetActiveNode(ctx, node)
	case v1base.StatusInactive:
		k.SetInactiveNode(ctx, node)
	default:
		panic(fmt.Errorf("failed to set the node %v", node))
	}
}

func (k *Keeper) HasNode(ctx sdk.Context, addr base.NodeAddress) bool {
	return k.HasActiveNode(ctx, addr) || k.HasInactiveNode(ctx, addr)
}

func (k *Keeper) GetNode(ctx sdk.Context, addr base.NodeAddress) (node v2.Node, found bool) {
	node, found = k.GetActiveNode(ctx, addr)
	if found {
		return
	}

	node, found = k.GetInactiveNode(ctx, addr)
	if found {
		return
	}

	return node, false
}

func (k *Keeper) GetNodes(ctx sdk.Context) (items v2.Nodes) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var item v2.Node
		k.cdc.MustUnmarshal(iterator.Value(), &item)
		items = append(items, item)
	}

	return items
}

func (k *Keeper) IterateNodes(ctx sdk.Context, fn func(index int, item v2.Node) (stop bool)) {
	store := k.Store(ctx)

	iterator := sdk.KVStorePrefixIterator(store, types.NodeKeyPrefix)
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		var node v2.Node
		k.cdc.MustUnmarshal(iterator.Value(), &node)

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetNodeForInactiveAt(ctx sdk.Context, at time.Time, addr base.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForInactiveAtKey(at, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) DeleteNodeForInactiveAt(ctx sdk.Context, at time.Time, addr base.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForInactiveAtKey(at, addr)
	)

	store.Delete(key)
}

func (k *Keeper) IterateNodesForInactiveAt(ctx sdk.Context, at time.Time, fn func(index int, item v2.Node) (stop bool)) {
	store := k.Store(ctx)

	iterator := store.Iterator(types.NodeForInactiveAtKeyPrefix, sdk.PrefixEndBytes(types.GetNodeForInactiveAtKeyPrefix(at)))
	defer iterator.Close()

	for i := 0; iterator.Valid(); iterator.Next() {
		node, found := k.GetNode(ctx, types.AddressFromNodeForInactiveAtKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("node for inactive at key %X does not exist", iterator.Key()))
		}

		if stop := fn(i, node); stop {
			break
		}
		i++
	}
}

func (k *Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForPlanKey(id, addr)
		value = k.cdc.MustMarshal(&protobuf.BoolValue{Value: true})
	)

	store.Set(key, value)
}

func (k *Keeper) HasNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) bool {
	var (
		store = k.Store(ctx)
		key   = types.NodeForPlanKey(id, addr)
	)

	return store.Has(key)
}

func (k *Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	var (
		store = k.Store(ctx)
		key   = types.NodeForPlanKey(id, addr)
	)

	store.Delete(key)
}

func (k *Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) (items v2.Nodes) {
	var (
		store    = k.Store(ctx)
		iterator = sdk.KVStorePrefixIterator(store, types.GetNodeForPlanKeyPrefix(id))
	)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		item, found := k.GetNode(ctx, types.AddressFromNodeForPlanKey(iterator.Key()))
		if !found {
			panic(fmt.Errorf("node for plan key %X does not exist", iterator.Key()))
		}

		items = append(items, item)
	}

	return items
}
