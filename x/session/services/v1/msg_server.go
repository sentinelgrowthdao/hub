package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v1"
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

func (k *msgServer) MsgStart(_ context.Context, _ *v1.MsgStartRequest) (*v1.MsgStartResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgUpdate(_ context.Context, _ *v1.MsgUpdateRequest) (*v1.MsgUpdateResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgEnd(_ context.Context, _ *v1.MsgEndRequest) (*v1.MsgEndResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
