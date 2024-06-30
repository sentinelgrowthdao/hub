package subscription

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	v1keeper "github.com/sentinel-official/hub/v12/x/subscription/keeper/v1"
	v2keeper "github.com/sentinel-official/hub/v12/x/subscription/keeper/v2"
	v3keeper "github.com/sentinel-official/hub/v12/x/subscription/keeper/v3"
	v1types "github.com/sentinel-official/hub/v12/x/subscription/types/v1"
	v2types "github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	v3types "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, cdc codec.BinaryCodec, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1keeper.NewMsgServiceServer(k))
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1keeper.NewQueryServiceServer(cdc, k))

	v2types.RegisterMsgServiceServer(configurator.MsgServer(), v2keeper.NewMsgServiceServer(k))
	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2keeper.NewQueryServiceServer(cdc, k))

	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3keeper.NewMsgServiceServer(k))
	v3types.RegisterQueryServiceServer(configurator.QueryServer(), v3keeper.NewQueryServiceServer(cdc, k))
}