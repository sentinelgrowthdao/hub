package v2

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
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

func (k *msgServer) MsgStart(_ context.Context, _ *v2.MsgStartRequest) (*v2.MsgStartResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgUpdateDetails(_ context.Context, _ *v2.MsgUpdateDetailsRequest) (*v2.MsgUpdateDetailsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgEnd(_ context.Context, _ *v2.MsgEndRequest) (*v2.MsgEndResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
