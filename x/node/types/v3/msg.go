package v3

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/types"
)

var (
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
)

func NewMsgStartSessionRequest(from sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, denom string) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		NodeAddress: nodeAddr.String(),
		Gigabytes:   gigabytes,
		Hours:       hours,
		Denom:       denom,
	}
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Gigabytes == 0 && m.Hours == 0 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "[gigabytes, hours] cannot be empty")
	}
	if m.Gigabytes != 0 && m.Hours != 0 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "[gigabytes, hours] cannot be non-empty")
	}
	if m.Gigabytes != 0 {
		if m.Gigabytes < 0 {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabytes cannot be negative")
		}
	}
	if m.Hours != 0 {
		if m.Hours < 0 {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, "hours cannot be negative")
		}
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
