package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateNode     = sdkerrors.Register(ModuleName, 201, "duplicate node")
	ErrorInvalidGigabytes  = sdkerrors.Register(ModuleName, 202, "invalid gigabytes")
	ErrorInvalidHours      = sdkerrors.Register(ModuleName, 203, "invalid hours")
	ErrorInvalidNodeStatus = sdkerrors.Register(ModuleName, 205, "invalid node status")
	ErrorInvalidPrices     = sdkerrors.Register(ModuleName, 204, "invalid prices")
	ErrorNodeNotFound      = sdkerrors.Register(ModuleName, 206, "node not found")
	ErrorPriceNotFound     = sdkerrors.Register(ModuleName, 207, "price not found")
	ErrorUnauthorized      = sdkerrors.Register(ModuleName, 208, "unauthorized")
)

func NewErrorDuplicateNode(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorDuplicateNode, "node %s already exist", addr)
}

func NewErrorInvalidGigabytes(gigabytes int64) error {
	return sdkerrors.Wrapf(ErrorInvalidGigabytes, "invalid gigabytes %d", gigabytes)
}

func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrorInvalidHours, "invalid hours %d", hours)
}

func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s", status, addr)
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

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
