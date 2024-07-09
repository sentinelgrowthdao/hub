package keeper

import (
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
	v1base "github.com/sentinel-official/hub/v12/types/v1"
	"github.com/sentinel-official/hub/v12/x/node/types"
	"github.com/sentinel-official/hub/v12/x/node/types/v2"
	"github.com/sentinel-official/hub/v12/x/node/types/v3"
)

func (k *Keeper) HandleMsgRegister(ctx sdk.Context, msg *v2.MsgRegisterRequest) (*v2.MsgRegisterResponse, error) {
	if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
		return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
	}
	if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
		return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr := base.NodeAddress(accAddr.Bytes())
	if k.HasNode(ctx, nodeAddr) {
		return nil, types.NewErrorDuplicateNode(nodeAddr)
	}

	deposit := k.Deposit(ctx)
	if err := k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	node := v2.Node{
		Address:        nodeAddr.String(),
		GigabytePrices: msg.GigabytePrices,
		HourlyPrices:   msg.HourlyPrices,
		RemoteURL:      msg.RemoteURL,
		InactiveAt:     time.Time{},
		Status:         v1base.StatusInactive,
		StatusAt:       ctx.BlockTime(),
	}

	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventRegister{
			Address: node.Address,
		},
	)

	return &v2.MsgRegisterResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateDetails(ctx sdk.Context, msg *v2.MsgUpdateDetailsRequest) (*v2.MsgUpdateDetailsResponse, error) {
	if msg.GigabytePrices != nil {
		if !k.IsValidGigabytePrices(ctx, msg.GigabytePrices) {
			return nil, types.NewErrorInvalidPrices(msg.GigabytePrices)
		}
	}
	if msg.HourlyPrices != nil {
		if !k.IsValidHourlyPrices(ctx, msg.HourlyPrices) {
			return nil, types.NewErrorInvalidPrices(msg.HourlyPrices)
		}
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	if msg.GigabytePrices != nil {
		node.GigabytePrices = msg.GigabytePrices
	}
	if msg.HourlyPrices != nil {
		node.HourlyPrices = msg.HourlyPrices
	}
	if msg.RemoteURL != "" {
		node.RemoteURL = msg.RemoteURL
	}

	k.SetNode(ctx, node)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateDetails{
			Address: node.Address,
		},
	)

	return &v2.MsgUpdateDetailsResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateStatus(ctx sdk.Context, msg *v2.MsgUpdateStatusRequest) (*v2.MsgUpdateStatusResponse, error) {
	nodeAddr, err := base.NodeAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}

	k.DeleteNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
	if node.Status.Equal(v1base.StatusActive) {
		k.DeleteActiveNode(ctx, nodeAddr)
	}
	if node.Status.Equal(v1base.StatusInactive) {
		k.DeleteInactiveNode(ctx, nodeAddr)
	}

	node.Status = msg.Status
	node.InactiveAt = time.Time{}
	node.StatusAt = ctx.BlockTime()

	if node.Status.Equal(v1base.StatusActive) {
		duration := k.ActiveDuration(ctx)
		node.InactiveAt = ctx.BlockTime().Add(duration)
	}

	k.SetNode(ctx, node)
	k.SetNodeForInactiveAt(ctx, node.InactiveAt, nodeAddr)
	ctx.EventManager().EmitTypedEvent(
		&v2.EventUpdateStatus{
			Status:  node.Status,
			Address: node.Address,
		},
	)

	return &v2.MsgUpdateStatusResponse{}, nil
}

func (k *Keeper) HandleMsgStartSession(ctx sdk.Context, msg *v3.MsgStartSessionRequest) (*v3.MsgStartSessionResponse, error) {
	if msg.Gigabytes != 0 {
		if ok := k.IsValidSessionGigabytes(ctx, msg.Gigabytes); !ok {
			return nil, types.NewErrorInvalidGigabytes(msg.Gigabytes)
		}
	}
	if msg.Hours != 0 {
		if ok := k.IsValidSessionHours(ctx, msg.Hours); !ok {
			return nil, types.NewErrorInvalidHours(msg.Hours)
		}
	}

	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	nodeAddr, err := base.NodeAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return nil, err
	}

	node, found := k.GetNode(ctx, nodeAddr)
	if !found {
		return nil, types.NewErrorNodeNotFound(nodeAddr)
	}
	if !node.Status.Equal(v1base.StatusActive) {
		return nil, types.NewErrorInvalidNodeStatus(nodeAddr, node.Status)
	}

	var (
		count   = k.session.GetCount(ctx)
		delay   = k.session.StatusChangeDelay(ctx)
		session = &v3.Session{
			ID:            count + 1,
			AccAddress:    accAddr.String(),
			NodeAddress:   nodeAddr.String(),
			Price:         sdk.Coin{},
			Deposit:       sdk.Coin{},
			DownloadBytes: sdkmath.ZeroInt(),
			UploadBytes:   sdkmath.ZeroInt(),
			MaxBytes:      sdkmath.ZeroInt(),
			Duration:      0,
			MaxDuration:   0,
			Status:        v1base.StatusActive,
			InactiveAt:    ctx.BlockTime().Add(delay),
			StatusAt:      ctx.BlockTime(),
		}
	)

	if msg.Gigabytes != 0 {
		price, found := node.GigabytePrice(msg.Denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(msg.Denom)
		}

		session.Price = price
		session.Deposit = sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Gigabytes),
		)
		session.MaxBytes = base.Gigabyte.MulRaw(msg.Gigabytes)
	}
	if msg.Hours != 0 {
		price, found := node.HourlyPrice(msg.Denom)
		if !found {
			return nil, types.NewErrorPriceNotFound(msg.Denom)
		}

		session.Price = price
		session.Deposit = sdk.NewCoin(
			price.Denom,
			price.Amount.MulRaw(msg.Hours),
		)
		session.MaxDuration = time.Duration(msg.Hours) * time.Hour
	}

	if err := k.AddDeposit(ctx, accAddr, session.Deposit); err != nil {
		return nil, err
	}

	k.session.SetCount(ctx, count+1)
	k.session.SetSession(ctx, session)
	k.session.SetSessionForAccount(ctx, accAddr, session.ID)
	k.session.SetSessionForNode(ctx, nodeAddr, session.ID)
	k.session.SetSessionForInactiveAt(ctx, session.InactiveAt, session.ID)

	return &v3.MsgStartSessionResponse{}, nil
}

func (k *Keeper) HandleMsgUpdateParams(ctx sdk.Context, msg *v3.MsgUpdateParamsRequest) (*v3.MsgUpdateParamsResponse, error) {
	if msg.From != k.authority {
		return nil, types.NewErrorUnauthorized(msg.From)
	}

	minGigabytePrices := k.MinGigabytePrices(ctx)
	minHourlyPrices := k.MinHourlyPrices(ctx)
	k.SetParams(ctx, msg.Params)

	minGigabytePricesModified := !msg.Params.MinGigabytePrices.IsEqual(minGigabytePrices)
	minHourlyPricesModified := !msg.Params.MinHourlyPrices.IsEqual(minHourlyPrices)

	if minGigabytePricesModified {
		minGigabytePrices = k.MinGigabytePrices(ctx)
	}
	if minHourlyPricesModified {
		minHourlyPrices = k.MinHourlyPrices(ctx)
	}

	if minGigabytePricesModified || minHourlyPricesModified {
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
			return false
		})
	}

	return &v3.MsgUpdateParamsResponse{}, nil
}
