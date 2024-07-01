package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateNode = sdkerrors.Register(ModuleName, 201, "duplicate node")
	ErrorInvalidPrices = sdkerrors.Register(ModuleName, 202, "invalid prices")
	ErrorNodeNotFound  = sdkerrors.Register(ModuleName, 203, "node not found")
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
