package v1

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types"
)

var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
	_ sdk.Msg = (*MsgRenewRequest)(nil)
	_ sdk.Msg = (*MsgEndRequest)(nil)
)

func NewMsgStartRequest(fromAddr base.ProvAddress, nodeAddr base.NodeAddress, hours int64, denom string, renewable bool) *MsgStartRequest {
	return &MsgStartRequest{
		From:        fromAddr.String(),
		NodeAddress: nodeAddr.String(),
		Hours:       hours,
		Denom:       denom,
		Renewable:   renewable,
	}
}

func (m *MsgStartRequest) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

func (m *MsgStartRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Hours == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hours cannot be empty")
	}
	if m.Hours < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hours cannot be negative")
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateDetailsRequest(fromAddr base.ProvAddress, id uint64, renewable bool) *MsgUpdateDetailsRequest {
	return &MsgUpdateDetailsRequest{
		From:      fromAddr.String(),
		ID:        id,
		Renewable: renewable,
	}
}

func (m *MsgUpdateDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgUpdateDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRenewRequest(fromAddr base.ProvAddress, id uint64, hours int64, denom string) *MsgRenewRequest {
	return &MsgRenewRequest{
		From:  fromAddr.String(),
		ID:    id,
		Hours: hours,
		Denom: denom,
	}
}

func (m *MsgRenewRequest) GetHours() time.Duration {
	return time.Duration(m.Hours) * time.Hour
}

func (m *MsgRenewRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Hours == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hours cannot be empty")
	}
	if m.Hours < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "hours cannot be negative")
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgRenewRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgEndRequest(fromAddr base.ProvAddress, id uint64) *MsgEndRequest {
	return &MsgEndRequest{
		From: fromAddr.String(),
		ID:   id,
	}
}

func (m *MsgEndRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgEndRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
