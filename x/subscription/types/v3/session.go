package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func (m *Session) GetDownloadBytes() sdkmath.Int { return m.DownloadBytes }
func (m *Session) GetUploadBytes() sdkmath.Int   { return m.UploadBytes }

func (m *Session) SetDownloadBytes(v sdkmath.Int) { m.DownloadBytes = v }
func (m *Session) SetDuration(v time.Duration)    { m.Duration = v }
func (m *Session) SetInactiveAt(v time.Time)      { m.InactiveAt = v }
func (m *Session) SetStatusAt(v time.Time)        { m.StatusAt = v }
func (m *Session) SetStatus(v v1base.Status)      { m.Status = v }
func (m *Session) SetUploadBytes(v sdkmath.Int)   { m.UploadBytes = v }
