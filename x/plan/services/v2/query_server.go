package v2

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/plan/keeper"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	keeper.Keeper
}

func NewQueryServiceServer(k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{k}
}

func (q *queryServer) QueryPlan(_ context.Context, _ *v2.QueryPlanRequest) (*v2.QueryPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPlans(_ context.Context, _ *v2.QueryPlansRequest) (*v2.QueryPlansResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (q *queryServer) QueryPlansForProvider(_ context.Context, _ *v2.QueryPlansForProviderRequest) (*v2.QueryPlansForProviderResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
