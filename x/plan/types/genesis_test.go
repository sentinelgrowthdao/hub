package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	var (
		state *GenesisState
	)

	state = DefaultGenesisState()
	require.Equal(t, &GenesisState{Plans: nil}, state)
}

func TestNewGenesisState(t *testing.T) {
	var (
		plans GenesisPlans
		state *GenesisState
	)

	state = NewGenesisState(nil)
	require.Equal(t, &GenesisState{Plans: nil}, state)
	require.Len(t, state.Plans, 0)

	state = NewGenesisState(plans)
	require.Equal(t, &GenesisState{Plans: plans}, state)
	require.Len(t, state.Plans, 0)

	plans = append(plans,
		GenesisPlan{},
	)
	state = NewGenesisState(plans)
	require.Equal(t, &GenesisState{Plans: plans}, state)
	require.Len(t, state.Plans, 1)

	plans = append(plans,
		GenesisPlan{},
	)
	state = NewGenesisState(plans)
	require.Equal(t, &GenesisState{Plans: plans}, state)
	require.Len(t, state.Plans, 2)
}
