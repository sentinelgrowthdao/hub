package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types/v1"
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

func (k *msgServer) MsgRegister(_ context.Context, _ *v1.MsgRegisterRequest) (*v1.MsgRegisterResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgUpdate(_ context.Context, _ *v1.MsgUpdateRequest) (*v1.MsgUpdateResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgSetStatus(_ context.Context, _ *v1.MsgSetStatusRequest) (*v1.MsgSetStatusResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
