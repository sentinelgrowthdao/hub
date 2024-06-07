package keeper

import (
	abcitypes "github.com/cometbft/cometbft/abci/types"
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

	k.Subscription.BeginBlock(ctx)
}

func (k *Keeper) EndBlock(ctx sdk.Context) abcitypes.ValidatorUpdates {
	ctx, write := cacheContext(ctx)
	defer write()

	k.Node.EndBlock(ctx)
	k.Session.EndBlock(ctx)
	k.Subscription.EndBlock(ctx)

	return nil
}
