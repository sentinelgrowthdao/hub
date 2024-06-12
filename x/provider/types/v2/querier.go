package v2

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/hub/v12/types"
)

func NewQueryProviderRequest(addr base.ProvAddress) *QueryProviderRequest {
	return &QueryProviderRequest{
		Address: addr.String(),
	}
}

func NewQueryProvidersRequest(status base.Status, pagination *query.PageRequest) *QueryProvidersRequest {
	return &QueryProvidersRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
