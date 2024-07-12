package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
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

func (q *queryServer) QueryNode(c context.Context, req *v2.QueryNodeRequest) (*v2.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNode(ctx, req)
}

func (q *queryServer) QueryNodes(c context.Context, req *v2.QueryNodesRequest) (res *v2.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNodes(ctx, req)
}

func (q *queryServer) QueryNodesForPlan(c context.Context, req *v2.QueryNodesForPlanRequest) (*v2.QueryNodesForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	return q.HandleQueryNodesForPlan(ctx, req)
}

func (q *queryServer) QueryParams(_ context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
