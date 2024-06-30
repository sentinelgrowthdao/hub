package v3

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/node/types"
)

var (
	_ sdk.Msg = (*MsgStartLeaseRequest)(nil)
	_ sdk.Msg = (*MsgUpdateLeaseDetailsRequest)(nil)
	_ sdk.Msg = (*MsgRenewLeaseRequest)(nil)
	_ sdk.Msg = (*MsgEndLeaseRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
)

func NewMsgStartLeaseRequest(fromAddr base.ProvAddress, nodeAddr base.NodeAddress, hours int64, denom string, renewable bool) *MsgStartLeaseRequest {
	return &MsgStartLeaseRequest{
		From:        fromAddr.String(),
		NodeAddress: nodeAddr.String(),
		Hours:       hours,
		Denom:       denom,
		Renewable:   renewable,
	}
}

func (m *MsgStartLeaseRequest) ValidateBasic() error {
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

func (m *MsgStartLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateLeaseDetailsRequest(fromAddr base.ProvAddress, id uint64, renewable bool) *MsgUpdateLeaseDetailsRequest {
	return &MsgUpdateLeaseDetailsRequest{
		From:      fromAddr.String(),
		ID:        id,
		Renewable: renewable,
	}
}

func (m *MsgUpdateLeaseDetailsRequest) ValidateBasic() error {
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

func (m *MsgUpdateLeaseDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRenewLeaseRequest(fromAddr base.ProvAddress, id uint64, hours int64, denom string, renewable bool) *MsgRenewLeaseRequest {
	return &MsgRenewLeaseRequest{
		From:      fromAddr.String(),
		ID:        id,
		Hours:     hours,
		Denom:     denom,
		Renewable: renewable,
	}
}

func (m *MsgRenewLeaseRequest) ValidateBasic() error {
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

func (m *MsgRenewLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgEndLeaseRequest(fromAddr base.ProvAddress, id uint64) *MsgEndLeaseRequest {
	return &MsgEndLeaseRequest{
		From: fromAddr.String(),
		ID:   id,
	}
}

func (m *MsgEndLeaseRequest) ValidateBasic() error {
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

func (m *MsgEndLeaseRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(fromAddr sdk.AccAddress, nodeAddr base.NodeAddress, gigabytes, hours int64, denom string) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        fromAddr.String(),
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
