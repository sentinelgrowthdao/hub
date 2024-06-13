// DO NOT COVER

package types

import (
	"github.com/spf13/pflag"

	v1 "github.com/sentinel-official/hub/v12/types/v1"
)

const (
	FlagStatus = "status"
)

func StatusFromFlags(flags *pflag.FlagSet) (v1.Status, error) {
	s, err := flags.GetString(FlagStatus)
	if err != nil {
		return v1.StatusUnspecified, err
	}

	return v1.StatusFromString(s), nil
}
