package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	protobuf "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

func (k *Keeper) SetAsset(ctx sdk.Context, asset types.Asset) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(asset.Denom)
		value = k.cdc.MustMarshal(&asset)
	)

	store.Set(key, value)
}

func (k *Keeper) GetAsset(ctx sdk.Context, denom string) (v types.Asset, found bool) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(denom)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	k.cdc.MustUnmarshal(value, &v)
	return v, true
}

func (k *Keeper) DeleteAsset(ctx sdk.Context, denom string) {
	var (
		store = k.Store(ctx)
		key   = types.AssetKey(denom)
	)

	store.Delete(key)
}

func (k *Keeper) IterateAssets(ctx sdk.Context, fn func(int, types.Asset) bool) {
	store := k.Store(ctx)

	iter := sdk.KVStorePrefixIterator(store, types.AssetKeyPrefix)
	defer iter.Close()

	for i := 0; iter.Valid(); iter.Next() {
		var item types.Asset
		k.cdc.MustUnmarshal(iter.Value(), &item)

		if stop := fn(i, item); stop {
			break
		}

		i++
	}
}

func (k *Keeper) SetDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64, denom string) {
	var (
		store = k.Store(ctx)
		key   = ibchost.PacketCommitmentKey(portID, channelID, sequence)
		value = k.cdc.MustMarshal(&protobuf.StringValue{Value: denom})
	)

	store.Set(key, value)
}

func (k *Keeper) GetDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) (v string, found bool) {
	var (
		store = k.Store(ctx)
		key   = ibchost.PacketCommitmentKey(portID, channelID, sequence)
		value = store.Get(key)
	)

	if value == nil {
		return v, false
	}

	var denom protobuf.StringValue
	k.cdc.MustUnmarshal(value, &denom)

	return denom.GetValue(), true
}

func (k *Keeper) DeleteDenomForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) {
	var (
		store = k.Store(ctx)
		key   = ibchost.PacketCommitmentKey(portID, channelID, sequence)
	)

	store.Delete(key)
}

func (k *Keeper) GetAssetForPacket(ctx sdk.Context, portID, channelID string, sequence uint64) (v types.Asset, err error) {
	denom, found := k.GetDenomForPacket(ctx, portID, channelID, sequence)
	if !found {
		return v, types.NewErrorDenomtNotFound(portID, channelID, sequence)
	}

	v, found = k.GetAsset(ctx, denom)
	if !found {
		return v, types.NewErrorAssetNotFound(denom)
	}

	return v, nil
}
