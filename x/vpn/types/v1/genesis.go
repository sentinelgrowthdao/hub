package v1

import (
	sdkerrors "cosmossdk.io/errors"

	deposittypes "github.com/sentinel-official/hub/v12/x/deposit/types/v1"
	nodetypes "github.com/sentinel-official/hub/v12/x/node/types/v3"
	plantypes "github.com/sentinel-official/hub/v12/x/plan/types/v2"
	providertypes "github.com/sentinel-official/hub/v12/x/provider/types/v2"
	sessiontypes "github.com/sentinel-official/hub/v12/x/session/types/v2"
	subscriptiontypes "github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func NewGenesisState(
	deposits *deposittypes.GenesisState,
	providers *providertypes.GenesisState,
	nodes *nodetypes.GenesisState,
	plans *plantypes.GenesisState,
	subscriptions *subscriptiontypes.GenesisState,
	sessions *sessiontypes.GenesisState,
) *GenesisState {
	return &GenesisState{
		Deposits:      deposits,
		Providers:     providers,
		Nodes:         nodes,
		Plans:         plans,
		Subscriptions: subscriptions,
		Sessions:      sessions,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		deposittypes.DefaultGenesisState(),
		providertypes.DefaultGenesisState(),
		nodetypes.DefaultGenesisState(),
		plantypes.DefaultGenesisState(),
		subscriptiontypes.DefaultGenesisState(),
		sessiontypes.DefaultGenesisState(),
	)
}

func (m *GenesisState) Validate() error {
	if err := deposittypes.ValidateGenesisState(m.Deposits); err != nil {
		return sdkerrors.Wrapf(err, "invalid deposit genesis")
	}
	if err := providertypes.ValidateGenesis(m.Providers); err != nil {
		return sdkerrors.Wrapf(err, "invalid provider genesis")
	}
	if err := nodetypes.ValidateGenesis(m.Nodes); err != nil {
		return sdkerrors.Wrapf(err, "invalid node genesis")
	}
	if err := plantypes.ValidateGenesis(m.Plans); err != nil {
		return sdkerrors.Wrapf(err, "invalid plan genesis")
	}
	if err := subscriptiontypes.ValidateGenesis(m.Subscriptions); err != nil {
		return sdkerrors.Wrapf(err, "invalid subscription genesis")
	}
	if err := sessiontypes.ValidateGenesis(m.Sessions); err != nil {
		return sdkerrors.Wrapf(err, "invalid session genesis")
	}

	return nil
}
