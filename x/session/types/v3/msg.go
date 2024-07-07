package v3

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
)

var (
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
)

func (m *MsgUpdateDetailsRequest) Proof() *Proof {
	return &Proof{
		ID:            m.ID,
		DownloadBytes: m.DownloadBytes,
		UploadBytes:   m.UploadBytes,
		Duration:      m.Duration,
	}
}

func (m *MsgUpdateDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "id cannot be zero")
	}
	if m.DownloadBytes.IsNil() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "download_bytes cannot be nil")
	}
	if !m.DownloadBytes.IsPositive() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "download_bytes must be positive")
	}
	if m.UploadBytes.IsNil() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "upload_bytes cannot be nil")
	}
	if !m.UploadBytes.IsPositive() {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "upload_bytes must be positive")
	}
	if m.Duration < 0 {
		return sdkerrors.Wrap(types.ErrorInvalidMessage, "duration cannot be negative")
	}
	if m.Signature != nil && len(m.Signature) != 64 {
		return sdkerrors.Wrapf(types.ErrorInvalidMessage, "signature length must be %d bytes", 64)
	}

	return nil
}

func (m *MsgUpdateDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
