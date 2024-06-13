package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	base "github.com/sentinel-official/hub/v12/types"
)

func NewQuerySubscriptionRequest(id uint64) *QuerySubscriptionRequest {
	return &QuerySubscriptionRequest{
		Id: id,
	}
}

func NewQuerySubscriptionsRequest(pagination *sdkquery.PageRequest) *QuerySubscriptionsRequest {
	return &QuerySubscriptionsRequest{
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForAccountRequest(addr sdk.AccAddress, pagination *sdkquery.PageRequest) *QuerySubscriptionsForAccountRequest {
	return &QuerySubscriptionsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForNodeRequest(addr base.NodeAddress, pagination *sdkquery.PageRequest) *QuerySubscriptionsForNodeRequest {
	return &QuerySubscriptionsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQuerySubscriptionsForPlanRequest(id uint64, pagination *sdkquery.PageRequest) *QuerySubscriptionsForPlanRequest {
	return &QuerySubscriptionsForPlanRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQueryAllocationRequest(id uint64, addr sdk.AccAddress) *QueryAllocationRequest {
	return &QueryAllocationRequest{
		Id:      id,
		Address: addr.String(),
	}
}

func NewQueryAllocationsRequest(id uint64, pagination *sdkquery.PageRequest) *QueryAllocationsRequest {
	return &QueryAllocationsRequest{
		Id:         id,
		Pagination: pagination,
	}
}

func NewQueryPayoutRequest(id uint64) *QueryPayoutRequest {
	return &QueryPayoutRequest{
		Id: id,
	}
}

func NewQueryPayoutsRequest(pagination *sdkquery.PageRequest) *QueryPayoutsRequest {
	return &QueryPayoutsRequest{
		Pagination: pagination,
	}
}

func NewQueryPayoutsForAccountRequest(addr sdk.AccAddress, pagination *sdkquery.PageRequest) *QueryPayoutsForAccountRequest {
	return &QueryPayoutsForAccountRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryPayoutsForNodeRequest(addr base.NodeAddress, pagination *sdkquery.PageRequest) *QueryPayoutsForNodeRequest {
	return &QueryPayoutsForNodeRequest{
		Address:    addr.String(),
		Pagination: pagination,
	}
}

func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
