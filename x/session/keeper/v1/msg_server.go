package v1

import (
	"context"

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

func (k *msgServer) MsgStart(c context.Context, req *v1.MsgStartRequest) (*v1.MsgStartResponse, error) {
	return &v1.MsgStartResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, req *v1.MsgUpdateRequest) (*v1.MsgUpdateResponse, error) {
	return &v1.MsgUpdateResponse{}, nil
}

func (k *msgServer) MsgEnd(c context.Context, req *v1.MsgEndRequest) (*v1.MsgEndResponse, error) {
	return &v1.MsgEndResponse{}, nil
}
