// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisSubscription represents the initial state for a subscription in the genesis block.
type GenesisSubscription struct {
	// Field 1: Subscription information stored as a serialized Any message.
	Subscription *types.Any `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// Field 2: Allocations associated with the subscription.
	// Each allocation contains information about granted and utilized bytes.
	Allocations []Allocation `protobuf:"bytes,2,rep,name=allocations,proto3" json:"allocations"`
}

func (m *GenesisSubscription) Reset()         { *m = GenesisSubscription{} }
func (m *GenesisSubscription) String() string { return proto.CompactTextString(m) }
func (*GenesisSubscription) ProtoMessage()    {}
func (*GenesisSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_794e3328e7a9d1e0, []int{0}
}
func (m *GenesisSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisSubscription.Merge(m, src)
}
func (m *GenesisSubscription) XXX_Size() int {
	return m.Size()
}
func (m *GenesisSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisSubscription proto.InternalMessageInfo

// GenesisState represents the initial state of the module in the genesis block.
type GenesisState struct {
	// Field 1: Subscriptions in the genesis block.
	// Each GenesisSubscription contains subscription information and associated allocations.
	Subscriptions []GenesisSubscription `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions"`
	// Field 2: Parameters for the module stored in the genesis block.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_794e3328e7a9d1e0, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisSubscription)(nil), "sentinel.subscription.v2.GenesisSubscription")
	proto.RegisterType((*GenesisState)(nil), "sentinel.subscription.v2.GenesisState")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v2/genesis.proto", fileDescriptor_794e3328e7a9d1e0)
}

var fileDescriptor_794e3328e7a9d1e0 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0xe3, 0x82, 0x3a, 0xb8, 0x65, 0x09, 0x1d, 0x42, 0x07, 0x53, 0x55, 0x80, 0xca, 0x50,
	0x5b, 0x84, 0x05, 0x16, 0xa4, 0x76, 0x61, 0x61, 0x40, 0x65, 0x02, 0x26, 0x27, 0x72, 0x5d, 0x4b,
	0xa9, 0x1d, 0xd5, 0x4e, 0x45, 0x6f, 0xc1, 0x05, 0xd8, 0xe1, 0x26, 0x1d, 0x3b, 0x32, 0x21, 0x48,
	0x2f, 0x82, 0x70, 0x12, 0x48, 0x24, 0xb2, 0xd9, 0xf2, 0xf7, 0xbf, 0xf7, 0xfe, 0x67, 0x78, 0xa2,
	0x99, 0x34, 0x42, 0xb2, 0x88, 0xe8, 0x24, 0xd0, 0xe1, 0x42, 0xc4, 0x46, 0x28, 0x49, 0x96, 0x3e,
	0xe1, 0x4c, 0x32, 0x2d, 0x34, 0x8e, 0x17, 0xca, 0x28, 0xd7, 0x2b, 0x38, 0x5c, 0xe6, 0xf0, 0xd2,
	0xef, 0x76, 0xb8, 0xe2, 0xca, 0x42, 0xe4, 0xe7, 0x94, 0xf1, 0xdd, 0x03, 0xae, 0x14, 0x8f, 0x18,
	0xb1, 0xb7, 0x20, 0x99, 0x12, 0x2a, 0x57, 0xf9, 0xd3, 0x69, 0xad, 0x25, 0x8d, 0x22, 0x15, 0x52,
	0x2b, 0x9c, 0xa1, 0xc7, 0xb5, 0x68, 0x4c, 0x17, 0x74, 0x9e, 0x87, 0xeb, 0xbf, 0x00, 0xb8, 0x7f,
	0x9d, 0xc5, 0xbd, 0x2b, 0x71, 0xee, 0x05, 0x6c, 0x97, 0xe7, 0x3c, 0xd0, 0x03, 0x83, 0x96, 0xdf,
	0xc1, 0x59, 0x36, 0x5c, 0x64, 0xc3, 0x23, 0xb9, 0x9a, 0x54, 0x48, 0xf7, 0x06, 0xb6, 0xfe, 0xc2,
	0x68, 0xaf, 0xd1, 0xdb, 0x19, 0xb4, 0xfc, 0x23, 0x5c, 0x57, 0x02, 0x1e, 0xfd, 0xc2, 0xe3, 0xdd,
	0xf5, 0xc7, 0xa1, 0x33, 0x29, 0x8f, 0xf7, 0xdf, 0x00, 0x6c, 0x17, 0xf9, 0x0c, 0x35, 0xcc, 0xbd,
	0x87, 0x7b, 0x65, 0x05, 0xed, 0x01, 0x6b, 0x30, 0xac, 0x37, 0xf8, 0x67, 0xbd, 0xdc, 0xa9, 0xaa,
	0xe4, 0x5e, 0xc1, 0x66, 0xd6, 0x8d, 0xd7, 0xb0, 0xdb, 0xf6, 0xea, 0x35, 0x6f, 0x2d, 0x97, 0xcb,
	0xe4, 0x53, 0xe3, 0xc7, 0xf5, 0x17, 0x72, 0x5e, 0x53, 0xe4, 0xac, 0x53, 0x04, 0x36, 0x29, 0x02,
	0x9f, 0x29, 0x02, 0xcf, 0x5b, 0xe4, 0x6c, 0xb6, 0xc8, 0x79, 0xdf, 0x22, 0xe7, 0xe1, 0x92, 0x0b,
	0x33, 0x4b, 0x02, 0x1c, 0xaa, 0x39, 0x29, 0xb4, 0x87, 0x6a, 0x3a, 0x15, 0xa1, 0xa0, 0x11, 0x99,
	0x25, 0x01, 0x59, 0x9e, 0xf9, 0xe4, 0xa9, 0xfa, 0x65, 0x66, 0x15, 0x33, 0x1d, 0x34, 0x6d, 0xe5,
	0xe7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xe8, 0x3a, 0x04, 0x76, 0x02, 0x00, 0x00,
}

func (m *GenesisSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Allocations) > 0 {
		for iNdEx := len(m.Allocations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Allocations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Subscription != nil {
		{
			size, err := m.Subscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Subscriptions) > 0 {
		for iNdEx := len(m.Subscriptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subscriptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscription != nil {
		l = m.Subscription.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Allocations) > 0 {
		for _, e := range m.Allocations {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for _, e := range m.Subscriptions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subscription == nil {
				m.Subscription = &types.Any{}
			}
			if err := m.Subscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allocations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Allocations = append(m.Allocations, Allocation{})
			if err := m.Allocations[len(m.Allocations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, GenesisSubscription{})
			if err := m.Subscriptions[len(m.Subscriptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
