package types

import (
	sdkerrors "cosmossdk.io/errors"

	base "github.com/sentinel-official/hub/v12/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateProvider = sdkerrors.Register(ModuleName, 201, "duplicate provider")
	ErrorProviderNotFound  = sdkerrors.Register(ModuleName, 202, "provider not found")
	ErrorUnauthorized      = sdkerrors.Register(ModuleName, 203, "unauthorized")
)

func NewErrorDuplicateProvider(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrorDuplicateProvider, "provider %s already exist", addr)
}

func NewErrorProviderNotFound(addr base.ProvAddress) error {
	return sdkerrors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorUnauthorized(addr string) error {
	return sdkerrors.Wrapf(ErrorUnauthorized, "address %s is not authorized", addr)
}
