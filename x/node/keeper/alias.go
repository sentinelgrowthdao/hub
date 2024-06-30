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

	return k.deposit.Add(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) SubtractDeposit(ctx sdk.Context, addr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.Subtract(ctx, addr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToAccount(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToAccount(ctx, fromAddr, toAddr, sdk.NewCoins(coin))
}

func (k *Keeper) SendCoinFromDepositToModule(ctx sdk.Context, fromAddr sdk.AccAddress, toModule string, coin sdk.Coin) error {
	if coin.IsZero() {
		return nil
	}

	return k.deposit.SendCoinsFromDepositToModule(ctx, fromAddr, toModule, sdk.NewCoins(coin))
}

func (k *Keeper) HasProvider(ctx sdk.Context, addr base.ProvAddress) bool {
	return k.provider.HasProvider(ctx, addr)
}

func (k *Keeper) LeaseInactivePreHook(ctx sdk.Context, id uint64) error {
	return nil
}
