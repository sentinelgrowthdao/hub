package v1

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/swap/keeper"
	"github.com/sentinel-official/hub/v12/x/swap/types/v1"
)

var (
	_ v1.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v1.QueryServiceServer {
	return &queryServer{k}
}

func (k *queryServer) QuerySwap(c context.Context, req *v1.QuerySwapRequest) (*v1.QuerySwapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySwap(ctx, req)
}

func (k *queryServer) QuerySwaps(c context.Context, req *v1.QuerySwapsRequest) (*v1.QuerySwapsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySwaps(ctx, req)
}

func (k *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryParams(ctx, req)
}