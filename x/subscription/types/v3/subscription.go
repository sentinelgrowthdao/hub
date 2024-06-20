package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *Subscription) GetAccAddress() sdk.AccAddress {
	if m.AccAddress == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(m.AccAddress)
	if err != nil {
		panic(err)
	}

	return addr
}
