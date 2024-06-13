package v1

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/sentinel-official/hub/v12/x/swap/types"
)

func NewQuerySwapRequest(txHash types.EthereumHash) *QuerySwapRequest {
	return &QuerySwapRequest{
		TxHash: txHash.Bytes(),
	}
}

func NewQuerySwapsRequest(pagination *query.PageRequest) *QuerySwapsRequest {
	return &QuerySwapsRequest{
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
