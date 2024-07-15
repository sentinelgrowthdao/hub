package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func cacheContext(ctx sdk.Context) (c sdk.Context, write func()) {
	cms := ctx.MultiStore().CacheMultiStore()
	c = ctx.WithMultiStore(cms)

	return c, cms.Write
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	ctx, write := cacheContext(ctx)
	defer write()

	k.Lease.BeginBlock(ctx)
	k.Node.BeginBlock(ctx)
	k.Session.BeginBlock(ctx)
	k.Subscription.BeginBlock(ctx)
}
