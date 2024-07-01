package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

// SessionInactiveHook is a function that handles the end of a session.
// It updates the allocation's utilized bytes, calculates and sends payments, and staking rewards.
func (k *Keeper) SessionInactiveHook(ctx sdk.Context, id uint64, utilisedBytes sdkmath.Int) error {
	// Retrieve the session associated with the provided session ID.
	session, found := k.GetSession(ctx, id)
	if !found {
		return fmt.Errorf("session %d does not exist", id)
	}

	// Check if the session has the correct status for processing.
	if !session.Status.Equal(v1base.StatusInactivePending) {
		return fmt.Errorf("invalid status %s for session %d", session.Status, session.ID)
	}

	// Retrieve the subscription associated with the session.
	subscription, found := k.GetSubscription(ctx, session.SubscriptionID)
	if !found {
		return fmt.Errorf("subscription %d does not exist", session.SubscriptionID)
	}

	accAddr, err := sdk.AccAddressFromBech32(session.Address)
	if err != nil {
		return err
	}

	// Retrieve the allocation associated with the subscription and account address.
	alloc, found := k.GetAllocation(ctx, subscription.ID, accAddr)
	if !found {
		return fmt.Errorf("subscription allocation %d/%s does not exist", subscription.ID, accAddr)
	}

	// Update the allocation's utilized bytes by adding the provided bytes.
	alloc.UtilisedBytes = alloc.UtilisedBytes.Add(utilisedBytes)

	// Ensure that the utilized bytes don't exceed the granted bytes.
	if alloc.UtilisedBytes.GT(alloc.GrantedBytes) {
		alloc.UtilisedBytes = alloc.GrantedBytes
	}

	// Save the updated allocation to the store.
	k.SetAllocation(ctx, alloc)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventAllocate{
			Address:       alloc.Address,
			GrantedBytes:  alloc.GrantedBytes,
			UtilisedBytes: alloc.UtilisedBytes,
			ID:            alloc.ID,
		},
	)

	return nil
}
