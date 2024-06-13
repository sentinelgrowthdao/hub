package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
)

var (
	_ v1.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(k Keeper) v1.MsgServiceServer {
	return &msgServer{k}
}

func (k *msgServer) MsgUpdateParams(c context.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if msg.From != k.authority {
		return nil, types.NewErrorInvalidSigner(msg.From, k.authority)
	}

	k.SetParams(ctx, msg.Params)
	return &v1.MsgUpdateParamsResponse{}, nil
}
