package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func (k *Keeper) FundCommunityPool(ctx sdk.Context, fromAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.distribution.FundCommunityPool(ctx, sdk.NewCoins(coin), fromAddr)
}

func (k *Keeper) AddDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.AddDeposit(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) NodeInactivePreHook(ctx sdk.Context, addr base.NodeAddress) error {
	return nil
}
