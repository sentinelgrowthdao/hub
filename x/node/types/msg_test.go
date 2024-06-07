package types

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestMsgRegisterRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From           string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: base.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: base.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: base.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:           base.TestBech32AccAddr10Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:           base.TestBech32AccAddr30Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsNil,
			},
			true,
		},
		{
			"gigabyte_prices empty",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsEmpty,
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			true,
		},
		{
			"hourly_prices empty",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsEmpty,
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "",
			},
			true,
		},
		{
			"remote_url 72 chars",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:           base.TestBech32AccAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRegisterRequest{
				From:           tt.fields.From,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateDetailsRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From           string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: base.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: base.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:           base.TestBech32NodeAddr10Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:           base.TestBech32NodeAddr30Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices empty",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsEmpty,
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: base.TestCoinsPositiveAmount,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices empty",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsEmpty,
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   base.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			false,
		},
		{
			"remote_url 72 chars",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:           base.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateDetailsRequest{
				From:           tt.fields.From,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateStatusRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From   string
		Status base.Status
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: base.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: base.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:   base.TestBech32NodeAddr10Bytes,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:   base.TestBech32NodeAddr20Bytes,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:   base.TestBech32NodeAddr30Bytes,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				From:   base.TestBech32NodeAddr20Bytes,
				Status: base.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				From:   base.TestBech32NodeAddr20Bytes,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				From:   base.TestBech32NodeAddr20Bytes,
				Status: base.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   base.TestBech32NodeAddr20Bytes,
				Status: base.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateStatusRequest{
				From:   tt.fields.From,
				Status: tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgSubscribeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
		NodeAddress string
		Hours       int64
		Gigabytes   int64
		Denom       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: base.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: base.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: base.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:        base.TestBech32AccAddr10Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:        base.TestBech32AccAddr30Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"node_address empty",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestAddrEmpty,
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestAddrInvalid,
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr10Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr30Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"hours negative and gigabytes negative",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       -1000,
				Gigabytes:   -1000,
			},
			true,
		},
		{
			"hours zero and gigabytes zero",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       0,
				Gigabytes:   0,
			},
			true,
		},
		{
			"hours positive and gigabytes positive",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   1000,
			},
			true,
		},
		{
			"hours negative",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       -1000,
			},
			true,
		},
		{
			"hours positive",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Denom:       base.TestDenomOne,
			},
			false,
		},
		{
			"gigabytes negative",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       0,
				Gigabytes:   -1000,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       0,
				Gigabytes:   1000,
				Denom:       base.TestDenomOne,
			},
			false,
		},
		{
			"denom empty",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "",
			},
			true,
		},
		{
			"denom invalid",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "o",
			},
			true,
		},
		{
			"denom valid",
			fields{
				From:        base.TestBech32AccAddr20Bytes,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeRequest{
				From:        tt.fields.From,
				NodeAddress: tt.fields.NodeAddress,
				Hours:       tt.fields.Hours,
				Gigabytes:   tt.fields.Gigabytes,
				Denom:       tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
