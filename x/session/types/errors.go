package types

import (
	sdkerrors "cosmossdk.io/errors"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorInvalidSessionStatus = sdkerrors.Register(ModuleName, 201, "invalid session status")
	ErrorInvalidSignature     = sdkerrors.Register(ModuleName, 202, "invalid signature")
	ErrorSessionNotFound      = sdkerrors.Register(ModuleName, 203, "session not found")
	ErrorUnauthorized         = sdkerrors.Register(ModuleName, 204, "unauthorized")
)

func NewErrorInvalidSessionStatus(id uint64, status v1base.Status) error {
	return sdkerrors.Wrapf(ErrorInvalidSessionStatus, "invalid status %s for session %d", status, id)
}

func NewErrorInvalidSignature(signature []byte) error {
	return sdkerrors.Wrapf(ErrorInvalidSignature, "invalid signature %X", signature)
}

func NewErrorSessionNotFound(id uint64) error {
	return sdkerrors.Wrapf(ErrorSessionNotFound, "session %d does not exist", id)
}

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
