package v1

import (
	"context"

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

func (k *msgServer) MsgSubscribeToNode(c context.Context, req *v1.MsgSubscribeToNodeRequest) (*v1.MsgSubscribeToNodeResponse, error) {
	return &v1.MsgSubscribeToNodeResponse{}, nil
}

func (k *msgServer) MsgSubscribeToPlan(c context.Context, req *v1.MsgSubscribeToPlanRequest) (*v1.MsgSubscribeToPlanResponse, error) {
	return &v1.MsgSubscribeToPlanResponse{}, nil
}

func (k *msgServer) MsgCancel(c context.Context, req *v1.MsgCancelRequest) (*v1.MsgCancelResponse, error) {
	return &v1.MsgCancelResponse{}, nil
}

func (k *msgServer) MsgAddQuota(c context.Context, req *v1.MsgAddQuotaRequest) (*v1.MsgAddQuotaResponse, error) {
	return &v1.MsgAddQuotaResponse{}, nil
}

func (k *msgServer) MsgUpdateQuota(c context.Context, req *v1.MsgUpdateQuotaRequest) (*v1.MsgUpdateQuotaResponse, error) {
	return &v1.MsgUpdateQuotaResponse{}, nil
}
