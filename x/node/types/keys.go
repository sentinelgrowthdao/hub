package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkaddress "github.com/cosmos/cosmos-sdk/types/address"

	base "github.com/sentinel-official/hub/v12/types"
)

const (
	ModuleName = "node"
)

var (
	LeaseCountKey = []byte{0x00}

	NodeKeyPrefix              = []byte{0x10}
	ActiveNodeKeyPrefix        = append(NodeKeyPrefix, 0x01)
	InactiveNodeKeyPrefix      = append(NodeKeyPrefix, 0x02)
	NodeForInactiveAtKeyPrefix = []byte{0x11}
	NodeForPlanKeyPrefix       = []byte{0x12}

	LeaseKeyPrefix                  = []byte{0x21}
	LeaseForInactiveAtKeyPrefix     = []byte{0x22}
	LeaseForNodeKeyPrefix           = []byte{0x23}
	LeaseForProviderKeyPrefix       = []byte{0x24}
	LeaseForProviderByNodeKeyPrefix = []byte{0x25}

	LeasePayoutForNextAtKeyPrefix = []byte{0x31}
)

func ActiveNodeKey(addr base.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func InactiveNodeKey(addr base.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetNodeForPlanKeyPrefix(id uint64) []byte {
	return append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func NodeForPlanKey(id uint64, addr base.NodeAddress) []byte {
	return append(GetNodeForPlanKeyPrefix(id), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func GetNodeForInactiveAtKeyPrefix(at time.Time) []byte {
	return append(NodeForInactiveAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func NodeForInactiveAtKey(at time.Time, addr base.NodeAddress) []byte {
	return append(GetNodeForInactiveAtKeyPrefix(at), sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func LeaseKey(id uint64) []byte {
	return append(LeaseKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForInactiveAtKeyPrefix(at time.Time) []byte {
	return append(LeaseForInactiveAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func LeaseForInactiveAtKey(at time.Time, id uint64) []byte {
	return append(GetLeaseForInactiveAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForNodeKeyPrefix(addr base.NodeAddress) []byte {
	return append(LeaseForNodeKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func LeaseForNodeKey(addr base.NodeAddress, id uint64) []byte {
	return append(GetLeaseForNodeKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForProviderKeyPrefix(addr base.ProvAddress) []byte {
	return append(LeaseForProviderKeyPrefix, sdkaddress.MustLengthPrefix(addr.Bytes())...)
}

func LeaseForProviderKey(addr base.ProvAddress, id uint64) []byte {
	return append(GetLeaseForProviderKeyPrefix(addr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeaseForProviderByNodeKeyPrefix(provAddr base.ProvAddress, nodeAddr base.NodeAddress) (key []byte) {
	return append(append(LeaseForProviderByNodeKeyPrefix, sdkaddress.MustLengthPrefix(provAddr.Bytes())...), sdkaddress.MustLengthPrefix(nodeAddr.Bytes())...)
}

func LeaseForProviderByNodeKey(provAddr base.ProvAddress, nodeAddr base.NodeAddress, id uint64) []byte {
	return append(GetLeaseForProviderByNodeKeyPrefix(provAddr, nodeAddr), sdk.Uint64ToBigEndian(id)...)
}

func GetLeasePayoutForNextAtKeyPrefix(at time.Time) []byte {
	return append(LeasePayoutForNextAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func LeasePayoutForNextAtKey(at time.Time, id uint64) []byte {
	return append(GetLeasePayoutForNextAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
}

func AddressFromNodeForPlanKey(key []byte) base.NodeAddress {
	// prefix (1 byte) | id (8 bytes) | addrLen (1 byte) | addr (addrLen bytes)

	addrLen := int(key[9])
	if len(key) != 10+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 10+addrLen))
	}

	return key[10:]
}

func AddressFromNodeForInactiveAtKey(key []byte) base.NodeAddress {
	// prefix (1 byte) | at (29 bytes) | addrLen (1 byte) | addr (addrLen bytes)

	addrLen := int(key[30])
	if len(key) != 31+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 31+addrLen))
	}

	return key[31:]
}

func IDFromLeaseForProviderByNodeKey(key []byte) uint64 {
	// prefix (1 byte) | provAddrLen(1 byte) | provAddr (provAddrLen bytes) | nodeAddrLen(1 byte) | nodeAddr (nodeAddrLen bytes) | id (8 bytes)

	provAddrLen, nodeAddrLen := int(key[1]), int(key[2+int(key[1])])
	if len(key) != 11+provAddrLen+nodeAddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 11+provAddrLen+nodeAddrLen))
	}

	return sdk.BigEndianToUint64(key[3+provAddrLen+nodeAddrLen:])
}

func IDFromLeasePayoutForNextAtKey(key []byte) uint64 {
	// prefix (1 byte) | at (29 bytes) | id (8 bytes)

	if len(key) != 38 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 38))
	}

	return sdk.BigEndianToUint64(key[30:])
}
