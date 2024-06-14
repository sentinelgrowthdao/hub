package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v1"
)

var (
	_ v1.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	codec.BinaryCodec
	keeper.Keeper
}

func NewQueryServiceServer(cdc codec.BinaryCodec, k keeper.Keeper) v1.QueryServiceServer {
	return &queryServer{
		BinaryCodec: cdc,
		Keeper:      k,
	}
}

func (k *queryServer) QuerySubscriptions(c context.Context, req *v1.QuerySubscriptionsRequest) (*v1.QuerySubscriptionsResponse, error) {
	return &v1.QuerySubscriptionsResponse{}, nil
}

func (k *queryServer) QuerySubscriptionsForAddress(c context.Context, req *v1.QuerySubscriptionsForAddressRequest) (*v1.QuerySubscriptionsForAddressResponse, error) {
	return &v1.QuerySubscriptionsForAddressResponse{}, nil
}

func (k *queryServer) QuerySubscription(c context.Context, req *v1.QuerySubscriptionRequest) (*v1.QuerySubscriptionResponse, error) {
	return &v1.QuerySubscriptionResponse{}, nil
}

func (k *queryServer) QueryQuota(c context.Context, req *v1.QueryQuotaRequest) (*v1.QueryQuotaResponse, error) {
	return &v1.QueryQuotaResponse{}, nil
}

func (k *queryServer) QueryQuotas(c context.Context, req *v1.QueryQuotasRequest) (*v1.QueryQuotasResponse, error) {
	return &v1.QueryQuotasResponse{}, nil
}

func (k *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return &v1.QueryParamsResponse{}, nil
}
