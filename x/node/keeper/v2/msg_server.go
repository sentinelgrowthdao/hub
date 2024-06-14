package v2

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/keeper"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
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

// MsgRegister registers a new node in the network.
// It validates the registration request, checks prices, and creates a new node.
func (k *msgServer) MsgRegister(c context.Context, msg *v2.MsgRegisterRequest) (*v2.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the provided GigabytePrices are valid, if not, return an error.
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}

	// Check if the provided HourlyPrices are valid, if not, return an error.
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert the account address to a `base.NodeAddress`.
	nodeAddr := base.NodeAddress(accAddr.Bytes())

	// Check if the node already exists in the network. If found, return an error to prevent duplicate registration.
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	// Get the required deposit for registering a new node.
	deposit := k.Deposit(ctx)

	// Fund the community pool with the required deposit from the registrant's account.
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Create a new node with the provided information and set its status to `Inactive`.
	node := v2.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteURL:      msg.RemoteURL,
		InactiveAt:     time.Time{},
		Status:         v1base.StatusInactive,
		StatusAt:       ctx.BlockTime(),
	}

	// Save the new node in the Store.
	k.SetNode(ctx, node)

	// Emit an event to notify that a new node has been registered.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventRegister{
			Address: node.Address,
		},
	)

	return &v2.MsgRegisterResponse{}, nil
}

// MsgUpdateDetails updates the details of a registered node.
// It validates the update details request, checks prices, and updates the node information.
func (k *msgServer) MsgUpdateDetails(c context.Context, msg *v2.MsgUpdateDetailsRequest) (*v2.MsgUpdateDetailsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the provided GigabytePrices are valid, if not, return an error.
	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
		}
	}

	// Check if the provided HourlyPrices are valid, if not, return an error.
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
		}
	}

	// Convert the `msg.From` address (in Bech32 format) to a `base.NodeAddress`.
	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the node from the Store based on the provided node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// Update the node's GigabytePrices, HourlyPrices, and RemoteURL with the provided values.
	if msg.GigabytePrices != nil {
		node.GigabytePrices = msg.GigabytePrices
	}
	if msg.HourlyPrices != nil {
		node.HourlyPrices = msg.HourlyPrices
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	// Save the updated node in the Store.
	k.SetNode(ctx, node)

	// Emit an event to notify that the node details have been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateDetails{
			Address: node.Address,
		},
	)

	return &v2.MsgUpdateDetailsResponse{}, nil
}

// MsgUpdateStatus updates the status of a registered node.
// It validates the update status request, checks the node's current status, and updates the status and inactive time accordingly.
func (k *msgServer) MsgUpdateStatus(c context.Context, msg *v2.MsgUpdateStatusRequest) (*v2.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to a `base.NodeAddress`.
	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the node from the Store based on the provided node address.
	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	// If the current status of the node is `Active`, handle the necessary updates for changing to `Inactive` status.
	if node.Status.Equal(v1base.StatusActive) {
		k.DeleteNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
		if msg.Status.Equal(v1base.StatusInactive) {
			k.DeleteActiveNode(ctx, nodeAddr)
		}
	}

	// If the current status of the node is `Inactive`, handle the necessary updates for changing to `Active` status.
	if node.Status.Equal(v1base.StatusInactive) {
		if msg.Status.Equal(v1base.StatusActive) {
			k.DeleteInactiveNode(ctx, nodeAddr)
		}
	}

	// If the new status is `Active`, update the node's inactive time based on the active duration.
	if msg.Status.Equal(v1base.StatusActive) {
		node.InactiveAt = ctx.BlockTime().Add(
			k.ActiveDuration(ctx),
		)
		k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
	}

	// If the new status is `Inactive`, set the node's inactive time to zero.
	if msg.Status.Equal(v1base.StatusInactive) {
		node.InactiveAt = time.Time{}
	}

	// Update the node's status and status timestamp.
	node.Status = msg.Status
	node.StatusAt = ctx.BlockTime()

	// Save the updated node in the Store.
	k.SetNode(ctx, node)

	// Emit an event to notify that the node status has been updated.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateStatus{
			Status:  node.Status,
			Address: node.Address,
		},
	)

	return &v2.MsgUpdateStatusResponse{}, nil
}

// MsgSubscribe subscribes to a node for a specific amount of gigabytes or hours.
// It validates the subscription request and creates a new subscription for the provided node and user account.
func (k *msgServer) MsgSubscribe(c context.Context, msg *v2.MsgSubscribeRequest) (*v2.MsgSubscribeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Check if the provided Gigabytes value is valid, if not, return an error.
	if msg.Gigabytes != 0 {
		if !k.IsValidSubscriptionGigabytes(ctx, msg.Gigabytes) {
			return nil, types.NewErrorInvalidGigabytes(msg.Gigabytes)
		}
	}

	// Check if the provided Hours value is valid, if not, return an error.
	if msg.Hours != 0 {
		if !k.IsValidSubscriptionHours(ctx, msg.Hours) {
			return nil, types.NewErrorInvalidHours(msg.Hours)
		}
	}

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert the `msg.NodeAddress` (node address) to a `base.NodeAddress`.
	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	// Create a new subscription for the provided node, user account, gigabytes, hours, and denom.
	subscription, err := k.CreateSubscriptionForNode(ctx, accAddr, nodeAddr, msg.Gigabytes, msg.Hours, msg.Denom)
	if err != nil {
		return nil, err
	}

	// Emit an event to notify that a new subscription has been created.
	ctx.EventManager().EmitTypedEvent(
		&v2.EventCreateSubscription{
			Address:     subscription.Address,
			NodeAddress: subscription.NodeAddress,
			ID:          subscription.ID,
		},
	)

	return &v2.MsgSubscribeResponse{}, nil
}
