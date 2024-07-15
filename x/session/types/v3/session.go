package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/gogoproto/proto"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

type Session interface {
	proto.Message

	GetAccAddress() string
	GetDownloadBytes() sdkmath.Int
	GetDuration() time.Duration
	GetID() uint64
	GetInactiveAt() time.Time
	GetNodeAddress() string
	GetStatus() v1base.Status
	GetStatusAt() time.Time
	GetUploadBytes() sdkmath.Int

	SetDownloadBytes(sdkmath.Int)
	SetDuration(time.Duration)
	SetInactiveAt(time.Time)
	SetStatusAt(time.Time)
	SetStatus(v1base.Status)
	SetUploadBytes(sdkmath.Int)
}
