package v3

import (
	"github.com/sentinel-official/hub/v12/x/subscription/types/v2"
)

func (m *Subscription) IsRenewable() bool {
	return !m.RenewalAt.IsZero()
}

func (m *Subscription) MsgCancelRequest() *v2.MsgCancelRequest {
	return &v2.MsgCancelRequest{
		From: m.AccAddress,
		ID:   m.ID,
	}
}

func (m *Subscription) MsgRenewRequest() *MsgRenewRequest {
	return &MsgRenewRequest{
		From:  m.AccAddress,
		ID:    m.ID,
		Denom: m.Price.Denom,
	}
}

func (m *Subscription) Validate() error {
	return nil
}
