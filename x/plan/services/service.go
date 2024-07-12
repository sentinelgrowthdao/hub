package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/plan/keeper"
	"github.com/sentinel-official/hub/v12/x/plan/services/v1"
	"github.com/sentinel-official/hub/v12/x/plan/services/v2"
	"github.com/sentinel-official/hub/v12/x/plan/services/v3"
	v1types "github.com/sentinel-official/hub/v12/x/plan/types/v1"
	v2types "github.com/sentinel-official/hub/v12/x/plan/types/v2"
	v3types "github.com/sentinel-official/hub/v12/x/plan/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1.NewQueryServiceServer(k))
	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2.NewQueryServiceServer(k))
}
