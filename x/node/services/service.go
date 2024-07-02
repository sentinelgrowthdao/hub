package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/services/v1"
	"github.com/sentinel-official/hub/v12/x/node/services/v2"
	"github.com/sentinel-official/hub/v12/x/node/services/v3"
	v1types "github.com/sentinel-official/hub/v12/x/node/types/v1"
	v2types "github.com/sentinel-official/hub/v12/x/node/types/v2"
	v3types "github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1.NewMsgServiceServer(k))
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1.NewQueryServiceServer(k))

	v2types.RegisterMsgServiceServer(configurator.MsgServer(), v2.NewMsgServiceServer(k))
	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2.NewQueryServiceServer(k))

	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))
	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3.NewQueryServiceServer(k))
}