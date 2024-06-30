// DO NOT COVER

package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr base.ProvAddress) bool
}

type NodeKeeper interface {
	HasNode(ctx sdk.Context, addr base.NodeAddress) bool
	SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress)
	DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress)
	GetNodesForPlan(ctx sdk.Context, id uint64) []nodetypes.Node
}

type SubscriptionKeeper interface {
	CreateSubscriptionForPlan(ctx sdk.Context, accAddr sdk.AccAddress, id uint64, denom string) (*subscriptiontypes.PlanSubscription, error)
}
