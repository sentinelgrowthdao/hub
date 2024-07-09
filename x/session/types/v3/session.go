package v3

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

type Session interface {
	proto.Message

	GetAccAddress() sdk.AccAddress
	GetID() uint64
	GetInactiveAt() time.Time
	GetNodeAddress() base.NodeAddress
	GetStatus() v1base.Status
	GetType()

	MsgEndRequest() *MsgEndRequest

	SetDownloadBytes(sdkmath.Int)
	SetDuration(time.Duration)
	SetInactiveAt(time.Time)
	SetStatusAt(time.Time)
	SetStatus(v1base.Status)
	SetUploadBytes(sdkmath.Int)
}
