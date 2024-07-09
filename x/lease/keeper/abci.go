package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	baseutils "github.com/sentinel-official/hub/v12/utils"
	"github.com/sentinel-official/hub/v12/x/lease/types/v1"
)

func (k *Keeper) handleInactiveLeases(ctx sdk.Context) {
	k.IterateLeasesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
		msg := item.MsgEndRequest()
		if _, err := k.HandleMsgEnd(ctx, msg); err != nil {
			panic(err)
		}

		return false
	})
}

func (k *Keeper) handleLeasePayouts(ctx sdk.Context) {
	share := k.StakingShare(ctx)

	k.IterateLeasesForPayoutAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) (stop bool) {
		k.DeleteLeaseForPayoutAt(ctx, item.PayoutAt, item.ID)

		provAddr, err := base.ProvAddressFromBech32(item.ProvAddress)
		if err != nil {
			panic(err)
		}

		reward := baseutils.GetProportionOfCoin(item.Price, share)
		if err := k.SendCoinFromDepositToModule(ctx, provAddr.Bytes(), k.feeCollectorName, reward); err != nil {
			panic(err)
		}

		nodeAddr, err := base.NodeAddressFromBech32(item.NodeAddress)
		if err != nil {
			panic(err)
		}

		payment := item.Price.Sub(reward)
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
		k.SetLeaseForPayoutAt(ctx, item.PayoutAt, item.ID)

		return false
	})
}

func (k *Keeper) handleLeaseRenewals(ctx sdk.Context) {
	k.IterateLeasesForRenewalAt(ctx, ctx.BlockTime(), func(_ int, item v1.Lease) bool {
		msg := &v1.MsgRenewRequest{
			From:  item.ProvAddress,
			ID:    item.ID,
			Hours: item.MaxHours,
			Denom: item.Price.Denom,
		}

		if _, err := k.HandleMsgRenew(ctx, msg); err != nil {
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
