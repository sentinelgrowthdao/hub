package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcicqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v7/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
)

// SendQueryPacket serializes query requests and sends them as an IBC packet to a destination chain.
func (k *Keeper) SendQueryPacket(
	ctx sdk.Context,
	channelCap *capabilitytypes.Capability,
	sourcePortID, sourceChannelID string,
	timeoutTimestamp uint64,
	reqs []abcitypes.RequestQuery,
) (uint64, error) {
	// Serialize the Cosmos query requests into binary format.
	data, err := ibcicqtypes.SerializeCosmosQuery(reqs)
	if err != nil {
		return 0, err
	}

	// Create packet data with the serialized queries and validate it.
	packetData := ibcicqtypes.InterchainQueryPacketData{Data: data}
	if err := packetData.ValidateBasic(); err != nil {
		return 0, err
	}

	// Use the ICS-04 interface to send the packet over IBC.
	return k.ics4.SendPacket(
		ctx,
		channelCap,
		sourcePortID,
		sourceChannelID,
		ibcclienttypes.ZeroHeight(),
		timeoutTimestamp,
		packetData.GetBytes(),
	)
}

func (k *Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet ibcchanneltypes.Packet,
	ack ibcchanneltypes.Acknowledgement,
) error {
	return nil
}
