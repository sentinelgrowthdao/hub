package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/swap/keeper"
	"github.com/sentinel-official/hub/v12/x/swap/types"
	"github.com/sentinel-official/hub/v12/x/swap/types/v1"
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

func (k *queryServer) QuerySwap(c context.Context, req *v1.QuerySwapRequest) (*v1.QuerySwapResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		ctx  = sdk.UnwrapSDKContext(c)
		hash = types.BytesToHash(req.TxHash)
	)

	item, found := k.GetSwap(ctx, hash)
	if !found {
		return nil, status.Errorf(codes.NotFound, "swap does not exist for hash %X", req.TxHash)
	}

	return &v1.QuerySwapResponse{Swap: item}, nil
}

func (k *queryServer) QuerySwaps(c context.Context, req *v1.QuerySwapsRequest) (*v1.QuerySwapsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v1.Swaps
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.SwapKeyPrefix)
	)

	pagination, err := sdkquery.FilteredPaginate(store, req.Pagination, func(_, value []byte, accumulate bool) (bool, error) {
		if accumulate {
			var item v1.Swap
			if err := k.Unmarshal(value, &item); err != nil {
				return false, err
			}

			items = append(items, item)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QuerySwapsResponse{Swaps: items, Pagination: pagination}, nil
}

func (k *queryServer) QueryParams(c context.Context, _ *v1.QueryParamsRequest) (*v1.QueryParamsResponse, error) {
	var (
		ctx    = sdk.UnwrapSDKContext(c)
		params = k.GetParams(ctx)
	)

	return &v1.QueryParamsResponse{Params: params}, nil
}
