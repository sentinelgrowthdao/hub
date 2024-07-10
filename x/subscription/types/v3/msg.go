package v3

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
	_ sdk.Msg = (*MsgRenewRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgStartRequest(from sdk.AccAddress, id uint64, denom string, renewable bool) *MsgStartRequest {
	return &MsgStartRequest{
		From:      from.String(),
		ID:        id,
		Denom:     denom,
		Renewable: renewable,
	}
}

func (m *MsgStartRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
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
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateRequest(from sdk.AccAddress, id uint64, renewable bool) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:      from.String(),
		ID:        id,
		Renewable: renewable,
	}
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRenewRequest(from sdk.AccAddress, id uint64, denom string) *MsgRenewRequest {
	return &MsgRenewRequest{
		From:  from.String(),
		ID:    id,
		Denom: denom,
	}
}

func (m *MsgRenewRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
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
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: nodeAddr.String(),
	}
}

func (m *MsgStartSessionRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgStartSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateParamsRequest(from sdk.AccAddress, params v2.Params) *MsgUpdateParamsRequest {
	return &MsgUpdateParamsRequest{
		From:   from.String(),
		Params: params,
	}
}

func (m *MsgUpdateParamsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if err := m.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (m *MsgUpdateParamsRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
