package v2

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/plan/types"
)

// The `types` package contains custom message types for the Cosmos SDK.

// The following variables implement the sdk.Msg interface for the provided message types.
// These variables ensure that the corresponding types can be used as messages in the Cosmos SDK.
var (
	_ sdk.Msg = (*MsgCreateRequest)(nil)
	_ sdk.Msg = (*MsgUpdateStatusRequest)(nil)
	_ sdk.Msg = (*MsgLinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUnlinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeRequest)(nil)
)

// NewMsgCreateRequest creates a new MsgCreateRequest instance with the given parameters.
func NewMsgCreateRequest(from base.ProvAddress, duration time.Duration, gigabytes int64, prices sdk.Coins) *MsgCreateRequest {
	return &MsgCreateRequest{
		From:      from.String(),
		Duration:  duration,
		Gigabytes: gigabytes,
		Prices:    prices,
	}
}

// ValidateBasic performs basic validation checks on the MsgCreateRequest fields.
// It checks if the 'From' field is not empty and represents a valid provider address,
// if the 'Duration' field is not negative or zero,
// if the 'Gigabytes' field is not negative or zero,
// and if the 'Prices' field is valid (not empty, not containing nil coins, and having valid coin denominations).
func (m *MsgCreateRequest) ValidateBasic() error {
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

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgCreateRequest.
func (m *MsgCreateRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateStatusRequest creates a new MsgUpdateStatusRequest instance with the given parameters.
func NewMsgUpdateStatusRequest(from base.ProvAddress, id uint64, status v1base.Status) *MsgUpdateStatusRequest {
	return &MsgUpdateStatusRequest{
		From:   from.String(),
		ID:     id,
		Status: status,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateStatusRequest fields.
// It checks if the 'From' field is not empty and represents a valid provider address,
// if the 'ID' field is not zero,
// and if the 'Status' field is one of the allowed values [active, inactive].
func (m *MsgUpdateStatusRequest) ValidateBasic() error {
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

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgUpdateStatusRequest.
func (m *MsgUpdateStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgLinkNodeRequest creates a new MsgLinkNodeRequest instance with the given parameters.
func NewMsgLinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgLinkNodeRequest {
	return &MsgLinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgLinkNodeRequest fields.
// It checks if the 'From' field is not empty and represents a valid provider address,
// if the 'ID' field is not zero,
// and if the 'Address' field is not empty and represents a valid node address.
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

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgLinkNodeRequest.
func (m *MsgLinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUnlinkNodeRequest creates a new MsgUnlinkNodeRequest instance with the given parameters.
func NewMsgUnlinkNodeRequest(from base.ProvAddress, id uint64, addr base.NodeAddress) *MsgUnlinkNodeRequest {
	return &MsgUnlinkNodeRequest{
		From:        from.String(),
		ID:          id,
		NodeAddress: addr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgUnlinkNodeRequest fields.
// It checks if the 'From' field is not empty and represents a valid provider address,
// if the 'ID' field is not zero,
// and if the 'Address' field is not empty and represents a valid node address.
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

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgUnlinkNodeRequest.
func (m *MsgUnlinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgSubscribeRequest creates a new MsgSubscribeRequest instance with the given parameters.
func NewMsgSubscribeRequest(from sdk.AccAddress, id uint64, denom string) *MsgSubscribeRequest {
	return &MsgSubscribeRequest{
		From:  from.String(),
		ID:    id,
		Denom: denom,
	}
}

// ValidateBasic performs basic validation checks on the MsgSubscribeRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'ID' field is not zero,
// and if the 'Denom' field is valid according to the Cosmos SDK's ValidateDenom function.
func (m *MsgSubscribeRequest) ValidateBasic() error {
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

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgSubscribeRequest.
func (m *MsgSubscribeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
