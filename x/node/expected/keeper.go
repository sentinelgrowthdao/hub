// DO NOT COVER

package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr base.ProvAddress) bool
}

type SubscriptionKeeper interface {
	CreateSubscriptionForNode(ctx sdk.Context, accAddr sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, denom string) (*subscriptiontypes.NodeSubscription, error)
}
