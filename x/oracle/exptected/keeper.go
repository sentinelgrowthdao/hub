package exptected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, portID, channelID string) (ibcchanneltypes.Channel, bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
}
