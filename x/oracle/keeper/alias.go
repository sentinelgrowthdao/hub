package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

func (k *Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.capability.AuthenticateCapability(ctx, cap, name)
}

func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.capability.ClaimCapability(ctx, cap, name)
}

func (k *Keeper) GetCapability(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool) {
	return k.capability.GetCapability(ctx, name)
}

func (k *Keeper) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool) {
	return k.channel.GetNextSequenceSend(ctx, portID, channelID)
}

func (k *Keeper) SendPacket(
	ctx sdk.Context, chanCap *capabilitytypes.Capability, sourcePort string, sourceChannel string,
	timeoutHeight ibcclienttypes.Height, timeoutTimestamp uint64, data []byte,
) (uint64, error) {
	return k.ics4.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}
