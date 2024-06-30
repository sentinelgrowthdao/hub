package v3

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/subscription/keeper"
	"github.com/sentinel-official/hub/v12/x/subscription/types"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v3"
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

func (k *queryServer) QuerySubscription(c context.Context, req *v3.QuerySubscriptionRequest) (*v3.QuerySubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetSubscription(ctx, req.ID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "subscription %d does not exist", req.ID)
	}

	return &v3.QuerySubscriptionResponse{Subscription: item}, nil
}

func (k *queryServer) QuerySubscriptions(c context.Context, req *v3.QuerySubscriptionsRequest) (*v3.QuerySubscriptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []v3.Subscription
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.SubscriptionKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v3.Subscription
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySubscriptionsResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySubscriptionsForAccount(c context.Context, req *v3.QuerySubscriptionsForAccountRequest) (*v3.QuerySubscriptionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []v3.Subscription
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForAccountKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySubscriptionsForAccountResponse{Subscriptions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySubscriptionsForPlan(c context.Context, req *v3.QuerySubscriptionsForPlanRequest) (*v3.QuerySubscriptionsForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []v3.Subscription
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSubscriptionForPlanKeyPrefix(req.ID))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSubscription(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("subscription for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v3.QuerySubscriptionsForPlanResponse{Subscriptions: items, Pagination: pagination}, nil
}
