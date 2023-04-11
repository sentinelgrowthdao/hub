package types

import (
	"fmt"
)

type (
	GenesisPlans []GenesisPlan
	GenesisState GenesisPlans
)

func NewGenesisState(plans GenesisPlans) GenesisState {
	return GenesisState(plans)
}

func DefaultGenesisState() GenesisState {
	return NewGenesisState(nil)
}

func ValidateGenesis(state GenesisState) error {
	m := make(map[uint64]bool)
	for _, item := range state {
		if m[item.Plan.ID] {
			return fmt.Errorf("found duplicate plan for id %d", item.Plan.ID)
		}

		m[item.Plan.ID] = true
	}

	for _, item := range state {
		m := make(map[string]bool)
		for _, addr := range item.Nodes {
			if m[addr] {
				return fmt.Errorf("found duplicate node %s for plan %d", addr, item.Plan.ID)
			}

			m[addr] = true
		}
	}

	for _, item := range state {
		if err := item.Plan.Validate(); err != nil {
			return err
		}
	}

	return nil
}
