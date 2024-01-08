// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/payout.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Payout represents information about a payout.
type Payout struct {
	// Field 1: Unique identifier for the payout.
	ID uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Field 2: Address associated with the payout.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Field 3: Node address associated with the payout.
	NodeAddress string `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	// Field 4: Duration, in hours, for which the payout is calculated.
	Hours int64 `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	// Field 5: Price of the payout, represented as a Cosmos Coin.
	// This field is not nullable.
	Price types.Coin `protobuf:"bytes,5,opt,name=price,proto3" json:"price"`
	// Field 6: Timestamp indicating when the next payout is scheduled.
	// This field is not nullable and is represented using the standard Timestamp format.
	NextAt time.Time `protobuf:"bytes,6,opt,name=next_at,json=nextAt,proto3,stdtime" json:"next_at"`
}

func (m *Payout) Reset()         { *m = Payout{} }
func (m *Payout) String() string { return proto.CompactTextString(m) }
func (*Payout) ProtoMessage()    {}
func (*Payout) Descriptor() ([]byte, []int) {
	return fileDescriptor_f20b0a92fdfecfe5, []int{0}
}
func (m *Payout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payout.Merge(m, src)
}
func (m *Payout) XXX_Size() int {
	return m.Size()
}
func (m *Payout) XXX_DiscardUnknown() {
	xxx_messageInfo_Payout.DiscardUnknown(m)
}

var xxx_messageInfo_Payout proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Payout)(nil), "sentinel.subscription.v2.Payout")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v2/payout.proto", fileDescriptor_f20b0a92fdfecfe5)
}

