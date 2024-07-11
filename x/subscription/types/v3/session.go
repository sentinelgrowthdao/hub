package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v3"
)

var (
	_ sessiontypes.Session = (*Session)(nil)
)

func (m *Session) GetDownloadBytes() sdkmath.Int { return m.DownloadBytes }
func (m *Session) GetUploadBytes() sdkmath.Int   { return m.UploadBytes }

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
