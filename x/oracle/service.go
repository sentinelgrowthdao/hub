package oracle

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/oracle/keeper"
	v1keeper "github.com/sentinel-official/hub/v12/x/oracle/keeper/v1"
	v1types "github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

func RegisterServices(configurator sdkmodule.Configurator, _ codec.BinaryCodec, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1keeper.NewMsgServiceServer(k))
}
