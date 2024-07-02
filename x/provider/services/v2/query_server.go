package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{k}
}

func (k *queryServer) QueryProvider(c context.Context, req *v2.QueryProviderRequest) (*v2.QueryProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryProvider(ctx, req)
}

func (k *queryServer) QueryProviders(c context.Context, req *v2.QueryProvidersRequest) (*v2.QueryProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryProviders(ctx, req)
}

func (k *queryServer) QueryParams(c context.Context, req *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryParams(ctx, req)
}
