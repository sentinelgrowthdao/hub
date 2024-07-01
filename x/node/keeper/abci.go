package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func (k *Keeper) handlePriceUpdates(ctx sdk.Context) {
	var (
		minGigabytePricesModified = k.IsMinGigabytePricesModified(ctx)
		minHourlyPricesModified   = k.IsMinHourlyPricesModified(ctx)
	)

	if !minGigabytePricesModified && !minHourlyPricesModified {
		return
	}

	minGigabytePrices := sdk.NewCoins()
	if minGigabytePricesModified {
		minGigabytePrices = k.MinGigabytePrices(ctx)
	}

	minHourlyPrices := sdk.NewCoins()
	if minHourlyPricesModified {
		minHourlyPrices = k.MinHourlyPrices(ctx)
	}

	k.IterateNodes(ctx, func(_ int, item v2.Node) bool {
		if minGigabytePricesModified {
			for _, coin := range minGigabytePrices {
				amount := item.GigabytePrices.AmountOf(coin.Denom)
				if amount.LT(coin.Amount) {
					item.GigabytePrices = item.GigabytePrices.Sub(
						sdk.NewCoin(coin.Denom, amount),
					).Add(coin)
				}
			}
		}

		if minHourlyPricesModified {
			for _, coin := range minHourlyPrices {
				amount := item.HourlyPrices.AmountOf(coin.Denom)
				if amount.LT(coin.Amount) {
					item.HourlyPrices = item.HourlyPrices.Sub(
						sdk.NewCoin(coin.Denom, amount),
					).Add(coin)
				}
			}
		}

		k.SetNode(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&v2.EventUpdateDetails{
				Address:        item.Address,
				GigabytePrices: item.GigabytePrices.String(),
				HourlyPrices:   item.HourlyPrices.String(),
			},
		)

		return false
	})
}

func (k *Keeper) handleInactiveNodes(ctx sdk.Context) {
	k.IterateNodesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v2.Node) bool {
		nodeAddr := item.GetAddress()

		k.DeleteActiveNode(ctx, nodeAddr)
		k.DeleteNodeForInactiveAt(ctx, item.InactiveAt, nodeAddr)

		item.InactiveAt = time.Time{}
		item.Status = v1base.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetNode(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&v2.EventUpdateStatus{
				Status:  v1base.StatusInactive,
				Address: item.Address,
			},
		)

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactiveNodes(ctx)
	k.handlePriceUpdates(ctx)
}
