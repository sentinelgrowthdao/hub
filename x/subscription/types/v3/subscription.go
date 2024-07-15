package v3

import (
	"time"
)

func (m *Subscription) RenewalAt() time.Time {
	if m.Renewable {
		return m.InactiveAt
	}

	return time.Time{}
}

func (m *Subscription) Validate() error {
	return nil
}
