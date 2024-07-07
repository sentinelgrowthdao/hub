package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

func (k *Keeper) InitGenesis(ctx sdk.Context, state *v2.GenesisState) {
	for _, item := range state.Plans {
		addr := item.Plan.GetProviderAddress()
		k.SetPlan(ctx, item.Plan)
		k.SetPlanForProvider(ctx, addr, item.Plan.ID)

		for _, node := range item.Nodes {
			addr, err := base.NodeAddressFromBech32(node)
			if err != nil {
				panic(err)
			}

			k.node.SetNodeForPlan(ctx, item.Plan.ID, addr)
		}
	}

	count := uint64(0)
	for _, item := range state.Plans {
		if item.Plan.ID > count {
			count = item.Plan.ID
		}
	}

	k.SetCount(ctx, count)
}

func (k *Keeper) ExportGenesis(ctx sdk.Context) *v2.GenesisState {
	var (
		plans = k.GetPlans(ctx)
		items = make(v2.GenesisPlans, 0, len(plans))
	)

	for _, plan := range plans {
		item := v2.GenesisPlan{
			Plan:  plan,
			Nodes: []string{},
		}

		nodes := k.node.GetNodesForPlan(ctx, plan.ID)
		for _, node := range nodes {
			item.Nodes = append(item.Nodes, node.Address)
		}

		items = append(items, item)
	}

	return v2.NewGenesisState(items)
}
