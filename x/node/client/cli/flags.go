package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
)

const (
	flagGigabytePrices = "gigabyte-prices"
	flagHourlyPrices   = "hourly-prices"
	flagNodeAddr       = "node-addr"
	flagPlanID         = "plan-id"
	flagProvAddr       = "prov-addr"
	flagRemoteURL      = "remote-url"
	flagRenewable      = "renewable"
)

func GetGigabytePrices(flags *pflag.FlagSet) (sdk.Coins, error) {
	s, err := flags.GetString(flagGigabytePrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseCoinsNormalized(s)
}

func GetHourlyPrices(flags *pflag.FlagSet) (sdk.Coins, error) {
	s, err := flags.GetString(flagHourlyPrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseCoinsNormalized(s)
}
