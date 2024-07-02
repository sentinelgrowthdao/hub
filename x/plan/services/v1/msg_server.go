package v1

import (
	"context"

	"github.com/sentinel-official/hub/v12/x/plan/keeper"
	"github.com/sentinel-official/hub/v12/x/plan/types/v1"
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

func (k *msgServer) MsgAdd(c context.Context, req *v1.MsgAddRequest) (*v1.MsgAddResponse, error) {
	return &v1.MsgAddResponse{}, nil
}

func (k *msgServer) MsgSetStatus(c context.Context, req *v1.MsgSetStatusRequest) (*v1.MsgSetStatusResponse, error) {
	return &v1.MsgSetStatusResponse{}, nil
}

func (k *msgServer) MsgAddNode(c context.Context, req *v1.MsgAddNodeRequest) (*v1.MsgAddNodeResponse, error) {
	return &v1.MsgAddNodeResponse{}, nil
}

func (k *msgServer) MsgRemoveNode(c context.Context, req *v1.MsgRemoveNodeRequest) (*v1.MsgRemoveNodeResponse, error) {
	return &v1.MsgRemoveNodeResponse{}, nil
}
