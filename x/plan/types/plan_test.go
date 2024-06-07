package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestPlan_GetProviderAddress(t *testing.T) {
	type fields struct {
		ProviderAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   base.ProvAddress
	}{
		{
			"empty",
			fields{
				ProviderAddress: base.TestAddrEmpty,
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
			},
			base.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				ProviderAddress: tt.fields.ProviderAddress,
			}
			if got := p.GetProviderAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProviderAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlan_Price(t *testing.T) {
	type fields struct {
		Prices sdk.Coins
	}
	type args struct {
		denom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"nil prices and empty denom",
			fields{
				Prices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"nil prices and one denom",
			fields{
				Prices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"nil prices and two denom",
			fields{
				Prices: base.TestCoinsNil,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"empty prices and empty denom",
			fields{
				Prices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"empty prices and one denom",
			fields{
				Prices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"empty prices and two denom",
			fields{
				Prices: base.TestCoinsEmpty,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"1000one prices and empty denom",
			fields{
				Prices: base.TestCoinsPositiveAmount,
			},
			args{
				denom: base.TestDenomEmpty,
			},
			base.TestCoinEmpty,
			false,
		},
		{
			"1000one prices and one denom",
			fields{
				Prices: base.TestCoinsPositiveAmount,
			},
			args{
				denom: base.TestDenomOne,
			},
			base.TestCoinPositiveAmount,
			true,
		},
		{
			"1000one prices and two denom",
			fields{
				Prices: base.TestCoinsPositiveAmount,
			},
			args{
				denom: base.TestDenomTwo,
			},
			base.TestCoinEmpty,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				Prices: tt.fields.Prices,
			}
			got, got1 := p.Price(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Price() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Price() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPlan_Validate(t *testing.T) {
	type fields struct {
		ID              uint64
		ProviderAddress string
		Duration        time.Duration
		Gigabytes       int64
		Prices          sdk.Coins
		Status          base.Status
		StatusAt        time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"id zero",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"id positive",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"provider_address empty",
			fields{
				ID:              1000,
				ProviderAddress: base.TestAddrEmpty,
			},
			true,
		},
		{
			"provider_address invalid",
			fields{
				ID:              1000,
				ProviderAddress: base.TestAddrInvalid,
			},
			true,
		},
		{
			"provider_address invalid prefix",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"provider_address 10 bytes",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr10Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"provider_address 20 bytes",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"provider_address 30 bytes",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr30Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"duration negative",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        0,
			},
			true,
		},
		{
			"duration positive",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"gigabytes negative",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       -1000,
			},
			true,
		},
		{
			"gigabytes zero",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       0,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"prices nil",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsNil,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			true,
		},
		{
			"prices empty",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsEmpty,
			},
			true,
		},
		{
			"prices empty denom",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"prices empty amount",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"prices invalid denom",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"prices negative amount",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"prices zero amount",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"prices positive amount",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"status inactive pending",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusInactive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
		{
			"status_at zero",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeZero,
			},
			true,
		},
		{
			"status_at positive",
			fields{
				ID:              1000,
				ProviderAddress: base.TestBech32ProvAddr20Bytes,
				Duration:        1000,
				Gigabytes:       1000,
				Prices:          base.TestCoinsPositiveAmount,
				Status:          base.StatusActive,
				StatusAt:        base.TestTimeNow,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				ID:              tt.fields.ID,
				ProviderAddress: tt.fields.ProviderAddress,
				Duration:        tt.fields.Duration,
				Gigabytes:       tt.fields.Gigabytes,
				Prices:          tt.fields.Prices,
				Status:          tt.fields.Status,
				StatusAt:        tt.fields.StatusAt,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
