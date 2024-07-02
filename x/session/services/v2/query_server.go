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
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{k}
}

func (k *queryServer) QuerySession(c context.Context, req *v2.QuerySessionRequest) (*v2.QuerySessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySession(ctx, req)
}

func (k *queryServer) QuerySessions(c context.Context, req *v2.QuerySessionsRequest) (*v2.QuerySessionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySessions(ctx, req)
}

func (k *queryServer) QuerySessionsForAccount(c context.Context, req *v2.QuerySessionsForAccountRequest) (*v2.QuerySessionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySessionsForAccount(ctx, req)
}

func (k *queryServer) QuerySessionsForNode(c context.Context, req *v2.QuerySessionsForNodeRequest) (*v2.QuerySessionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySessionsForNode(ctx, req)
}

func (k *queryServer) QuerySessionsForSubscription(c context.Context, req *v2.QuerySessionsForSubscriptionRequest) (*v2.QuerySessionsForSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySessionsForSubscription(ctx, req)
}

func (k *queryServer) QuerySessionsForAllocation(c context.Context, req *v2.QuerySessionsForAllocationRequest) (*v2.QuerySessionsForAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySessionsForAllocation(ctx, req)
}

func (k *queryServer) QueryParams(c context.Context, req *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryParams(ctx, req)
}
