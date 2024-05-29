package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	var (
		store = k.Store(ctx)
		value = k.cdc.MustMarshal(&params)
	)

	store.Set(types.ParamsKey, value)
}

func (k *Keeper) GetParams(ctx sdk.Context) (v types.Params) {
	var (
		store = k.Store(ctx)
		value = store.Get(types.ParamsKey)
	)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

func (k *Keeper) GetChannelID(ctx sdk.Context) string {
	return k.GetParams(ctx).ChannelID
}

func (k *Keeper) GetPortID(ctx sdk.Context) string {
	return k.GetParams(ctx).PortID
}

func (k *Keeper) GetTimeoutDuration(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).TimeoutDuration
}

func (k *Keeper) GetTimeout(ctx sdk.Context) int64 {
	return ctx.BlockTime().Add(
		k.GetTimeoutDuration(ctx),
	).UnixNano()
}
