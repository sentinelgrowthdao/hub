package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

var (
	_ sessiontypes.Session = (*Session)(nil)
)

func (m *Session) GetAccAddress() sdk.AccAddress {
	if m.AccAddress == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(m.AccAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Session) GetNodeAddress() base.NodeAddress {
	if m.NodeAddress == "" {
		return nil
	}

	addr, err := base.NodeAddressFromBech32(m.NodeAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (m *Session) GetID() uint64            { return m.ID }
func (m *Session) GetInactiveAt() time.Time { return m.InactiveAt }
func (m *Session) GetStatus() v1base.Status { return m.Status }
func (m *Session) GetType()                 {}

func (m *Session) MsgEndRequest() *sessiontypes.MsgEndRequest {
	return &sessiontypes.MsgEndRequest{
		From: m.AccAddress,
		ID:   m.ID,
	}
}

func (m *Session) SetDownloadBytes(v sdkmath.Int) { m.DownloadBytes = v }
func (m *Session) SetDuration(v time.Duration)    { m.Duration = v }
func (m *Session) SetInactiveAt(v time.Time)      { m.InactiveAt = v }
func (m *Session) SetStatusAt(v time.Time)        { m.StatusAt = v }
func (m *Session) SetStatus(v v1base.Status)      { m.Status = v }
func (m *Session) SetUploadBytes(v sdkmath.Int)   { m.UploadBytes = v }
