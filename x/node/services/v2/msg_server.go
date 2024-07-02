package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

var (
	_ v2.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v2.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgRegister(c context.Context, msg *v2.MsgRegisterRequest) (*v2.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgRegister(ctx, msg)
}

func (k *msgServer) MsgUpdateDetails(c context.Context, msg *v2.MsgUpdateDetailsRequest) (*v2.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateDetails(ctx, msg)
}

func (k *msgServer) MsgUpdateStatus(c context.Context, msg *v2.MsgUpdateStatusRequest) (*v2.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateStatus(ctx, msg)
}

func (k *msgServer) MsgSubscribe(_ context.Context, _ *v2.MsgSubscribeRequest) (*v2.MsgSubscribeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
