package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/plan/types"
	"github.com/sentinel-official/hub/v12/x/plan/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	Keeper
}

func NewQueryServiceServer(keeper Keeper) v2.QueryServiceServer {
	return &queryServer{Keeper: keeper}
}

func (q *queryServer) QueryPlan(c context.Context, req *v2.QueryPlanRequest) (*v2.QueryPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := q.GetPlan(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "plan does not exist for id %d", req.Id)
	}

	return &v2.QueryPlanResponse{Plan: item}, nil
}

func (q *queryServer) QueryPlans(c context.Context, req *v2.QueryPlansRequest) (res *v2.QueryPlansResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items     v2.Plans
		keyPrefix []byte
		ctx       = sdk.UnwrapSDKContext(c)
	)

	switch req.Status {
	case base.StatusActive:
		keyPrefix = types.ActivePlanKeyPrefix
	case base.StatusInactive:
		keyPrefix = types.InactivePlanKeyPrefix
	default:
		keyPrefix = types.PlanKeyPrefix
	}

	store := prefix.NewStore(q.Store(ctx), keyPrefix)
	pagination, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Plan
		if err := q.cdc.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryPlansResponse{Plans: items, Pagination: pagination}, nil
}

func (q *queryServer) QueryPlansForProvider(c context.Context, req *v2.QueryPlansForProviderRequest) (res *v2.QueryPlansForProviderResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Plans
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(q.Store(ctx), types.GetPlanForProviderKeyPrefix(addr))
	)

	pagination, err := query.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		if !accumulate {
			return false, nil
		}

		item, found := q.GetPlan(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return false, fmt.Errorf("plan for key %X does not exist", key)
		}

		if req.Status.Equal(base.StatusUnspecified) || item.Status.Equal(req.Status) {
			items = append(items, item)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryPlansForProviderResponse{Plans: items, Pagination: pagination}, nil
}
