package oracle

import (
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

	"github.com/sentinel-official/hub/v12/x/oracle/client/cli"
	"github.com/sentinel-official/hub/v12/x/oracle/keeper"
	"github.com/sentinel-official/hub/v12/x/oracle/types"
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

func (a AppModuleBasic) Name() string { return types.ModuleName }

func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(context client.Context, mux *runtime.ServeMux) {}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

type AppModule struct {
	AppModuleBasic
	k keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		k: k,
	}
}

func (a AppModule) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	state := types.DefaultGenesisState()
	return jsonCodec.MustMarshalJSON(state)
}

func (a AppModule) ValidateGenesis(jsonCodec codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state types.GenesisState
	if err := jsonCodec.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

func (a AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state types.GenesisState
	jsonCodec.MustUnmarshalJSON(message, &state)
	a.k.InitGenesis(ctx, &state)

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	state := a.k.ExportGenesis(ctx)
	return jsonCodec.MustMarshalJSON(state)
}

func (a AppModule) GenerateGenesisState(_ *sdkmodule.SimulationState) {}

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(_ sdkmodule.SimulationState) []sdksimulation.WeightedOperation {
	return nil
}

func (a AppModule) BeginBlock(ctx sdk.Context, req abcitypes.RequestBeginBlock) {}

func (a AppModule) EndBlock(ctx sdk.Context, req abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (a AppModule) ConsensusVersion() uint64 { return 1 }

func (a AppModule) RegisterServices(configurator sdkmodule.Configurator) {
	types.RegisterMsgServiceServer(configurator.MsgServer(), keeper.NewMsgServiceServer(a.k))
}
