package v3

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
)

var (
	_ sdk.Msg = (*MsgCreatePlanRequest)(nil)
	_ sdk.Msg = (*MsgLinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUnlinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUpdatePlanStatusRequest)(nil)
	_ sdk.Msg = (*MsgStartSessionRequest)(nil)
)

func NewMsgCreatePlanRequest(from base.ProvAddress, duration time.Duration, gigabytes int64, prices sdk.Coins) *MsgCreatePlanRequest {
	return &MsgCreatePlanRequest{
		From:      from.String(),
		Duration:  duration,
		Gigabytes: gigabytes,
		Prices:    prices,
	}
}

func (m *MsgCreatePlanRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Duration < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "duration cannot be negative")
	}
	if m.Duration == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "duration cannot be zero")
	}
	if m.Gigabytes < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabytes cannot be negative")
	}
	if m.Gigabytes == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "gigabytes cannot be zero")
	}
	if m.Prices == nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "prices cannot be nil")
	}
	if m.Prices.Len() == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "prices cannot be empty")
	}
	if m.Prices.IsAnyNil() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "prices cannot contain nil")
	}
	if !m.Prices.IsValid() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "prices must be valid")
	}

	return nil
}

func (m *MsgCreatePlanRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgLinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgLinkNodeRequest {
	return &MsgLinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

func (m *MsgLinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
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

func (m *MsgLinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUnlinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgUnlinkNodeRequest {
	return &MsgUnlinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

func (m *MsgUnlinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
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

func (m *MsgUnlinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdatePlanStatusRequest(from base.ProvAddress, id uint64, status v1base.Status) *MsgUpdatePlanStatusRequest {
	return &MsgUpdatePlanStatusRequest{
		From:   from.String(),
		ID:     id,
		Status: status,
	}
}

func (m *MsgUpdatePlanStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if !m.Status.IsOneOf(v1base.StatusActive, v1base.StatusInactive) {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgUpdatePlanStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgStartSessionRequest(from sdk.AccAddress, id uint64, denom string, renewable bool, nodeAddr base.NodeAddress) *MsgStartSessionRequest {
	return &MsgStartSessionRequest{
		From:        from.String(),
		ID:          id,
		Denom:       denom,
		Renewable:   renewable,
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
	if m.Denom == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
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
