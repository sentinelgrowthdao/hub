package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

func (k *Keeper) Deposit(ctx sdk.Context) (v sdk.Coin) {
	k.params.Get(ctx, v2.KeyDeposit, &v)
	return
}

func (k *Keeper) StakingShare(ctx sdk.Context) (v sdkmath.LegacyDec) {
	k.params.Get(ctx, v2.KeyStakingShare, &v)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params v2.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) v2.Params {
	return v2.NewParams(
		k.Deposit(ctx),
		k.StakingShare(ctx),
	)
}
