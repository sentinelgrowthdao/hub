package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	depositkeeper "github.com/sentinel-official/hub/v12/x/deposit/keeper"
	leasekeeper "github.com/sentinel-official/hub/v12/x/lease/keeper"
	nodekeeper "github.com/sentinel-official/hub/v12/x/node/keeper"
	plankeeper "github.com/sentinel-official/hub/v12/x/plan/keeper"
	providerkeeper "github.com/sentinel-official/hub/v12/x/provider/keeper"
	sessionkeeper "github.com/sentinel-official/hub/v12/x/session/keeper"
	subscriptionkeeper "github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/vpn/expected"
)

type Keeper struct {
	Deposit      depositkeeper.Keeper
	Lease        leasekeeper.Keeper
	Node         nodekeeper.Keeper
	Plan         plankeeper.Keeper
	Provider     providerkeeper.Keeper
	Session      sessionkeeper.Keeper
	Subscription subscriptionkeeper.Keeper
}

func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, accountKeeper expected.AccountKeeper,
	bankKeeper expected.BankKeeper, distributionKeeper expected.DistributionKeeper, authority, feeCollectorName string,
) Keeper {
	k := Keeper{
		Deposit:      depositkeeper.NewKeeper(cdc, key),
		Lease:        leasekeeper.NewKeeper(cdc, key, authority, feeCollectorName),
		Node:         nodekeeper.NewKeeper(cdc, key, authority, feeCollectorName),
		Plan:         plankeeper.NewKeeper(cdc, key),
		Provider:     providerkeeper.NewKeeper(cdc, key, authority),
		Session:      sessionkeeper.NewKeeper(cdc, key, authority, feeCollectorName),
		Subscription: subscriptionkeeper.NewKeeper(cdc, key, authority, feeCollectorName),
	}

	k.Deposit.WithBankKeeper(bankKeeper)

	k.Lease.WithDepositKeeper(&k.Deposit)
	k.Lease.WithNodeKeeper(&k.Node)
	k.Lease.WithProviderKeeper(&k.Provider)

	k.Node.WithDepositKeeper(&k.Deposit)
	k.Node.WithDistributionKeeper(distributionKeeper)

	k.Plan.WithBankKeeper(bankKeeper)
	k.Plan.WithNodeKeeper(&k.Node)
	k.Plan.WithProviderKeeper(&k.Provider)

	k.Provider.WithDistributionKeeper(distributionKeeper)

	k.Session.WithAccountKeeper(accountKeeper)
	k.Session.WithBankKeeper(bankKeeper)
	k.Session.WithDepositKeeper(&k.Deposit)
	k.Session.WithNodeKeeper(&k.Node)
	k.Session.WithPlanKeeper(&k.Plan)
	k.Session.WithSubscriptionKeeper(nil)

	k.Subscription.WithBankKeeper(bankKeeper)
	k.Subscription.WithDepositKeeper(&k.Deposit)
	k.Subscription.WithNodeKeeper(&k.Node)
	k.Subscription.WithPlanKeeper(&k.Plan)
	k.Subscription.WithProviderKeeper(&k.Provider)
	k.Subscription.WithSessionKeeper(&k.Session)

	return k
}
