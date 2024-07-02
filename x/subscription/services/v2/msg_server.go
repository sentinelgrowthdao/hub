package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
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

func (k *msgServer) MsgCancel(c context.Context, msg *v2.MsgCancelRequest) (*v2.MsgCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgCancel(ctx, msg)
}

func (k *msgServer) MsgAllocate(c context.Context, msg *v2.MsgAllocateRequest) (*v2.MsgAllocateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgAllocate(ctx, msg)
}
