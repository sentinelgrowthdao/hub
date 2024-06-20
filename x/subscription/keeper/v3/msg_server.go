package v3

import (
	"context"

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

func (k *msgServer) MsgStartSession(_ context.Context, _ *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
