package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types/v1"
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

func (k *queryServer) QueryNodes(_ context.Context, _ *v1.QueryNodesRequest) (*v1.QueryNodesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryNodesForProvider(_ context.Context, _ *v1.QueryNodesForProviderRequest) (*v1.QueryNodesForProviderResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryNode(_ context.Context, _ *v1.QueryNodeRequest) (*v1.QueryNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryParams(_ context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
