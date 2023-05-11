package types

import (
	"fmt"
)

func (l *Lease) Validate() error {
	if l.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if l.Bytes.IsNil() && l.Minutes == 0 {
		return fmt.Errorf("[bytes, minutes] cannot be empty")
	}
	if !l.Bytes.IsNil() && l.Minutes != 0 {
		return fmt.Errorf("[bytes, minutes] cannot be non-empty")
	}
	if !l.Bytes.IsNil() {
		if l.Bytes.IsNegative() {
			return fmt.Errorf("bytes cannot be negative")
		}
		if l.Bytes.IsZero() {
			return fmt.Errorf("bytes cannot be zero")
		}
	}
	if l.Minutes != 0 {
		if l.Minutes < 0 {
			return fmt.Errorf("minutes cannot be negative")
		}
	}
	if l.Price.Denom != "" {
		if l.Price.IsNil() {
			return fmt.Errorf("price cannot be nil")
		}
		if l.Price.IsNegative() {
			return fmt.Errorf("price cannot be negative")
		}
		if l.Price.IsZero() {
			return fmt.Errorf("price cannot be zero")
		}
		if !l.Price.IsValid() {
			return fmt.Errorf("price must be valid")
		}
	}

	return nil
}

type (
	Leases []Lease
)
