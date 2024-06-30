package v3

import (
	"fmt"

	"github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func NewGenesisState(nodes []v2.Node, leases []Lease, params Params) *GenesisState {
	return &GenesisState{
		Nodes:  nodes,
		Leases: leases,
		Params: params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	nodeMap := make(map[string]bool)
	for _, node := range state.Nodes {
		if nodeMap[node.Address] {
			return fmt.Errorf("found a duplicate node for address %s", node.Address)
		}

		nodeMap[node.Address] = true

		if err := node.Validate(); err != nil {
			return err
		}
	}

	leaseMap := make(map[uint64]bool)
	for _, lease := range state.Leases {
		if leaseMap[lease.ID] {
			return fmt.Errorf("found a duplicate lease with ID %d", lease.ID)
		}

		leaseMap[lease.ID] = true

		if err := lease.Validate(); err != nil {
			return err
		}
	}

	return nil
}
