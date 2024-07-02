package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v1"
)

var (
	_ v1.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	keeper.Keeper
}

func NewMsgServiceServer(k keeper.Keeper) v1.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgSubscribeToNode(_ context.Context, _ *v1.MsgSubscribeToNodeRequest) (*v1.MsgSubscribeToNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgSubscribeToPlan(_ context.Context, _ *v1.MsgSubscribeToPlanRequest) (*v1.MsgSubscribeToPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgCancel(_ context.Context, _ *v1.MsgCancelRequest) (*v1.MsgCancelResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgAddQuota(_ context.Context, _ *v1.MsgAddQuotaRequest) (*v1.MsgAddQuotaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgUpdateQuota(_ context.Context, _ *v1.MsgUpdateQuotaRequest) (*v1.MsgUpdateQuotaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
