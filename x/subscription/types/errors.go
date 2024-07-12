package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorAllocationNotFound        = sdkerrors.Register(ModuleName, 201, "allocation not found")
	ErrorInsufficientBytes         = sdkerrors.Register(ModuleName, 202, "insufficient bytes")
	ErrorInvalidAllocation         = sdkerrors.Register(ModuleName, 203, "invalid allocation")
	ErrorInvalidNodeStatus         = sdkerrors.Register(ModuleName, 204, "invalid node status")
	ErrorInvalidPlanStatus         = sdkerrors.Register(ModuleName, 205, "invalid plan status")
	ErrorInvalidSubscriptionStatus = sdkerrors.Register(ModuleName, 206, "invalid subscription status")
	ErrorNodeNotFound              = sdkerrors.Register(ModuleName, 207, "node not found")
	ErrorPlanNotFound              = sdkerrors.Register(ModuleName, 208, "plan not found")
	ErrorPriceNotFound             = sdkerrors.Register(ModuleName, 209, "price not found")
	ErrorSubscriptionNotFound      = sdkerrors.Register(ModuleName, 210, "subscription not found")
	ErrorUnauthorized              = sdkerrors.Register(ModuleName, 211, "unauthorized")
)

func NewErrorAllocationNotFound(id uint64, addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrorAllocationNotFound, "allocation %d/%s does not exist", id, addr)
}

func NewErrorInsufficientBytes(id uint64, bytes sdkmath.Int) error {
	return sdkerrors.Wrapf(ErrorInsufficientBytes, "insufficient bytes %s for subscription %d", bytes, id)
}

func NewErrorInvalidAllocation(id uint64, addr sdk.AccAddress) error {
	return sdkerrors.Wrapf(ErrorInvalidAllocation, "invalid allocation %d/%s", id, addr)
}

func NewErrorInvalidNodeStatus(addr base.NodeAddress, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidNodeStatus, "invalid status %s for node %d", status, addr)
}

func NewErrorInvalidPlanStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidPlanStatus, "invalid status %s for plan %d", status, id)
}

func NewErrorInvalidSubscriptionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidSubscriptionStatus, "invalid status %s for subscription %d", status, id)
}

func NewErrorNodeNotFound(addr base.NodeAddress) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorPriceNotFound(denom string) error {
	return sdkerrors.Wrapf(ErrorPriceNotFound, "price for denom %s does not exist", denom)
}

func NewErrorSubscriptionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorSubscriptionNotFound, "subscription %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
