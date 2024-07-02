package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/swap/keeper"
	"github.com/sentinel-official/hub/v12/x/swap/types/v1"
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

func (k *msgServer) MsgSwap(c context.Context, msg *v1.MsgSwapRequest) (*v1.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgSwap(ctx, msg)
}
