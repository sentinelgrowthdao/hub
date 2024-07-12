package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v1"
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

func (q *queryServer) QuerySubscriptions(_ context.Context, _ *v1.QuerySubscriptionsRequest) (*v1.QuerySubscriptionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySubscriptionsForAddress(_ context.Context, _ *v1.QuerySubscriptionsForAddressRequest) (*v1.QuerySubscriptionsForAddressResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QuerySubscription(_ context.Context, _ *v1.QuerySubscriptionRequest) (*v1.QuerySubscriptionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryQuota(_ context.Context, _ *v1.QueryQuotaRequest) (*v1.QueryQuotaResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryQuotas(_ context.Context, _ *v1.QueryQuotasRequest) (*v1.QueryQuotasResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryParams(_ context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
