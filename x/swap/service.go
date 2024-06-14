package swap

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/swap/keeper"
	v1keeper "github.com/sentinel-official/hub/v12/x/swap/keeper/v1"
	v1types "github.com/sentinel-official/hub/v12/x/swap/types/v1"
)

func RegisterServices(configurator sdkmodule.Configurator, cdc codec.BinaryCodec, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1keeper.NewMsgServiceServer(k))
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1keeper.NewQueryServiceServer(cdc, k))
}
