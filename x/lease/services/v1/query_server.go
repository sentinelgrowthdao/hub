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
	_ v1.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v1.QueryServiceServer {
	return &queryServer{k}
}

func (k *queryServer) QueryLease(c context.Context, req *v1.QueryLeaseRequest) (*v1.QueryLeaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryLease(ctx, req)
}

func (k *queryServer) QueryLeases(c context.Context, req *v1.QueryLeasesRequest) (*v1.QueryLeasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryLeases(ctx, req)
}

func (k *queryServer) QueryLeasesForProvider(c context.Context, req *v1.QueryLeasesForProviderRequest) (*v1.QueryLeasesForProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryLeasesForProvider(ctx, req)
}

func (k *queryServer) QueryLeasesForNode(c context.Context, req *v1.QueryLeasesForNodeRequest) (*v1.QueryLeasesForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryLeasesForNode(ctx, req)
}

func (k *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryParams(ctx, req)
}
