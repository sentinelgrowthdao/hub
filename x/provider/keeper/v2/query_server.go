package v2

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
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

func (k *queryServer) QueryProvider(c context.Context, req *v2.QueryProviderRequest) (*v2.QueryProviderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := base.ProvAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetProvider(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "provider %s does not exist", req.Address)
	}

	return &v2.QueryProviderResponse{Provider: item}, nil
}

func (k *queryServer) QueryProviders(c context.Context, req *v2.QueryProvidersRequest) (*v2.QueryProvidersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items     v2.Providers
		keyPrefix []byte
		ctx       = sdk.UnwrapSDKContext(c)
	)

	switch req.Status {
	case v1base.StatusActive:
		keyPrefix = types.ActiveProviderKeyPrefix
	case v1base.StatusInactive:
		keyPrefix = types.InactiveProviderKeyPrefix
	default:
		keyPrefix = types.ProviderKeyPrefix
	}

	store := prefix.NewStore(k.Store(ctx), keyPrefix)
	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_, value []byte) error {
		var item v2.Provider
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v2.QueryProvidersResponse{Providers: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v2.QueryParamsRequest) (*v2.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = k.GetParams(ctx)
	)

	return &v2.QueryParamsResponse{Params: params}, nil
}
