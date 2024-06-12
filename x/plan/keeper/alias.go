package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types"
)

func (k *Keeper) HasProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}

func (k *Keeper) HasNode(ctx sdk.Context, addr base.NodeAddress) bool {
	return k.node.HasNode(ctx, addr)
}

func (k *Keeper) SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	k.node.SetNodeForPlan(ctx, id, addr)
}

func (k *Keeper) DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress) {
	k.node.DeleteNodeForPlan(ctx, id, addr)
}

func (k *Keeper) GetNodesForPlan(ctx sdk.Context, id uint64) nodetypes.Nodes {
	return k.node.GetNodesForPlan(ctx, id)
}

func (k *Keeper) CreateSubscriptionForPlan(ctx sdk.Context, accAddr sdk.AccAddress, id uint64, denom string) (*subscriptiontypes.PlanSubscription, error) {
	return k.subscription.CreateSubscriptionForPlan(ctx, accAddr, id, denom)
}
