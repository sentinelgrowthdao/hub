package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
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

func (k *msgServer) MsgRegister(c context.Context, msg *v2.MsgRegisterRequest) (*v2.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgRegister(ctx, msg)
}

func (k *msgServer) MsgUpdate(c context.Context, msg *v2.MsgUpdateRequest) (*v2.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdate(ctx, msg)
}

func (k *msgServer) MsgUpdateParams(c context.Context, msg *v2.MsgUpdateParamsRequest) (*v2.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateParams(ctx, msg)
}
