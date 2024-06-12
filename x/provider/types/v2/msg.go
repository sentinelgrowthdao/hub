package v2

import (
	"net/url"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/provider/types"
)

// The `types` package contains custom message types for the Cosmos SDK.

// The following variables implement the sdk.Msg interface for MsgRegisterRequest and MsgUpdateRequest.
// These variables ensure that the corresponding types can be used as messages in the Cosmos SDK.
var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
)

// NewMsgRegisterRequest creates a new MsgRegisterRequest instance with the given parameters.
func NewMsgRegisterRequest(from sdk.AccAddress, name, identity, website, description string) *MsgRegisterRequest {
	return &MsgRegisterRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

// ValidateBasic performs basic validation checks on the MsgRegisterRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'Name' field is not empty and its length is not greater than 64 characters,
// if the 'Identity' field's length is not greater than 64 characters,
// if the 'Website' field's length is not greater than 64 characters (if not empty),
// if the 'Website' field represents a valid URL (if not empty),
// and if the 'Description' field's length is not greater than 256 characters.
func (m *MsgRegisterRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.Name == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "name cannot be empty")
	}
	if len(m.Name) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
		}
	}
	if len(m.Description) > 256 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "description length cannot be greater than %d chars", 256)
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgRegisterRequest.
func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgUpdateRequest creates a new MsgUpdateRequest instance with the given parameters.
func NewMsgUpdateRequest(from base.ProvAddress, name, identity, website, description string, status base.Status) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
		Status:      status,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateRequest fields.
// It checks if the 'From' field is not empty and represents a valid provider address,
// if the 'Name' field's length is not greater than 64 characters,
// if the 'Identity' field's length is not greater than 64 characters,
// if the 'Website' field's length is not greater than 64 characters (if not empty),
// if the 'Website' field represents a valid URL (if not empty),
// if the 'Description' field's length is not greater than 256 characters,
// and if the 'Status' field is one of the allowed values [unspecified, active, inactive].
func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.ProvAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if len(m.Name) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
		}
	}
	if len(m.Description) > 256 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "description length cannot be greater than %d chars", 256)
	}
	if !m.Status.IsOneOf(base.StatusUnspecified, base.StatusActive, base.StatusInactive) {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "status must be one of [unspecified, active, inactive]")
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgUpdateRequest.
func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := base.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
