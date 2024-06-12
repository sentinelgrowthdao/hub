// DO NOT COVER

package expected

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v2"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v2"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v2"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type DepositKeeper interface {
	Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coins sdk.Coins) error
	SendCoinsFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coins sdk.Coins) error
}

type ProviderKeeper interface {
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
}

type NodeKeeper interface {
	StakingShare(ctx sdk.Context) sdkmath.LegacyDec
	GetNode(ctx sdk.Context, addr base.NodeAddress) (nodetypes.Node, bool)
}

type PlanKeeper interface {
	GetPlan(ctx sdk.Context, id uint64) (plantypes.Plan, bool)
}

type SessionKeeper interface {
	GetSession(ctx sdk.Context, id uint64) (sessiontypes.Session, bool)
	SubscriptionInactivePendingHook(ctx sdk.Context, id uint64) error
}
