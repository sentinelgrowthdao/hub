package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
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

func (k *msgServer) MsgStart(_ context.Context, _ *v2.MsgStartRequest) (*v2.MsgStartResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgUpdateDetails(_ context.Context, _ *v2.MsgUpdateDetailsRequest) (*v2.MsgUpdateDetailsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *msgServer) MsgEnd(c context.Context, msg *v2.MsgEndRequest) (*v2.MsgEndResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgEnd(ctx, msg)
}
