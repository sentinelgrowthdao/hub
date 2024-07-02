package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
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

func (k *queryServer) QueryAllocation(c context.Context, req *v2.QueryAllocationRequest) (*v2.QueryAllocationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryAllocation(ctx, req)
}

func (k *queryServer) QueryAllocations(c context.Context, req *v2.QueryAllocationsRequest) (*v2.QueryAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryAllocations(ctx, req)
}

func (k *queryServer) QueryParams(c context.Context, req *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQueryParams(ctx, req)
}

func (k *queryServer) QuerySubscriptions(_ context.Context, _ *v2.QuerySubscriptionsRequest) (*v2.QuerySubscriptionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscriptionsForAccount(_ context.Context, _ *v2.QuerySubscriptionsForAccountRequest) (*v2.QuerySubscriptionsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscriptionsForNode(_ context.Context, _ *v2.QuerySubscriptionsForNodeRequest) (*v2.QuerySubscriptionsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscriptionsForPlan(_ context.Context, _ *v2.QuerySubscriptionsForPlanRequest) (*v2.QuerySubscriptionsForPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscription(_ context.Context, _ *v2.QuerySubscriptionRequest) (*v2.QuerySubscriptionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayouts(_ context.Context, _ *v2.QueryPayoutsRequest) (*v2.QueryPayoutsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayoutsForAccount(_ context.Context, _ *v2.QueryPayoutsForAccountRequest) (*v2.QueryPayoutsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayoutsForNode(_ context.Context, _ *v2.QueryPayoutsForNodeRequest) (*v2.QueryPayoutsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayout(_ context.Context, _ *v2.QueryPayoutRequest) (*v2.QueryPayoutResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
