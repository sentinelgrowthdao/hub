package v2

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

// The following line asserts that the `msgServer` type implements the `types.MsgServiceServer` interface.
var (
	_ v2.MsgServiceServer = (*msgServer)(nil)
)

// msgServer is a message server that implements the `types.MsgServiceServer` interface.
type msgServer struct {
	keeper.Keeper // Keeper is an instance of the main keeper for the module.
}

// NewMsgServiceServer creates a new instance of `types.MsgServiceServer` using the provided Keeper.
func NewMsgServiceServer(k keeper.Keeper) v2.MsgServiceServer {
	return &msgServer{k}
}

// MsgRegister registers a new provider with the provided details and stores it in the Store.
// It validates the registration request, checks for provider existence, and assigns a unique address to the provider.
func (k *msgServer) MsgRegister(c context.Context, msg *v2.MsgRegisterRequest) (*v2.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert the `accAddr` to a `base.ProvAddress` to represent the provider address.
	provAddr := base.ProvAddress(accAddr.Bytes())

	// Check if the provider with the given address exists in the network. If yes, return an error.
	if k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorDuplicateProvider(provAddr)
	}

	// Get the deposit value from the Store and fund the community pool with the deposit from the provider.
	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Create a new provider with the provided details and set its status as `Inactive`.
	provider := v2.Provider{
		Address:     provAddr.String(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	// Save the new provider in the Store.
	k.SetProvider(ctx, provider)

	// Emit an event to notify that a new provider has been registered.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventRegister{
			Address: provider.Address,
		},
	)

	return &v2.MsgRegisterResponse{}, nil
}

// MsgUpdate updates the details of a provider.
// It validates the update request, checks for provider existence, and updates the provider's details and status.
func (k *msgServer) MsgUpdate(c context.Context, msg *v2.MsgUpdateRequest) (*v2.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to a `base.ProvAddress`.
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the provider from the Store based on the provided provider address.
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	// Update the provider's details (name, identity, website, description) if they are provided in the message.
	if len(msg.Name) > 0 {
		provider.Name = msg.Name
	}
	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.Description = msg.Description

	// If the status is provided in the message and it is not `StatusUnspecified`, update the provider's status.
	if !msg.Status.Equal(v1base.StatusUnspecified) {
		// If the current status of the provider is `Active`, handle the necessary updates for changing to `Inactive` status.
		if provider.Status.Equal(v1base.StatusActive) {
			if msg.Status.Equal(v1base.StatusInactive) {
				k.DeleteActiveProvider(ctx, provAddr)
			}
		}
		// If the current status of the provider is `Inactive`, handle the necessary updates for changing to `Active` status.
		if provider.Status.Equal(v1base.StatusInactive) {
			if msg.Status.Equal(v1base.StatusActive) {
				k.DeleteInactiveProvider(ctx, provAddr)
			}
		}

		// Update the provider's status and status timestamp with the provided new status and current block time.
		provider.Status = msg.Status
		provider.StatusAt = ctx.BlockTime()
	}

	// Save the updated provider in the Store.
	k.SetProvider(ctx, provider)

	// Emit an event to notify that the provider details have been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdate{
			Address: provider.Address,
		},
	)

	return &v2.MsgUpdateResponse{}, nil
}
