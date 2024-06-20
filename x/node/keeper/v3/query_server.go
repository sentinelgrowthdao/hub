package v3

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

var (
	_ v3.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	codec.BinaryCodec
	keeper.Keeper
}

func NewQueryServiceServer(cdc codec.BinaryCodec, k keeper.Keeper) v3.QueryServiceServer {
	return &queryServer{
		BinaryCodec: cdc,
		Keeper:      k,
	}
}

func (k *queryServer) QueryLease(c context.Context, req *v3.QueryLeaseRequest) (*v3.QueryLeaseResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryLeases(c context.Context, req *v3.QueryLeasesRequest) (*v3.QueryLeasesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryLeasesForProvider(c context.Context, req *v3.QueryLeasesForProviderRequest) (*v3.QueryLeasesForProviderResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryLeasesForNode(c context.Context, req *v3.QueryLeasesForNodeRequest) (*v3.QueryLeasesForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
