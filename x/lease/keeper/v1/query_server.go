package v1

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
	"github.com/sentinel-official/hub/v12/x/lease/keeper"
	"github.com/sentinel-official/hub/v12/x/lease/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
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

func (k *queryServer) QueryLease(c context.Context, req *v1.QueryLeaseRequest) (*v1.QueryLeaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetLease(ctx, req.Id)
	if !found {
		return nil, status.Errorf(codes.NotFound, "lease %d does not exist", req.Id)
	}

	return &v1.QueryLeaseResponse{Lease: item}, nil
}

func (k *queryServer) QueryLeases(c context.Context, req *v1.QueryLeasesRequest) (*v1.QueryLeasesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items []v1.Lease
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.LeaseKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v1.Lease
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryLeasesResponse{Leases: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryLeasesForProvider(c context.Context, req *v1.QueryLeasesForProviderRequest) (*v1.QueryLeasesForProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []v1.Lease
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetLeaseForProviderKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetLease(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("lease for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryLeasesForProviderResponse{Leases: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryLeasesForNode(c context.Context, req *v1.QueryLeasesForNodeRequest) (*v1.QueryLeasesForNodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.NodeAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	var (
		items []v1.Lease
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.GetLeaseForNodeKeyPrefix(addr))
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(key, _ []byte) error {
		item, found := k.GetLease(ctx, sdk.BigEndianToUint64(key))
		if !found {
			return fmt.Errorf("lease for key %X does not exist", key)
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryLeasesForNodeResponse{Leases: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = k.GetParams(ctx)
	)

	return &v1.QueryParamsResponse{Params: params}, nil
}
