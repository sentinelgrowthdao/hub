package v2

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	base "github.com/sentinel-official/hub/v12/types"
)

func TestMsgCreateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From      string
		Duration  time.Duration
		Gigabytes int64
		Prices    sdk.Coins
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
				From:      base.TestBech32ProvAddr10Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:      base.TestBech32ProvAddr30Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"duration negative",
			fields{
				From:     base.TestBech32ProvAddr20Bytes,
				Duration: -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				From:     base.TestBech32ProvAddr20Bytes,
				Duration: 0,
			},
			true,
		},
		{
			"duration positive",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"gigabytes negative",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: -1000,
			},
			true,
		},
		{
			"gigabytes zero",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 0,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"prices nil",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsNil,
			},
			true,
		},
		{
			"prices empty",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsEmpty,
			},
			true,
		},
		{
			"prices empty denom",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"prices empty amount",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"prices invalid denom",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"prices negative amount",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"prices zero amount",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"prices positive amount",
			fields{
				From:      base.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    base.TestCoinsPositiveAmount,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCreateRequest{
				From:      tt.fields.From,
				Duration:  tt.fields.Duration,
				Gigabytes: tt.fields.Gigabytes,
				Prices:    tt.fields.Prices,
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
		ID     uint64
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
				From:   base.TestBech32ProvAddr10Bytes,
				ID:     1000,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:   base.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:   base.TestBech32ProvAddr30Bytes,
				ID:     1000,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: base.TestBech32ProvAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:   base.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				From:   base.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: base.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				From:   base.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: base.StatusActive,
			},
			false,
		},
		{
			"status inactive pending",
			fields{
				From:   base.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: base.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   base.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: base.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateStatusRequest{
				From:   tt.fields.From,
				ID:     tt.fields.ID,
				Status: tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgLinkNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
		ID          uint64
		NodeAddress string
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
				From:        base.TestBech32ProvAddr10Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:        base.TestBech32ProvAddr30Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: base.TestBech32ProvAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address empty",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestAddrEmpty,
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestAddrInvalid,
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgLinkNodeRequest{
				From:        tt.fields.From,
				ID:          tt.fields.ID,
				NodeAddress: tt.fields.NodeAddress,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUnlinkNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
		ID          uint64
		NodeAddress string
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
				From:        base.TestBech32ProvAddr10Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:        base.TestBech32ProvAddr30Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: base.TestBech32ProvAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address empty",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestAddrEmpty,
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestAddrInvalid,
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				From:        base.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: base.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUnlinkNodeRequest{
				From:        tt.fields.From,
				ID:          tt.fields.ID,
				NodeAddress: tt.fields.NodeAddress,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgSubscribeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From  string
		ID    uint64
		Denom string
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
				From: base.TestBech32ProvAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:  base.TestBech32AccAddr10Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:  base.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:  base.TestBech32AccAddr30Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"id zero",
			fields{
				From: base.TestBech32AccAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:  base.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"denom empty",
			fields{
				From:  base.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "",
			},
			true,
		},
		{
			"denom invalid",
			fields{
				From:  base.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "o",
			},
			true,
		},
		{
			"denom one",
			fields{
				From:  base.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeRequest{
				From:  tt.fields.From,
				ID:    tt.fields.ID,
				Denom: tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
