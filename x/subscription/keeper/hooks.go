package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v3"
)

func (k *Keeper) SessionInactivePreHook(ctx sdk.Context, id uint64) error {
	item, found := k.session.GetSession(ctx, id)
	if !found {
		return fmt.Errorf("session %d does not exist", id)
	}
	if !item.GetStatus().Equal(v1base.StatusInactivePending) {
		return fmt.Errorf("invalid status %s for session %d", item.GetStatus(), item.GetStatus())
	}

	session, ok := item.(*subscriptiontypes.Session)
	if !ok {
		return nil
	}

	subscription, found := k.GetSubscription(ctx, session.SubscriptionID)
	if !found {
		return fmt.Errorf("subscription %d does not exist", session.SubscriptionID)
	}

	var (
		accAddr       = session.GetAccAddress()
		utilisedBytes = session.DownloadBytes.Add(session.UploadBytes)
	)

	alloc, found := k.GetAllocation(ctx, subscription.ID, accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", subscription.ID, accAddr)
	}

	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(utilisedBytes)
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	k.SetAllocation(ctx, alloc)
	return nil
}
