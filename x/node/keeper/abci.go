package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
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

func (k *Keeper) handleInactiveLeases(ctx sdk.Context) {
	k.IterateLeasesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v3.Lease) bool {
		k.DeleteLeaseForInactiveAt(ctx, item.InactiveAt, item.ID)

		if err := k.LeaseInactivePreHook(ctx, item.ID); err != nil {
			panic(err)
		}

		var (
			nodeAddr = item.GetNodeAddress()
			provAddr = item.GetProvAddress()
		)

		k.DeleteLease(ctx, item.ID)
		k.DeleteLeaseForNode(ctx, nodeAddr, item.ID)
		k.DeleteLeaseForPayoutAt(ctx, item.PayoutAt, item.ID)
		k.DeleteLeaseForProvider(ctx, provAddr, item.ID)
		k.DeleteLeaseForProviderByNode(ctx, provAddr, nodeAddr, item.ID)
		k.DeleteLeaseForRenewAt(ctx, item.RenewAt, item.ID)

		return false
	})
}

func (k *Keeper) handleLeasePayouts(ctx sdk.Context) {
	k.IterateLeasesForPayoutAt(ctx, ctx.BlockTime(), func(_ int, item v3.Lease) (stop bool) {
		k.DeleteLeaseForPayoutAt(ctx, item.PayoutAt, item.ID)

		var (
			nodeAddr     = item.GetNodeAddress()
			provAddr     = item.GetProvAddress()
			stakingShare = k.StakingShare(ctx)
		)

		stakingReward := baseutils.GetProportionOfCoin(item.Price, stakingShare)
		if err := k.SendCoinFromDepositToModule(ctx, provAddr.Bytes(), k.feeCollectorName, stakingReward); err != nil {
			panic(err)
		}

		payment := item.Price.Sub(stakingReward)
		if err := k.SendCoinFromDepositToAccount(ctx, provAddr.Bytes(), nodeAddr.Bytes(), payment); err != nil {
			panic(err)
		}

		item.Hours = item.Hours + 1
		if item.Hours < item.MaxHours {
			item.PayoutAt = item.PayoutAt.Add(time.Hour)
		} else {
			item.PayoutAt = time.Time{}
		}

		k.SetLease(ctx, item)

		if item.Hours < item.MaxHours {
			k.SetLeaseForPayoutAt(ctx, item.PayoutAt, item.ID)
		}

		return false
	})
}

func (k *Keeper) handleLeaseRenewals(ctx sdk.Context) {
	k.IterateLeasesForRenewAt(ctx, ctx.BlockTime(), func(_ int, item v3.Lease) bool {
		k.DeleteLeaseForRenewAt(ctx, item.RenewAt, item.ID)

		msg := &v3.MsgRenewLeaseRequest{
			From:      "",
			ID:        item.ID,
			Hours:     item.MaxHours,
			Denom:     item.Price.Denom,
			Renewable: true,
		}

		if _, err := k.RenewLease(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactiveLeases(ctx)
	k.handleInactiveNodes(ctx)
	k.handleLeasePayouts(ctx)
	k.handleLeaseRenewals(ctx)
	k.handlePriceUpdates(ctx)
}
