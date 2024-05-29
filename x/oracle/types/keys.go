package types

const (
	ModuleName = "oracle"
	PortID     = ModuleName
	StoreKey   = ModuleName
	Version    = "oracle-1"
)

var (
	ParamsKey = []byte{0x00}

	AssetKeyPrefix = []byte{0x10}
)

func AssetKey(denom string) []byte {
	return append(AssetKeyPrefix, []byte(denom)...)
}
