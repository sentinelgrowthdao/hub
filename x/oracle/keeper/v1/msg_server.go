package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/oracle/keeper"
	"github.com/sentinel-official/hub/v12/x/oracle/types"
	"github.com/sentinel-official/hub/v12/x/oracle/types/v1"
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

func (k *msgServer) MsgUpdateParams(c context.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	authority := k.GetAuthority()
	if msg.From != authority {
		return nil, types.NewErrorInvalidSigner(msg.From, authority)
	}

	k.SetParams(ctx, msg.Params)
	return &v1.MsgUpdateParamsResponse{}, nil
}
