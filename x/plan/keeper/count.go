package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/v12/x/plan/types"
)

func (k *Keeper) SetPlanCount(ctx sdk.Context, count uint64) {
	var (
		store = k.Store(ctx)
		key   = types.PlanCountKey
		value = k.cdc.MustMarshal(&protobuf.UInt64Value{Value: count})
	)

	store.Set(key, value)
}

func (k *Keeper) GetPlanCount(ctx sdk.Context) uint64 {
	var (
		store = k.Store(ctx)
		key   = types.PlanCountKey
		value = store.Get(key)
	)

	if value == nil {
		return 0
	}

	var count protobuf.UInt64Value
	k.cdc.MustUnmarshal(value, &count)

	return count.GetValue()
}
