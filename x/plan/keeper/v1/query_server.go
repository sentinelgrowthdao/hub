package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/sentinel-official/hub/v12/x/plan/keeper"
	"github.com/sentinel-official/hub/v12/x/plan/types/v1"
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

func (k *queryServer) QueryPlans(c context.Context, req *v1.QueryPlansRequest) (*v1.QueryPlansResponse, error) {
	return &v1.QueryPlansResponse{}, nil
}

func (k *queryServer) QueryPlansForProvider(c context.Context, req *v1.QueryPlansForProviderRequest) (*v1.QueryPlansForProviderResponse, error) {
	return &v1.QueryPlansForProviderResponse{}, nil
}

func (k *queryServer) QueryPlan(c context.Context, req *v1.QueryPlanRequest) (*v1.QueryPlanResponse, error) {
	return &v1.QueryPlanResponse{}, nil
}

func (k *queryServer) QueryNodesForPlan(c context.Context, req *v1.QueryNodesForPlanRequest) (*v1.QueryNodesForPlanResponse, error) {
	return &v1.QueryNodesForPlanResponse{}, nil
}
