package v2

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

var (
	_ v2.QueryServiceServer = (*queryServer)(nil)
)

type queryServer struct {
	codec.BinaryCodec
	keeper.Keeper
}

func NewQueryServiceServer(cdc codec.BinaryCodec, k keeper.Keeper) v2.QueryServiceServer {
	return &queryServer{
		BinaryCodec: cdc,
		Keeper:      k,
	}
}

func (k *queryServer) QueryAllocation(c context.Context, req *v2.QueryAllocationRequest) (*v2.QueryAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetAllocation(ctx, req.Id, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "allocation %d/%s does not exist", req.Id, req.Address)
	}

	return &v2.QueryAllocationResponse{Allocation: item}, nil
}

func (k *queryServer) QueryAllocations(c context.Context, req *v2.QueryAllocationsRequest) (*v2.QueryAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v2.Allocations
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetAllocationForSubscriptionKeyPrefix(req.Id))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Allocation
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryAllocationsResponse{Allocations: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = k.GetParams(ctx)
	)

	return &v2.QueryParamsResponse{Params: params}, nil
}

func (k *queryServer) QuerySubscriptions(_ context.Context, _ *v2.QuerySubscriptionsRequest) (*v2.QuerySubscriptionsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscriptionsForAccount(_ context.Context, _ *v2.QuerySubscriptionsForAccountRequest) (*v2.QuerySubscriptionsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscriptionsForNode(_ context.Context, _ *v2.QuerySubscriptionsForNodeRequest) (*v2.QuerySubscriptionsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscriptionsForPlan(_ context.Context, _ *v2.QuerySubscriptionsForPlanRequest) (*v2.QuerySubscriptionsForPlanResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QuerySubscription(_ context.Context, _ *v2.QuerySubscriptionRequest) (*v2.QuerySubscriptionResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayouts(_ context.Context, _ *v2.QueryPayoutsRequest) (*v2.QueryPayoutsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayoutsForAccount(_ context.Context, _ *v2.QueryPayoutsForAccountRequest) (*v2.QueryPayoutsForAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayoutsForNode(_ context.Context, _ *v2.QueryPayoutsForNodeRequest) (*v2.QueryPayoutsForNodeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (k *queryServer) QueryPayout(_ context.Context, _ *v2.QueryPayoutRequest) (*v2.QueryPayoutResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
