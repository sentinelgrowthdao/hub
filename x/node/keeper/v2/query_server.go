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
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
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

func (k *queryServer) QueryNode(c context.Context, req *v2.QueryNodeRequest) (*v2.QueryNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetNode(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "node does not exist for address %s", req.Address)
	}

	return &v2.QueryNodeResponse{Node: item}, nil
}

func (k *queryServer) QueryNodes(c context.Context, req *v2.QueryNodesRequest) (res *v2.QueryNodesResponse, err error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items     []v2.Node
		keyPrefix []byte
		ctx       = sdk.UnwrapSDKContext(c)
	)

	switch req.Status {
	case v1base.StatusActive:
		keyPrefix = types.ActiveNodeKeyPrefix
	case v1base.StatusInactive:
		keyPrefix = types.InactiveNodeKeyPrefix
	default:
		keyPrefix = types.NodeKeyPrefix
	}

	store := prefix.NewStore(k.Store(ctx), keyPrefix)
	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Node
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryNodesResponse{Nodes: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryNodesForPlan(c context.Context, req *v2.QueryNodesForPlanRequest) (*v2.QueryNodesForPlanResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []v2.Node
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetNodeForPlanKeyPrefix(req.Id))
	)

	pagination, err := sdkquery.FilteredPaginate(store, req.Pagination, func(key, _ []byte, accumulate bool) (bool, error) {
		if !accumulate {
			return false, nil
		}

		item, found := k.GetNode(ctx, key[1:])
		if !found {
			return false, fmt.Errorf("node for key %X does not exist", key)
		}

		if req.Status.Equal(v1base.StatusUnspecified) || item.Status.Equal(req.Status) {
			items = append(items, item)
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryNodesForPlanResponse{Nodes: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
