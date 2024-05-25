package types

import (
	"fmt"
)

const (
	ModuleName = "oracle"
	PortID     = ModuleName
	Version    = "oracle-1"
)

var (
	PortIDKey = []byte{0x00}

	AssetKeyPrefix           = []byte{0x10}
	AssetPriceKeyPrefix      = []byte{0x20}
	DenomForRequestKeyPrefix = []byte{0x30, 0x00}
)

func AssetKey(denom string) []byte {
	return append(AssetKeyPrefix, []byte(denom)...)
}

func AssetPriceKey(denom string) []byte {
	return append(AssetPriceKeyPrefix, []byte(denom)...)
}

func DenomForRequestKey(sequence uint64, index int) []byte {
	s := fmt.Sprintf("%d/%d", sequence, index)
	return append(DenomForRequestKeyPrefix, []byte(s)...)
}
