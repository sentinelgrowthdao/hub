package cli

import (
	"github.com/spf13/pflag"

	base "github.com/sentinel-official/hub/v12/types"
)

const (
	flagProvider = "provider"
)

func GetProvider(flags *pflag.FlagSet) (base.ProvAddress, error) {
	s, err := flags.GetString(flagProvider)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return base.ProvAddressFromBech32(s)
}
