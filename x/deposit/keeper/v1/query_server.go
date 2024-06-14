package v1

import (
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/v12/x/deposit/keeper"
	"github.com/sentinel-official/hub/v12/x/deposit/types"
	"github.com/sentinel-official/hub/v12/x/deposit/types/v1"
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

func (k *queryServer) QueryDeposit(c context.Context, req *v1.QueryDepositRequest) (*v1.QueryDepositResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address %s", req.Address)
	}

	ctx := sdk.UnwrapSDKContext(c)

	item, found := k.GetDeposit(ctx, addr)
	if !found {
		return nil, status.Errorf(codes.NotFound, "deposit does not exist for address %s", req.Address)
	}

	return &v1.QueryDepositResponse{Deposit: item}, nil
}

func (k *queryServer) QueryDeposits(c context.Context, req *v1.QueryDepositsRequest) (*v1.QueryDepositsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	var (
		items v1.Deposits
		ctx   = sdk.UnwrapSDKContext(c)
		store = prefix.NewStore(k.Store(ctx), types.DepositKeyPrefix)
	)

	pagination, err := sdkquery.Paginate(store, req.Pagination, func(_ []byte, value []byte) error {
		var item v1.Deposit
		if err := k.Unmarshal(value, &item); err != nil {
			return err
		}

		items = append(items, item)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v1.QueryDepositsResponse{Deposits: items, Pagination: pagination}, nil
}
