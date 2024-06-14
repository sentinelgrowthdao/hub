package v2

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/keeper"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
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

func (k *queryServer) QuerySession(c context.Context, req *v2.QuerySessionRequest) (*v2.QuerySessionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetSession(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "session does not exist for id %d", req.Id)
	}

	return &v2.QuerySessionResponse{Session: item}, nil
}

func (k *queryServer) QuerySessions(c context.Context, req *v2.QuerySessionsRequest) (*v2.QuerySessionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v2.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.SessionKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Session
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySessionsResponse{Sessions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySessionsForAccount(c context.Context, req *v2.QuerySessionsForAccountRequest) (*v2.QuerySessionsForAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForAccountKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySessionsForAccountResponse{Sessions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySessionsForNode(c context.Context, req *v2.QuerySessionsForNodeRequest) (*v2.QuerySessionsForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForNodeKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySessionsForNodeResponse{Sessions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySessionsForSubscription(c context.Context, req *v2.QuerySessionsForSubscriptionRequest) (*v2.QuerySessionsForSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v2.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForSubscriptionKeyPrefix(req.Id))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySessionsForSubscriptionResponse{Sessions: items, Pagination: pagination}, nil
}

func (k *queryServer) QuerySessionsForAllocation(c context.Context, req *v2.QuerySessionsForAllocationRequest) (*v2.QuerySessionsForAllocationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Sessions
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetSessionForAllocationKeyPrefix(req.Id, addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetSession(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("session for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QuerySessionsForAllocationResponse{Sessions: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = k.GetParams(ctx)
	)

	return &v2.QueryParamsResponse{Params: params}, nil
}
