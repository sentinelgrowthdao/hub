package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type DepositKeeper interface {
	AddDeposit(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
	SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool)
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr base.ProvAddress) bool
}
