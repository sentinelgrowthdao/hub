package upgrades

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	ibcicahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	ibcicatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
)

func Handler(mm *module.Manager, configurator module.Configurator, wasmKeeper wasmkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		controllerParams := ibcicacontrollertypes.Params{}
		hostParams := ibcicahosttypes.Params{
			HostEnabled: true,
			AllowMessages: []string{
				sdk.MsgTypeURL(&authz.MsgExec{}),
				sdk.MsgTypeURL(&authz.MsgGrant{}),
				sdk.MsgTypeURL(&authz.MsgRevoke{}),
				sdk.MsgTypeURL(&banktypes.MsgSend{}),
				sdk.MsgTypeURL(&distributiontypes.MsgFundCommunityPool{}),
				sdk.MsgTypeURL(&distributiontypes.MsgSetWithdrawAddress{}),
				sdk.MsgTypeURL(&distributiontypes.MsgWithdrawDelegatorReward{}),
				sdk.MsgTypeURL(&distributiontypes.MsgWithdrawValidatorCommission{}),
				sdk.MsgTypeURL(&govtypes.MsgVote{}),
				sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
				sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
				sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
			},
		}

		icaModule, ok := mm.Modules[ibcicatypes.ModuleName].(ibcica.AppModule)
		if !ok {
			panic("mm.Modules[ibcicatypes.ModuleName] is not of type ibcica.AppModule")
		}

		icaModule.InitModule(ctx, controllerParams, hostParams)

		wasmParams := wasmKeeper.GetParams(ctx)
		wasmParams.CodeUploadAccess = wasmtypes.AllowNobody
		wasmKeeper.SetParams(ctx, wasmParams)

		fromVM[ibcicatypes.ModuleName] = mm.Modules[ibcicatypes.ModuleName].ConsensusVersion()

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
