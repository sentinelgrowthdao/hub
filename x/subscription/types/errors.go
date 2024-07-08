package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorInsufficientBytes = sdkerrors.Register(ModuleName, 201, "insufficient bytes")
	ErrorInvalidAllocation = sdkerrors.Register(ModuleName, 202, "invalid allocation")
	ErrorInvalidStatus     = sdkerrors.Register(ModuleName, 203, "invalid status")
	ErrorNotFound          = sdkerrors.Register(ModuleName, 204, "not found")
	ErrorUnauthorized      = sdkerrors.Register(ModuleName, 205, "unauthorized")
)

func NewErrorAllocationNotFound(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorNotFound, "allocation %d/%s does not exist", id, addr)
}

func NewErrorInsufficientBytes(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInsufficientBytes, "insufficient bytes for allocation %d/%s", id, addr)
}

func NewErrorInvalidAllocation(id uint64, addr interface{}) error {
	return sdkerrors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidPlanStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidStatus, "invalid status %s for plan %d", status, id)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorNotFound, "plan %d does not exist", id)
}

func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorNotFound, "price for denom %s does not exist", denom)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}

func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNotFound, "node %s does not exist", addr)
}

func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidStatus, "invalid status %s for node %d", status, addr)
}
