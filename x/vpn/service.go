package vpn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/deposit"
	"github.com/sentinel-official/hub/v12/x/node"
	"github.com/sentinel-official/hub/v12/x/plan"
	"github.com/sentinel-official/hub/v12/x/provider"
	"github.com/sentinel-official/hub/v12/x/session"
	"github.com/sentinel-official/hub/v12/x/subscription"
	"github.com/sentinel-official/hub/v12/x/vpn/keeper"
)

func RegisterServices(configurator sdkmodule.Configurator, cdc codec.BinaryCodec, k keeper.Keeper) {
	deposit.RegisterServices(configurator, cdc, k.Deposit)
	node.RegisterServices(configurator, cdc, k.Node)
	plan.RegisterServices(configurator, cdc, k.Plan)
	provider.RegisterServices(configurator, cdc, k.Provider)
	session.RegisterServices(configurator, cdc, k.Session)
	subscription.RegisterServices(configurator, cdc, k.Subscription)
}
