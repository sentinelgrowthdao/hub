package v3

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

var (
	_ sdk.Msg = (*MsgCancelSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateSessionRequest)(nil)
	_ sdk.Msg = (*MsgUpdateParamsRequest)(nil)
)

func NewMsgCancelSessionRequest(from sdk.AccAddress, id uint64) *MsgCancelSessionRequest {
	return &MsgCancelSessionRequest{
		From: from.String(),
		ID:   id,
	}
}

func (m *MsgCancelSessionRequest) ValidateBasic() error {
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

func (m *MsgCancelSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateSessionRequest(from base.NodeAddress, id uint64, downloadBytes, uploadBytes sdkmath.Int, duration time.Duration, signature []byte) *MsgUpdateSessionRequest {
	return &MsgUpdateSessionRequest{
		From:          from.String(),
		ID:            id,
		DownloadBytes: downloadBytes,
		UploadBytes:   uploadBytes,
		Duration:      duration,
		Signature:     signature,
	}
}

func (m *MsgUpdateSessionRequest) Proof() *Proof {
	return &Proof{
		ID:            m.ID,
		DownloadBytes: m.DownloadBytes,
		UploadBytes:   m.UploadBytes,
		Duration:      m.Duration,
	}
}

func (m *MsgUpdateSessionRequest) ValidateBasic() error {
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

func (m *MsgUpdateSessionRequest) GetSigners() []sdk.AccAddress {
	from, err := base.NodeAddressFromBech32(m.From)
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
