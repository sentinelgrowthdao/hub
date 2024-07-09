package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

// SetParams stores the oracle module parameters in the module's KVStore.
func (k *Keeper) SetParams(ctx sdk.Context, params v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := k.cdc.MustMarshal(&params)

	store.Set(key, value)
}

// GetParams retrieves the oracle module parameters from the module's KVStore.
func (k *Keeper) GetParams(ctx sdk.Context) (v v1.Params) {
	store := k.Store(ctx)
	key := types.ParamsKey
	value := store.Get(key)

	k.cdc.MustUnmarshal(value, &v)
	return v
}

// GetChannelID retrieves the ChannelID parameter from the module's parameters.
func (k *Keeper) GetChannelID(ctx sdk.Context) string {
	return k.GetParams(ctx).ChannelID
}

// GetPortID retrieves the PortID parameter from the module's parameters.
func (k *Keeper) GetPortID(ctx sdk.Context) string {
	return k.GetParams(ctx).PortID
}

// GetTimeoutDuration retrieves the TimeoutDuration parameter from the module's parameters.
func (k *Keeper) GetTimeoutDuration(ctx sdk.Context) time.Duration {
	return k.GetParams(ctx).TimeoutDuration
}

// GetTimeout calculates and returns the timeout time in UnixNano format based on the current block time and the TimeoutDuration parameter.
func (k *Keeper) GetTimeout(ctx sdk.Context) int64 {
	duration := k.GetTimeoutDuration(ctx)
	return ctx.BlockTime().Add(duration).UnixNano()
}
