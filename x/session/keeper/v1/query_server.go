package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types/v1"
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

func (k *queryServer) QuerySessions(c context.Context, req *v1.QuerySessionsRequest) (*v1.QuerySessionsResponse, error) {
	return &v1.QuerySessionsResponse{}, nil
}

func (k *queryServer) QuerySessionsForAddress(c context.Context, req *v1.QuerySessionsForAddressRequest) (*v1.QuerySessionsForAddressResponse, error) {
	return &v1.QuerySessionsForAddressResponse{}, nil
}

func (k *queryServer) QuerySession(c context.Context, req *v1.QuerySessionRequest) (*v1.QuerySessionResponse, error) {
	return &v1.QuerySessionResponse{}, nil
}

func (k *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return &v1.QueryParamsResponse{}, nil
}
