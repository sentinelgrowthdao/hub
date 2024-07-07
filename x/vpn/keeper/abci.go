package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func cacheContext(c sdk.Context) (cc sdk.Context, writeCache func()) {
	cms := c.MultiStore().CacheMultiStore()
	cc = c.WithMultiStore(cms)
	return cc, cms.Write
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	ctx, write := cacheContext(ctx)
	defer write()

	k.Node.BeginBlock(ctx)
	k.Session.BeginBlock(ctx)
	k.Subscription.BeginBlock(ctx)
}
