package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func (k *Keeper) StatusChangeDelay(ctx sdk.Context) (duration time.Duration) {
	k.params.Get(ctx, v2.KeyStatusChangeDelay, &duration)
	return
}

func (k *Keeper) SetParams(ctx sdk.Context, params v2.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) v2.Params {
	return v2.NewParams(
		k.StatusChangeDelay(ctx),
	)
}
