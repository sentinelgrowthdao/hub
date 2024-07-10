package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v3"
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

func (k *msgServer) MsgEnd(c context.Context, msg *v3.MsgEndRequest) (*v3.MsgEndResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgEnd(ctx, msg)
}

func (k *msgServer) MsgUpdate(c context.Context, msg *v3.MsgUpdateRequest) (*v3.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdate(ctx, msg)
}

func (k *msgServer) MsgUpdateParams(c context.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateParams(ctx, msg)
}
