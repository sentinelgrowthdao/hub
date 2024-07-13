package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type NodeKeeper interface {
	HasNode(ctx sdk.Context, addr base.NodeAddress) bool
	SetNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress)
	DeleteNodeForPlan(ctx sdk.Context, id uint64, addr base.NodeAddress)
	GetNodesForPlan(ctx sdk.Context, id uint64) []nodetypes.Node
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr base.ProvAddress) bool
}

type SubscriptionKeeper interface {
	HandleMsgStartSession(ctx sdk.Context, msg *subscriptiontypes.MsgStartSessionRequest) (*subscriptiontypes.MsgStartSessionResponse, error)
	HandleMsgStartSubscription(ctx sdk.Context, msg *subscriptiontypes.MsgStartSubscriptionRequest) (*subscriptiontypes.MsgStartSubscriptionResponse, error)
}
