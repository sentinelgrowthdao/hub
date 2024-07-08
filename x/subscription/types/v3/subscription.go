package v3

func (m *Subscription) IsRenewable() bool {
	return !m.RenewalAt.IsZero()
}

func (m *Subscription) Validate() error {
	return nil
}
