package deposit

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/deposit/keeper"
	v1keeper "github.com/sentinel-official/hub/v12/x/deposit/keeper/v1"
	v1types "github.com/sentinel-official/hub/v12/x/deposit/types/v1"
)

func RegisterServices(configurator sdkmodule.Configurator, cdc codec.BinaryCodec, k keeper.Keeper) {
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1keeper.NewQueryServiceServer(cdc, k))
}
