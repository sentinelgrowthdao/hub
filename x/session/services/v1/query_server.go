package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v1"
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

func (k *queryServer) QuerySessions(_ context.Context, _ *v1.QuerySessionsRequest) (*v1.QuerySessionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySessionsForAddress(_ context.Context, _ *v1.QuerySessionsForAddressRequest) (*v1.QuerySessionsForAddressResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySession(_ context.Context, _ *v1.QuerySessionRequest) (*v1.QuerySessionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryParams(_ context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
