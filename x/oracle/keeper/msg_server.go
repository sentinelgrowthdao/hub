package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgUpdateParams(c context.Context, msg *types.MsgUpdateParamsRequest) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if msg.From != k.authority {
		return nil, types.NewErrorInvalidSigner(msg.From, k.authority)
	}

	k.SetParams(ctx, msg.Params)
	return &types.MsgUpdateParamsResponse{}, nil
}
