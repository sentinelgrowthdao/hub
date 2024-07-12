package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/plan/keeper"
	"github.com/sentinel-official/hub/v12/x/plan/types/v1"
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

func (q *queryServer) QueryPlans(_ context.Context, _ *v1.QueryPlansRequest) (*v1.QueryPlansResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPlansForProvider(_ context.Context, _ *v1.QueryPlansForProviderRequest) (*v1.QueryPlansForProviderResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPlan(_ context.Context, _ *v1.QueryPlanRequest) (*v1.QueryPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryNodesForPlan(_ context.Context, _ *v1.QueryNodesForPlanRequest) (*v1.QueryNodesForPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
