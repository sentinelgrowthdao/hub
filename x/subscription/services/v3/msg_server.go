package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

var (
	_ v3.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v3.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgStart(c context.Context, msg *v3.MsgStartRequest) (*v3.MsgStartResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgStart(ctx, msg)
}

func (k *msgServer) MsgUpdate(c context.Context, msg *v3.MsgUpdateRequest) (*v3.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdate(ctx, msg)
}

func (k *msgServer) MsgRenew(_ context.Context, _ *v3.MsgRenewRequest) (*v3.MsgRenewResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgStartSession(c context.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgStartSession(ctx, msg)
}

func (k *msgServer) MsgUpdateParams(c context.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateParams(ctx, msg)
}
