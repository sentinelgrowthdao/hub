package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/types/v1"
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

func (k *queryServer) QueryProviders(_ context.Context, _ *v1.QueryProvidersRequest) (*v1.QueryProvidersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryProvider(_ context.Context, _ *v1.QueryProviderRequest) (*v1.QueryProviderResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryParams(_ context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
