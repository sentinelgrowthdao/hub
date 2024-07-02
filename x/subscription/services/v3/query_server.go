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

func (k *queryServer) QuerySubscription(c context.Context, req *v3.QuerySubscriptionRequest) (*v3.QuerySubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySubscription(ctx, req)
}

func (k *queryServer) QuerySubscriptions(c context.Context, req *v3.QuerySubscriptionsRequest) (*v3.QuerySubscriptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySubscriptions(ctx, req)
}

func (k *queryServer) QuerySubscriptionsForAccount(c context.Context, req *v3.QuerySubscriptionsForAccountRequest) (*v3.QuerySubscriptionsForAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySubscriptionsForAccount(ctx, req)
}

func (k *queryServer) QuerySubscriptionsForPlan(c context.Context, req *v3.QuerySubscriptionsForPlanRequest) (*v3.QuerySubscriptionsForPlanResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return k.HandleQuerySubscriptionsForPlan(ctx, req)
}
