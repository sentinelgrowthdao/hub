package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/sentinel-official/hub/v12/x/oracle/exptected"
)

type Keeper struct {
	cdc        codec.Codec
	key        storetypes.StoreKey
	capability ibcexported.ScopedKeeper
	channel    exptected.ChannelKeeper
	ics4       ibcporttypes.ICS4Wrapper
}

func NewKeeper(
	cdc codec.Codec, key storetypes.StoreKey, capability ibcexported.ScopedKeeper, channel exptected.ChannelKeeper,
	ics4 ibcporttypes.ICS4Wrapper,
) Keeper {
	return Keeper{
		cdc:        cdc,
		key:        key,
		capability: capability,
		channel:    channel,
		ics4:       ics4,
	}
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}
