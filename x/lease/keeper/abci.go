package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) handleInactiveLeases(ctx sdk.Context) {
	k.IterateLeasesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
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
		k.DeleteLeaseForProvider(ctx, provAddr, item.ID)
		k.DeleteLeaseForProviderByNode(ctx, provAddr, nodeAddr, item.ID)
		k.DeleteLeaseForPayoutAt(ctx, item.PayoutAt, item.ID)
		k.DeleteLeaseForRenewalAt(ctx, item.RenewalAt, item.ID)

		return false
	})
}

func (k *Keeper) handleLeasePayouts(ctx sdk.Context) {
	k.IterateLeasesForPayoutAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) (stop bool) {
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
	k.IterateLeasesForRenewalAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
		k.DeleteLeaseForRenewalAt(ctx, item.RenewalAt, item.ID)

		msg := &v1.MsgRenewRequest{
			From:  item.ProvAddress,
			ID:    item.ID,
			Hours: item.MaxHours,
			Denom: item.Price.Denom,
		}

		if _, err := k.HandleMsgRenewLease(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})
}

func (k *Keeper) BeginBlock(ctx sdk.Context) {
	k.handleInactiveLeases(ctx)
	k.handleLeasePayouts(ctx)
	k.handleLeaseRenewals(ctx)
}
