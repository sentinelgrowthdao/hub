package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateLease   = sdkerrors.Register(ModuleName, 201, "duplicate lease")
	ErrorDuplicateNode    = sdkerrors.Register(ModuleName, 202, "duplicate node")
	ErrorInvalidGigabytes = sdkerrors.Register(ModuleName, 203, "invalid gigabytes")
	ErrorInvalidHours     = sdkerrors.Register(ModuleName, 204, "invalid hours")
	ErrorInvalidPrices    = sdkerrors.Register(ModuleName, 205, "invalid prices")
	ErrorLeaseNotFound    = sdkerrors.Register(ModuleName, 206, "lease not found")
	ErrorNodeNotFound     = sdkerrors.Register(ModuleName, 207, "node not found")
	ErrorPriceNotFound    = sdkerrors.Register(ModuleName, 208, "price not found")
	ErrorProviderNotFound = sdkerrors.Register(ModuleName, 209, "provider not found")
	ErrorUnauthorised     = sdkerrors.Register(ModuleName, 210, "unauthorised")
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

func NewErrorInvalidPrices(prices sdk.Coins) error {
	return sdkerrors.Wrapf(ErrorInvalidPrices, "invalid prices %s", prices)
}

func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorPriceNotFound, "price for denom %s does not exist", denom)
}

func NewErrorDuplicateLease(provAddr base.ProvAddress, nodeAddr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorDuplicateLease, "lease for provider %s by node %s already exist", provAddr, nodeAddr)
}

func NewErrorLeaseNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorLeaseNotFound, "lease %d does not exist", id)
}

func NewErrorUnauthorised(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorised, "address %s is not authorised", addr)
}
