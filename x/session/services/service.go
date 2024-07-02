package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/services/v1"
	"github.com/sentinel-official/hub/v12/x/session/services/v2"
	v1types "github.com/sentinel-official/hub/v12/x/session/types/v1"
	v2types "github.com/sentinel-official/hub/v12/x/session/types/v2"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1.NewMsgServiceServer(k))
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1.NewQueryServiceServer(k))

	v2types.RegisterMsgServiceServer(configurator.MsgServer(), v2.NewMsgServiceServer(k))
	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2.NewQueryServiceServer(k))
}
