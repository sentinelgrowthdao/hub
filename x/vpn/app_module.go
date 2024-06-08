package vpn

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

	depositkeeper "github.com/sentinel-official/hub/v12/x/deposit/keeper"
	deposittypes "github.com/sentinel-official/hub/v12/x/deposit/types"
	nodekeeper "github.com/sentinel-official/hub/v12/x/node/keeper"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types"
	plankeeper "github.com/sentinel-official/hub/v12/x/plan/keeper"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types"
	providerkeeper "github.com/sentinel-official/hub/v12/x/provider/keeper"
	providertypes "github.com/sentinel-official/hub/v12/x/provider/types"
	sessionkeeper "github.com/sentinel-official/hub/v12/x/session/keeper"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types"
	subscriptionkeeper "github.com/sentinel-official/hub/v12/x/subscription/keeper"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/vpn/client/cli"
	"github.com/sentinel-official/hub/v12/x/vpn/expected"
	"github.com/sentinel-official/hub/v12/x/vpn/keeper"
	"github.com/sentinel-official/hub/v12/x/vpn/types"
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
	_ = deposittypes.RegisterQueryServiceHandlerClient(context.Background(), mux, deposittypes.NewQueryServiceClient(ctx))
	_ = nodetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, nodetypes.NewQueryServiceClient(ctx))
	_ = plantypes.RegisterQueryServiceHandlerClient(context.Background(), mux, plantypes.NewQueryServiceClient(ctx))
	_ = providertypes.RegisterQueryServiceHandlerClient(context.Background(), mux, providertypes.NewQueryServiceClient(ctx))
	_ = sessiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, sessiontypes.NewQueryServiceClient(ctx))
	_ = subscriptiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, subscriptiontypes.NewQueryServiceClient(ctx))
}

func (amb AppModuleBasic) GetTxCmd() *cobra.Command { return cli.GetTxCmd() }

func (amb AppModuleBasic) GetQueryCmd() *cobra.Command { return cli.GetQueryCmd() }

type AppModule struct {
	AppModuleBasic
	cdc     codec.Codec
	account expected.AccountKeeper
	bank    expected.BankKeeper
	keeper  keeper.Keeper
}

func NewAppModule(cdc codec.Codec, account expected.AccountKeeper, bank expected.BankKeeper, k keeper.Keeper) AppModule {
	return AppModule{
		cdc:     cdc,
		account: account,
		bank:    bank,
		keeper:  k,
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

func (am AppModule) BeginBlock(ctx sdk.Context, _ abcitypes.RequestBeginBlock) {
	am.keeper.BeginBlock(ctx)
}

func (am AppModule) EndBlock(ctx sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return am.keeper.EndBlock(ctx)
}

func (am AppModule) GenerateGenesisState(_ *sdkmodule.SimulationState) {}

func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (am AppModule) WeightedOperations(_ sdkmodule.SimulationState) []sdksimulation.WeightedOperation {
	return nil
}

func (am AppModule) ConsensusVersion() uint64 { return 3 }

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) RegisterServices(configurator sdkmodule.Configurator) {
	nodetypes.RegisterMsgServiceServer(configurator.MsgServer(), nodekeeper.NewMsgServiceServer(am.keeper.Node))
	plantypes.RegisterMsgServiceServer(configurator.MsgServer(), plankeeper.NewMsgServiceServer(am.keeper.Plan))
	providertypes.RegisterMsgServiceServer(configurator.MsgServer(), providerkeeper.NewMsgServiceServer(am.keeper.Provider))
	sessiontypes.RegisterMsgServiceServer(configurator.MsgServer(), sessionkeeper.NewMsgServiceServer(am.keeper.Session))
	subscriptiontypes.RegisterMsgServiceServer(configurator.MsgServer(), subscriptionkeeper.NewMsgServiceServer(am.keeper.Subscription))

	deposittypes.RegisterQueryServiceServer(configurator.QueryServer(), depositkeeper.NewQueryServiceServer(am.keeper.Deposit))
	nodetypes.RegisterQueryServiceServer(configurator.QueryServer(), nodekeeper.NewQueryServiceServer(am.keeper.Node))
	plantypes.RegisterQueryServiceServer(configurator.QueryServer(), plankeeper.NewQueryServiceServer(am.keeper.Plan))
	providertypes.RegisterQueryServiceServer(configurator.QueryServer(), providerkeeper.NewQueryServiceServer(am.keeper.Provider))
	sessiontypes.RegisterQueryServiceServer(configurator.QueryServer(), sessionkeeper.NewQueryServiceServer(am.keeper.Session))
	subscriptiontypes.RegisterQueryServiceServer(configurator.QueryServer(), subscriptionkeeper.NewQueryServiceServer(am.keeper.Subscription))
}
