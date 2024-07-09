package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/provider/types"
	"github.com/sentinel-official/hub/v12/x/provider/types/v2"
)

func (k *Keeper) HandleMsgRegister(ctx sdk.Context, msg *v2.MsgRegisterRequest) (*v2.MsgRegisterResponse, error) {
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provAddr := base.ProvAddress(accAddr.Bytes())
	if k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorDuplicateProvider(provAddr)
	}

	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	provider := v2.Provider{
		Address:     provAddr.String(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
		Status:      v1base.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventRegister{
			Address: provider.Address,
		},
	)

	return &v2.MsgRegisterResponse{}, nil
}

func (k *Keeper) HandleMsgUpdate(ctx sdk.Context, msg *v2.MsgUpdateRequest) (*v2.MsgUpdateResponse, error) {
	provAddr, err := base.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	if msg.Status.Equal(v1base.StatusInactive) {
		if err := k.ProviderInactivePreHook(ctx, provAddr); err != nil {
			return nil, err
		}
	}

	if msg.Status.Equal(v1base.StatusActive) {
		if provider.Status.Equal(v1base.StatusInactive) {
			k.DeleteInactiveProvider(ctx, provAddr)
		}
	}
	if msg.Status.Equal(v1base.StatusInactive) {
		if provider.Status.Equal(v1base.StatusActive) {
			k.DeleteActiveProvider(ctx, provAddr)
		}
	}

	if msg.Name != "" {
		provider.Name = msg.Name
	}
	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.Description = msg.Description

	if !msg.Status.Equal(v1base.StatusUnspecified) {
		provider.Status = msg.Status
		provider.StatusAt = ctx.BlockTime()
	}

	k.SetProvider(ctx, provider)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdate{
			Address: provider.Address,
		},
	)

	return &v2.MsgUpdateResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v2.MsgUpdateParamsRequest) (*v2.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	k.SetParams(ctx, msg.Params)
	return &v2.MsgUpdateParamsResponse{}, nil
}
