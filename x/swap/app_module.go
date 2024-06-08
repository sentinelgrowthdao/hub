package swap

import (
	"context"
	"encoding/json"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v12/x/swap/client/cli"
	"github.com/sentinel-official/hub/v12/x/swap/keeper"
	"github.com/sentinel-official/hub/v12/x/swap/types"
)

var (
	_ sdkmodule.AppModuleBasic      = AppModuleBasic{}
	_ sdkmodule.AppModuleGenesis    = AppModule{}
	_ sdkmodule.AppModuleSimulation = AppModule{}
	_ sdkmodule.BeginBlockAppModule = AppModule{}
	_ sdkmodule.EndBlockAppModule   = AppModule{}
	_ sdkmodule.HasConsensusVersion = AppModule{}
	_ sdkmodule.HasServices         = AppModule{}
)

type AppModuleBasic struct{}

func (amb AppModuleBasic) Name() string { return types.ModuleName }

func (amb AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (amb AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (amb AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	_ = types.RegisterQueryServiceHandlerClient(context.Background(), mux, types.NewQueryServiceClient(ctx))
}

func (amb AppModuleBasic) GetTxCmd() *cobra.Command { return cli.GetTxCmd() }

func (amb AppModuleBasic) GetQueryCmd() *cobra.Command { return cli.GetQueryCmd() }

type AppModule struct {
	AppModuleBasic
	cdc    codec.Codec
	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: k,
	}
}

func (am AppModule) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	state := types.DefaultGenesisState()
	return jsonCodec.MustMarshalJSON(state)
}

func (am AppModule) ValidateGenesis(jsonCodec codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state types.GenesisState
	if err := jsonCodec.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

func (am AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state types.GenesisState
	jsonCodec.MustUnmarshalJSON(message, &state)
	am.keeper.InitGenesis(ctx, &state)

	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	state := am.keeper.ExportGenesis(ctx)
	return jsonCodec.MustMarshalJSON(state)
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abcitypes.RequestBeginBlock) {}

func (am AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (AppModule) GenerateGenesisState(_ *sdkmodule.SimulationState) {}

func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (am AppModule) WeightedOperations(_ sdkmodule.SimulationState) []sdksimulation.WeightedOperation {
	return nil
}

func (am AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) RegisterServices(configurator sdkmodule.Configurator) {
	types.RegisterMsgServiceServer(configurator.MsgServer(), keeper.NewMsgServiceServer(am.keeper))
	types.RegisterQueryServiceServer(configurator.QueryServer(), keeper.NewQueryServiceServer(am.keeper))
}
