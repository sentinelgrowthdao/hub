package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	channelID := k.GetChannelID(ctx)
	portID := k.GetPortID(ctx)
	timeout := k.GetTimeout(ctx)

	channelCap, found := k.GetCapability(ctx, ibchost.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil
	}

	k.IterateAssets(ctx, func(_ int, asset types.Asset) bool {
		sequence, err := k.SendQueryPacket(
			ctx, channelCap, portID, channelID, uint64(timeout),
			asset.QueryGetProtoRevPoolRequest(k.cdc),
			asset.SpotPriceRequest(k.cdc),
		)
		if err != nil {
			panic(err)
		}

		k.SetDenomForPacket(ctx, portID, channelID, sequence, asset.Denom)
		return false
	})

	return nil
}
