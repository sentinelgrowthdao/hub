package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateLease   = sdkerrors.Register(ModuleName, 201, "duplicate lease")
	ErrorInvalidHours     = sdkerrors.Register(ModuleName, 202, "invalid hours")
	ErrorLeaseNotFound    = sdkerrors.Register(ModuleName, 203, "lease not found")
	ErrorNodeNotFound     = sdkerrors.Register(ModuleName, 204, "node not found")
	ErrorPriceNotFound    = sdkerrors.Register(ModuleName, 205, "price not found")
	ErrorProviderNotFound = sdkerrors.Register(ModuleName, 206, "provider not found")
	ErrorUnauthorized     = sdkerrors.Register(ModuleName, 207, "unauthorized")
)

func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrorInvalidHours, "invalid hours %d", hours)
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

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
