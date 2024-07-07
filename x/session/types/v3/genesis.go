package v3

import (
	"github.com/sentinel-official/hub/v12/x/session/types/v2"
)

func NewGenesisState(_ []Session, params v2.Params) *GenesisState {
	return &GenesisState{
		Sessions: nil,
		Params:   params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, v2.DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	return nil
}
