package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) SendCoinFromDepositToAccount(ctx sdk.Context, from, to sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToModule(ctx sdk.Context, from sdk.AccAddress, to string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToModule(ctx, from, to, sdk.NewCoins(coin))
}

func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	return nil
}
