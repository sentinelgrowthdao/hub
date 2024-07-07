package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
)

func (k *Keeper) HandleMsgEnd(ctx sdk.Context, msg *v2.MsgEndRequest) (*v2.MsgEndResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}
	if !session.GetStatus().Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.GetID(), session.GetStatus())
	}

	accAddr := session.GetAccAddress()
	if !fromAddr.Equals(accAddr) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	statusChangeDelay := k.StatusChangeDelay(ctx)
	k.DeleteSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

	session.SetStatus(v1base.StatusInactivePending)
	session.SetInactiveAt(ctx.BlockTime().Add(statusChangeDelay))
	session.SetStatusAt(ctx.BlockTime())

	k.SetSession(ctx, session)
	k.SetSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

	return &v2.MsgEndResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateDetails(ctx sdk.Context, msg *v3.MsgUpdateDetailsRequest) (*v3.MsgUpdateDetailsResponse, error) {
	fromAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}
	if session.GetStatus().Equal(v1base.StatusInactive) {
		return nil, types.NewErrorInvalidSessionStatus(session.GetID(), session.GetStatus())
	}

	nodeAddr := session.GetNodeAddress()
	if !fromAddr.Equals(nodeAddr) {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if k.ProofVerificationEnabled(ctx) {
		accAddr := session.GetAccAddress()
		if err := k.VerifySignature(ctx, accAddr, msg.Proof(), msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	session.SetDownloadBytes(msg.DownloadBytes)
	session.SetUploadBytes(msg.UploadBytes)
	session.SetDuration(msg.Duration)

	k.SetSession(ctx, session)

	if session.GetStatus().Equal(v1base.StatusActive) {
		k.DeleteSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())

		statusChangeDelay := k.StatusChangeDelay(ctx)
		session.SetInactiveAt(ctx.BlockTime().Add(statusChangeDelay))
		k.SetSessionForInactiveAt(ctx, session.GetInactiveAt(), session.GetID())
	}

	return &v3.MsgUpdateDetailsResponse{}, nil
}
