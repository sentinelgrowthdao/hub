package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/plan/keeper"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
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

func (k *msgServer) MsgCreate(c context.Context, msg *v2.MsgCreateRequest) (*v2.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgCreate(ctx, msg)
}

func (k *msgServer) MsgUpdateStatus(c context.Context, msg *v2.MsgUpdateStatusRequest) (*v2.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateStatus(ctx, msg)
}

func (k *msgServer) MsgLinkNode(c context.Context, msg *v2.MsgLinkNodeRequest) (*v2.MsgLinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgLinkNode(ctx, msg)
}

func (k *msgServer) MsgUnlinkNode(c context.Context, msg *v2.MsgUnlinkNodeRequest) (*v2.MsgUnlinkNodeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUnlinkNode(ctx, msg)
}

func (k *msgServer) MsgSubscribe(_ context.Context, _ *v2.MsgSubscribeRequest) (*v2.MsgSubscribeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
