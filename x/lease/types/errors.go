package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateLease        = sdkerrors.Register(ModuleName, 201, "duplicate lease")
	ErrorInvalidHours          = sdkerrors.Register(ModuleName, 202, "invalid hours")
	ErrorInvalidNodeStatus     = sdkerrors.Register(ModuleName, 203, "invalid node status")
	ErrorInvalidProviderStatus = sdkerrors.Register(ModuleName, 204, "invalid provider status")
	ErrorLeaseNotFound         = sdkerrors.Register(ModuleName, 205, "lease not found")
	ErrorNodeNotFound          = sdkerrors.Register(ModuleName, 206, "node not found")
	ErrorPriceNotFound         = sdkerrors.Register(ModuleName, 207, "price not found")
	ErrorProviderNotFound      = sdkerrors.Register(ModuleName, 208, "provider not found")
	ErrorUnauthorized          = sdkerrors.Register(ModuleName, 209, "unauthorized")
)

// NewErrorDuplicateLease returns an error indicating that a lease for the specified provider and node already exists.
func NewErrorDuplicateLease(provAddr base.ProvAddress, nodeAddr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorDuplicateLease, "lease for provider %s by node %s already exists", provAddr, nodeAddr)
}

// NewErrorInvalidHours returns an error indicating that the provided hours are invalid.
func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrorInvalidHours, "invalid hours %d", hours)
}

// NewErrorInvalidNodeStatus returns an error indicating that the provided status is invalid for the given node.
func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %s", status, addr)
}

// NewErrorInvalidProviderStatus returns an error indicating that the provided status is invalid for the given provider.
func NewErrorInvalidProviderStatus(addr base.ProvAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidProviderStatus, "invalid status %s for provider %s", status, addr)
}

// NewErrorLeaseNotFound returns an error indicating that the specified lease does not exist.
func NewErrorLeaseNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorLeaseNotFound, "lease %d does not exist", id)
}

// NewErrorNodeNotFound returns an error indicating that the specified node does not exist.
func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

// NewErrorPriceNotFound returns an error indicating that the price for the specified denom does not exist.
func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorPriceNotFound, "price for denom %s does not exist", denom)
}

// NewErrorProviderNotFound returns an error indicating that the specified provider does not exist.
func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

// NewErrorUnauthorized returns an error indicating that the specified address is not authorized.
func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
