package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/session/types"
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

func (k *Keeper) HandleQuerySession(ctx sdk.Context, req *v2.QuerySessionRequest) (*v2.QuerySessionResponse, error) {
	item, found := k.GetSession(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "session does not exist for id %d", req.Id)
	}

	return &v2.QuerySessionResponse{Session: item}, nil
}

func (k *Keeper) HandleQuerySessions(ctx sdk.Context, req *v2.QuerySessionsRequest) (*v2.QuerySessionsResponse, error) {
	var (
		items v2.Sessions
		store = prefix.NewStore(k.Store(ctx), types.SessionKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Session
		if err := k.cdc.Unmarshal(value, &item); err != nil {
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

func (k *Keeper) HandleQuerySessionsForAccount(ctx sdk.Context, req *v2.QuerySessionsForAccountRequest) (*v2.QuerySessionsForAccountResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Sessions
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

func (k *Keeper) HandleQuerySessionsForNode(ctx sdk.Context, req *v2.QuerySessionsForNodeRequest) (*v2.QuerySessionsForNodeResponse, error) {
	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Sessions
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

func (k *Keeper) HandleQuerySessionsForSubscription(ctx sdk.Context, req *v2.QuerySessionsForSubscriptionRequest) (*v2.QuerySessionsForSubscriptionResponse, error) {
	var (
		items v2.Sessions
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

func (k *Keeper) HandleQuerySessionsForAllocation(ctx sdk.Context, req *v2.QuerySessionsForAllocationRequest) (*v2.QuerySessionsForAllocationResponse, error) {
	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items v2.Sessions
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

func (k *Keeper) HandleQueryParams(ctx sdk.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	params := k.GetParams(ctx)
	return &v2.QueryParamsResponse{Params: params}, nil
}
