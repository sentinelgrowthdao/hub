package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

func (k *Keeper) HandleMsgUpdateDetails(ctx sdk.Context, msg *v2.MsgUpdateDetailsRequest) (*v2.MsgUpdateDetailsResponse, error) {
	session, found := k.GetSession(ctx, msg.Proof.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.Proof.ID)
	}
	if session.Status.Equal(v1base.StatusInactive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	if msg.From != session.NodeAddress {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	if k.ProofVerificationEnabled(ctx) {
		if err := k.VerifySignature(ctx, session.GetAddress(), msg.Proof, msg.Signature); err != nil {
			return nil, types.NewErrorInvalidSignature(msg.Signature)
		}
	}

	if session.Status.Equal(v1base.StatusActive) {
		statusChangeDelay := k.StatusChangeDelay(ctx)
		k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

		session.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)
		k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)
	}

	session.Bandwidth = msg.Proof.Bandwidth
	session.Duration = msg.Proof.Duration

	k.SetSession(ctx, session)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateDetails{
			Address:        session.Address,
			NodeAddress:    session.NodeAddress,
			ID:             session.ID,
			PlanID:         0,
			SubscriptionID: session.SubscriptionID,
		},
	)

	return &v2.MsgUpdateDetailsResponse{}, nil
}

func (k *Keeper) HandleMsgEnd(ctx sdk.Context, msg *v2.MsgEndRequest) (*v2.MsgEndResponse, error) {
	session, found := k.GetSession(ctx, msg.ID)
	if !found {
		return nil, types.NewErrorSessionNotFound(msg.ID)
	}
	if !session.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidSessionStatus(session.ID, session.Status)
	}

	if msg.From != session.Address {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	statusChangeDelay := k.StatusChangeDelay(ctx)
	k.DeleteSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	session.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)
	session.Status = v1base.StatusInactivePending
	session.StatusAt = ctx.BlockTime()

	k.SetSession(ctx, session)
	k.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateStatus{
			Status:         v1base.StatusInactivePending,
			Address:        session.Address,
			NodeAddress:    session.NodeAddress,
			ID:             session.ID,
			PlanID:         0,
			SubscriptionID: session.SubscriptionID,
		},
	)

	return &v2.MsgEndResponse{}, nil
}
