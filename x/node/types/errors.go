package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateNode = sdkerrors.Register(ModuleName, 201, "duplicate node")
	ErrorInvalidPrices = sdkerrors.Register(ModuleName, 202, "invalid prices")
	ErrorNodeNotFound  = sdkerrors.Register(ModuleName, 203, "node not found")
	ErrorPriceNotFound = sdkerrors.Register(ModuleName, 204, "price not found")
	ErrorInvalidStatus = sdkerrors.Register(ModuleName, 205, "invalid status")
)

func NewErrorDuplicateNode(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorDuplicateNode, "node %s already exist", addr)
}

func NewErrorInvalidPrices(prices sdk.Coins) error {
	return sdkerrors.Wrapf(ErrorInvalidPrices, "invalid prices %s", prices)
}

func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorPriceNotFound, "price for denom %s does not exist", denom)
}

func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidStatus, "invalid status %s for node %s", status, addr)
}
