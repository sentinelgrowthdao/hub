// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	ibcporttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	ibcerrors "github.com/cosmos/ibc-go/v7/modules/core/errors"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorAssetNotFound      = sdkerrors.Register(ModuleName, 201, "asset not found")
	ErrorAssetPriceNotFound = sdkerrors.Register(ModuleName, 202, "asset price not found")
)

func NewErrorInvalidVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid version %s; expected %s", version, expected)
}

func NewErrorInvalidCounterpartyVersion(version, expected string) error {
	return sdkerrors.Wrapf(ibcerrors.ErrInvalidVersion, "invalid counteryparty version %s; expected %s", version, expected)
}

func NewErrorInvalidChannelOrdering(order, expected ibcchanneltypes.Order) error {
	return sdkerrors.Wrapf(ibcchanneltypes.ErrInvalidChannelOrdering, "invalid channel order %s; expected %s", order, expected)
}

func NewErrorInvalidPort(portID, expected string) error {
	return sdkerrors.Wrapf(ibcporttypes.ErrInvalidPort, "invalid port %s; expected %s", portID, expected)
}

func NewErrorAssetNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorAssetNotFound, "asset %s does not exist", denom)
}

func NewErrorAssetPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorAssetPriceNotFound, "price for asset %s does not exist", denom)
}
