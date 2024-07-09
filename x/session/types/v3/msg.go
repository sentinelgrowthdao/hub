package v3

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
)

var (
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
)

func NewMsgEndRequest(from sdk.AccAddress, id uint64) *MsgEndRequest {
	return &MsgEndRequest{
		From: from.String(),
		ID:   id,
	}
}

func (m *MsgEndRequest) ValidateBasic() error {
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

func (m *MsgEndRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateDetailsRequest(
	from base.NodeAddress, id uint64, downloadBytes, uploadBytes sdkmath.Int, duration time.Duration, signature []byte,
) *MsgUpdateDetailsRequest {
	return &MsgUpdateDetailsRequest{
		From:          from.String(),
		ID:            id,
		DownloadBytes: downloadBytes,
		UploadBytes:   uploadBytes,
		Duration:      duration,
		Signature:     signature,
	}
}

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
