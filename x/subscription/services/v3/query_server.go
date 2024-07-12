package v3

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

var (
	_ v3.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v3.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QuerySubscription(c context.Context, req *v3.QuerySubscriptionRequest) (*v3.QuerySubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySubscription(ctx, req)
}

func (q *queryServer) QuerySubscriptions(c context.Context, req *v3.QuerySubscriptionsRequest) (*v3.QuerySubscriptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySubscriptions(ctx, req)
}

func (q *queryServer) QuerySubscriptionsForAccount(c context.Context, req *v3.QuerySubscriptionsForAccountRequest) (*v3.QuerySubscriptionsForAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySubscriptionsForAccount(ctx, req)
}

func (q *queryServer) QuerySubscriptionsForPlan(c context.Context, req *v3.QuerySubscriptionsForPlanRequest) (*v3.QuerySubscriptionsForPlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQuerySubscriptionsForPlan(ctx, req)
}
