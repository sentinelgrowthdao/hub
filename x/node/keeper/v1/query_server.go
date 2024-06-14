package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"

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

func (k *queryServer) QueryNodes(c context.Context, req *v1.QueryNodesRequest) (*v1.QueryNodesResponse, error) {
	return &v1.QueryNodesResponse{}, nil
}

func (k *queryServer) QueryNodesForProvider(c context.Context, req *v1.QueryNodesForProviderRequest) (*v1.QueryNodesForProviderResponse, error) {
	return &v1.QueryNodesForProviderResponse{}, nil
}

func (k *queryServer) QueryNode(c context.Context, req *v1.QueryNodeRequest) (*v1.QueryNodeResponse, error) {
	return &v1.QueryNodeResponse{}, nil
}

func (k *queryServer) QueryParams(c context.Context, req *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	return &v1.QueryParamsResponse{}, nil
}