var fileDescriptor_f20b0a92fdfecfe5 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbb, 0x8e, 0x9b, 0x40,
	0x14, 0x86, 0x19, 0x6c, 0xe3, 0x64, 0x9c, 0x0a, 0x59, 0x11, 0x71, 0x31, 0x90, 0x48, 0x91, 0x68,
	0x32, 0x23, 0x13, 0xa5, 0x48, 0x91, 0xc2, 0x24, 0x4d, 0xba, 0x08, 0xa5, 0x4a, 0x0a, 0x8b, 0xcb,
	0x18, 0x8f, 0x64, 0x38, 0x88, 0x19, 0x90, 0xfd, 0x16, 0x7e, 0x8c, 0x3c, 0x8a, 0x4b, 0x97, 0xa9,
	0xbc, 0xbb, 0xb8, 0xda, 0xb7, 0x58, 0x01, 0x46, 0xda, 0xed, 0xce, 0xe5, 0xfb, 0xa5, 0xf3, 0xe9,
	0xe0, 0x8f, 0x92, 0xe7, 0x4a, 0xe4, 0x7c, 0xc7, 0x64, 0x15, 0xc9, 0xb8, 0x14, 0x85, 0x12, 0x90,
	0xb3, 0xda, 0x63, 0x45, 0x78, 0x80, 0x4a, 0xd1, 0xa2, 0x04, 0x05, 0xa6, 0x35, 0x60, 0xf4, 0x39,
	0x46, 0x6b, 0x6f, 0x41, 0x62, 0x90, 0x19, 0x48, 0x16, 0x85, 0x92, 0xb3, 0x7a, 0x19, 0x71, 0x15,
	0x2e, 0x59, 0x0c, 0x22, 0xef, 0x93, 0x8b, 0x79, 0x0a, 0x29, 0x74, 0x25, 0x6b, 0xab, 0xdb, 0xd4,
	0x4e, 0x01, 0xd2, 0x1d, 0x67, 0x5d, 0x17, 0x55, 0x1b, 0xa6, 0x44, 0xc6, 0xa5, 0x0a, 0xb3, 0xa2,
	0x07, 0x3e, 0x3c, 0x22, 0x6c, 0xfc, 0xea, 0x2e, 0x30, 0xdf, 0x62, 0x5d, 0x24, 0x16, 0x72, 0x90,
	0x3b, 0xf6, 0x8d, 0xe6, 0x62, 0xeb, 0x3f, 0x7f, 0x04, 0xba, 0x48, 0x4c, 0x0b, 0x4f, 0xc3, 0x24,
	0x29, 0xb9, 0x94, 0x96, 0xee, 0x20, 0xf7, 0x75, 0x30, 0xb4, 0xe6, 0x7b, 0xfc, 0x26, 0x87, 0x84,
	0xaf, 0x87, 0xf5, 0xa8, 0x5b, 0xcf, 0xda, 0xd9, 0xea, 0x86, 0xcc, 0xf1, 0x64, 0x0b, 0x55, 0x29,
	0xad, 0xb1, 0x83, 0xdc, 0x51, 0xd0, 0x37, 0xe6, 0x17, 0x3c, 0x29, 0x4a, 0x11, 0x73, 0x6b, 0xe2,
	0x20, 0x77, 0xe6, 0xbd, 0xa3, 0xbd, 0x1c, 0x6d, 0xe5, 0xe8, 0x4d, 0x8e, 0x7e, 0x07, 0x91, 0xfb,
	0xe3, 0xd3, 0xc5, 0xd6, 0x82, 0x9e, 0x36, 0xbf, 0xe1, 0x69, 0xce, 0xf7, 0x6a, 0x1d, 0x2a, 0xcb,
	0xe8, 0x82, 0x0b, 0xda, 0xfb, 0xd1, 0xc1, 0x8f, 0xfe, 0x1e, 0xfc, 0xfc, 0x57, 0x6d, 0xf2, 0x78,
	0x67, 0xa3, 0xc0, 0x68, 0x43, 0x2b, 0xe5, 0xff, 0x3d, 0x3d, 0x10, 0xed, 0x5f, 0x43, 0xb4, 0x53,
	0x43, 0xd0, 0xb9, 0x21, 0xe8, 0xbe, 0x21, 0xe8, 0x78, 0x25, 0xda, 0xf9, 0x4a, 0xb4, 0xff, 0x57,
	0xa2, 0xfd, 0xf9, 0x9a, 0x0a, 0xb5, 0xad, 0x22, 0x1a, 0x43, 0xc6, 0x86, 0x4f, 0x7c, 0x82, 0xcd,
	0x46, 0xc4, 0x22, 0xdc, 0xb1, 0x6d, 0x15, 0xb1, 0x7a, 0xe9, 0xb1, 0xfd, 0xcb, 0x1f, 0xaa, 0x43,
	0xc1, 0x65, 0x64, 0x74, 0x27, 0x7c, 0x7e, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x56, 0x7d, 0x4e, 0xc2,
	0xe9, 0x01, 0x00, 0x00,
}

func (m *Payout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.NextAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NextAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPayout(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPayout(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Hours != 0 {
		i = encodeVarintPayout(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintPayout(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPayout(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintPayout(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPayout(dAtA []byte, offset int, v uint64) int {
	offset -= sovPayout(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Payout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovPayout(uint64(m.ID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPayout(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovPayout(uint64(l))
	}
	if m.Hours != 0 {
		n += 1 + sovPayout(uint64(m.Hours))
	}
	l = m.Price.Size()
	n += 1 + l + sovPayout(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.NextAt)
	n += 1 + l + sovPayout(uint64(l))
	return n
}

func sovPayout(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPayout(x uint64) (n int) {
	return sovPayout(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Payout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayout
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
			return fmt.Errorf("proto: Payout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
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
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayout
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
				return ErrInvalidLengthPayout
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPayout
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.NextAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayout(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPayout
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
func skipPayout(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayout
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
					return 0, ErrIntOverflowPayout
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
					return 0, ErrIntOverflowPayout
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
				return 0, ErrInvalidLengthPayout
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPayout
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPayout
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPayout        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayout          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPayout = fmt.Errorf("proto: unexpected end of group")
)
