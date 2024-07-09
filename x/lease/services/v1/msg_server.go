package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/lease/keeper"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
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

func (k *msgServer) MsgStart(c context.Context, msg *v1.MsgStartRequest) (*v1.MsgStartResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgStart(ctx, msg)
}

func (k *msgServer) MsgUpdateDetails(c context.Context, msg *v1.MsgUpdateDetailsRequest) (*v1.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateDetails(ctx, msg)
}

func (k *msgServer) MsgEnd(c context.Context, msg *v1.MsgEndRequest) (*v1.MsgEndResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgEnd(ctx, msg)
}

func (k *msgServer) MsgUpdateParams(c context.Context, msg *v1.MsgUpdateParamsRequest) (*v1.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleMsgUpdateParams(ctx, msg)
}

func (k *msgServer) MsgRenew(_ context.Context, _ *v1.MsgRenewRequest) (*v1.MsgRenewResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
