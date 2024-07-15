package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	k.IterateLeasesForNode(ctx, addr, func(_ int, item v1.Lease) bool {
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		handler := k.router.Handler(msg)
		if _, err := handler(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}

func (k *Keeper) ProviderInactivePreHook(ctx sdk.Context, addr base.ProvAddress) error {
	k.IterateLeasesForProvider(ctx, addr, func(_ int, item v1.Lease) bool {
		msg := &v1.MsgEndLeaseRequest{
			From: item.ProvAddress,
			ID:   item.ID,
		}

		handler := k.router.Handler(msg)
		if _, err := handler(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})

	return nil
}
