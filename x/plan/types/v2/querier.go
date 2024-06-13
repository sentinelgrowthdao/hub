package v2

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
)

func NewQueryPlanRequest(id uint64) *QueryPlanRequest {
	return &QueryPlanRequest{
		Id: id,
	}
}

func NewQueryPlansRequest(status v1base.Status, pagination *query.PageRequest) *QueryPlansRequest {
	return &QueryPlansRequest{
		Status:     status,
		Pagination: pagination,
	}
}

func NewQueryPlansForProviderRequest(addr base.ProvAddress, status v1base.Status, pagination *query.PageRequest) *QueryPlansForProviderRequest {
	return &QueryPlansForProviderRequest{
		Address:    addr.String(),
		Status:     status,
		Pagination: pagination,
	}
}
