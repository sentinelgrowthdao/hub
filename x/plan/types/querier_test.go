package types

import (
	"crypto/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestNewQueryPlanRequest(t *testing.T) {
	for i := 0; i < 20; i++ {
		require.Equal(
			t,
			&QueryPlanRequest{
				Id: uint64(i),
			},
			NewQueryPlanRequest(uint64(i)),
		)
	}
}

func TestNewQueryPlansForProviderRequest(t *testing.T) {
	var (
		address    []byte
		status     base.Status
		pagination *query.PageRequest
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		status = base.Status(i % 4)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(address)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryPlansForProviderRequest{
				Address:    base.ProvAddress(address).String(),
				Status:     status,
				Pagination: pagination,
			},
			NewQueryPlansForProviderRequest(address, status, pagination),
		)
	}
}

func TestNewQueryPlansRequest(t *testing.T) {
	var (
		status     base.Status
		pagination *query.PageRequest
	)

	for i := 0; i < 20; i++ {
		status = base.Status(i % 4)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryPlansRequest{
				Status:     status,
				Pagination: pagination,
			},
			NewQueryPlansRequest(status, pagination),
		)
	}
}
