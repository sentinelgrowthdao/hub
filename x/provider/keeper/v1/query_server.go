package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/types/v1"
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

func (k *queryServer) QueryProviders(c context.Context, req *v1.QueryProvidersRequest) (*v1.QueryProvidersResponse, error) {
	return &v1.QueryProvidersResponse{}, nil
}

func (k *queryServer) QueryProvider(c context.Context, req *v1.QueryProviderRequest) (*v1.QueryProviderResponse, error) {
	return &v1.QueryProviderResponse{}, nil
}

func (k *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return &v1.QueryParamsResponse{}, nil
}
