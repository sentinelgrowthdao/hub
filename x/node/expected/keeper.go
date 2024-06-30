// DO NOT COVER

package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
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
	Add(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
	Subtract(ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, addr base.ProvAddress) bool
}
