package v1

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func (m *Lease) IsRenewable() bool {
	return !m.RenewalAt.IsZero()
}

func (m *Lease) MsgEndRequest() *MsgEndRequest {
	return &MsgEndRequest{
		From: m.ProvAddress,
		ID:   m.ID,
	}
}

func (m *Lease) MsgRenewRequest() *MsgRenewRequest {
	return &MsgRenewRequest{
		From:  m.ProvAddress,
		ID:    m.ID,
		Hours: m.MaxHours,
		Denom: m.Price.Denom,
	}
}

func (m *Lease) RefundAmount() sdk.Coin {
	hours := m.MaxHours - m.Hours
	return sdk.NewCoin(
		m.Price.Denom,
		m.Price.Amount.MulRaw(hours),
	)
}

func (m *Lease) Validate() error {
	if m.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if m.ProvAddress == "" {
		return fmt.Errorf("prov_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.ProvAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid prov_address %s", m.ProvAddress)
	}
	if m.NodeAddress == "" {
		return fmt.Errorf("node_address cannot be empty")
	}
	if _, err := base.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrapf(err, "invalid node_address %s", m.NodeAddress)
	}
	if !m.Price.IsValid() {
		return fmt.Errorf("price must be valid")
	}
	if m.Price.IsZero() {
		return fmt.Errorf("price cannot be zero")
	}
	if !m.Deposit.IsValid() {
		return fmt.Errorf("deposit must be valid")
	}
	if m.Deposit.IsZero() {
		return fmt.Errorf("deposit cannot be zero")
	}
	if m.Hours <= 0 {
		return fmt.Errorf("hours must be greater than zero")
	}
	if m.MaxHours <= 0 {
		return fmt.Errorf("max_hours must be greater than zero")
	}
	if m.MaxHours < m.Hours {
		return fmt.Errorf("max_hours cannot be less than hours")
	}
	if m.PayoutAt.IsZero() {
		return fmt.Errorf("payout_at cannot be zero")
	}
	if m.InactiveAt.IsZero() && m.RenewalAt.IsZero() {
		return fmt.Errorf("both inactive_at and renewal_at cannot be zero")
	}
	if !m.InactiveAt.IsZero() && !m.RenewalAt.IsZero() {
		return fmt.Errorf("either inactive_at or renewal_at must be zero")
	}

	return nil
}
