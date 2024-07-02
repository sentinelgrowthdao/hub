package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) HandleMsgStart(ctx sdk.Context, msg *v3.MsgStartRequest) (*v3.MsgStartResponse, error) {
	return &v3.MsgStartResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateDetails(ctx sdk.Context, msg *v3.MsgUpdateDetailsRequest) (*v3.MsgUpdateDetailsResponse, error) {
	return &v3.MsgUpdateDetailsResponse{}, nil
}

func (k *Keeper) HandleMsgRenew(ctx sdk.Context, msg *v3.MsgRenewRequest) (*v3.MsgRenewResponse, error) {
	return &v3.MsgRenewResponse{}, nil
}

func (k *Keeper) HandleMsgCancel(ctx sdk.Context, msg *v2.MsgCancelRequest) (*v2.MsgCancelResponse, error) {
	return &v2.MsgCancelResponse{}, nil
}

func (k *Keeper) HandleMsgAllocate(ctx sdk.Context, msg *v2.MsgAllocateRequest) (*v2.MsgAllocateResponse, error) {
	return &v2.MsgAllocateResponse{}, nil
}

func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	return &v3.MsgStartSessionResponse{}, nil
}
