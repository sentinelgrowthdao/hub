package v1

import (
	"context"

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

func (k *msgServer) MsgRegister(c context.Context, req *v1.MsgRegisterRequest) (*v1.MsgRegisterResponse, error) {
	return &v1.MsgRegisterResponse{}, nil
}

func (k *msgServer) MsgUpdate(c context.Context, req *v1.MsgUpdateRequest) (*v1.MsgUpdateResponse, error) {
	return &v1.MsgUpdateResponse{}, nil
}

func (k *msgServer) MsgSetStatus(c context.Context, req *v1.MsgSetStatusRequest) (*v1.MsgSetStatusResponse, error) {
	return &v1.MsgSetStatusResponse{}, nil
}
