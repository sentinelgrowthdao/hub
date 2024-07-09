package expected

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	base "github.com/sentinel-official/hub/v12/types"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
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
}

type SessionKeeper interface {
	GetCount(ctx sdk.Context) uint64
	SetCount(ctx sdk.Context, count uint64)
	SetSession(ctx sdk.Context, session sessiontypes.Session)
	SetSessionForAccount(ctx sdk.Context, addr sdk.AccAddress, id uint64)
	SetSessionForInactiveAt(ctx sdk.Context, at time.Time, id uint64)
	SetSessionForNode(ctx sdk.Context, addr base.NodeAddress, id uint64)
	StatusChangeDelay(ctx sdk.Context) time.Duration
}
