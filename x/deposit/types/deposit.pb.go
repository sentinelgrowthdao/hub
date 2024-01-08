// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/deposit/v1/deposit.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Deposit represents a message for handling deposits.
type Deposit struct {
	// Field 1: Deposit address represented as a string.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Field 2: List of coins involved in the deposit.
	// - (gogoproto.nullable) = false: Field is not nullable.
	// - (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins":
	//   Type to cast to when repeating this field.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6348f01112f3831f, []int{0}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Deposit)(nil), "sentinel.deposit.v1.Deposit")
}

func init() { proto.RegisterFile("sentinel/deposit/v1/deposit.proto", fileDescriptor_6348f01112f3831f) }

var fileDescriptor_6348f01112f3831f = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3d, 0x52, 0xf3, 0x30,
	0x10, 0x86, 0xa5, 0xef, 0x1b, 0xc8, 0x60, 0x3a, 0x43, 0x61, 0x52, 0x28, 0x81, 0xca, 0x4d, 0xb4,
	0x38, 0x70, 0x82, 0xc0, 0x09, 0x52, 0x50, 0xd0, 0xf9, 0x47, 0x71, 0x34, 0x24, 0x5e, 0x4f, 0x56,
	0xf1, 0xc0, 0x05, 0xa8, 0x39, 0x06, 0xc3, 0x49, 0x5c, 0xa6, 0xa4, 0xe2, 0xc7, 0xbe, 0x08, 0x63,
	0x2b, 0x62, 0xa8, 0xf4, 0x6a, 0x47, 0xfb, 0x3c, 0xda, 0xf5, 0xce, 0x49, 0x15, 0x46, 0x17, 0x6a,
	0x05, 0x99, 0x2a, 0x91, 0xb4, 0x81, 0x2a, 0x72, 0x51, 0x96, 0x1b, 0x34, 0xe8, 0x9f, 0xb8, 0x27,
	0xd2, 0xd5, 0xab, 0x68, 0x28, 0x52, 0xa4, 0x35, 0x12, 0x24, 0x31, 0x29, 0xa8, 0xa2, 0x44, 0x99,
	0x38, 0x82, 0x14, 0x75, 0x61, 0x9b, 0x86, 0xa7, 0x39, 0xe6, 0xd8, 0x47, 0xe8, 0x92, 0xad, 0x5e,
	0x3c, 0x73, 0x6f, 0x70, 0x6b, 0x21, 0x7e, 0xe0, 0x0d, 0xe2, 0x2c, 0xdb, 0x28, 0xa2, 0x80, 0x8f,
	0x79, 0x78, 0x34, 0x77, 0x57, 0x3f, 0xf6, 0x0e, 0x3a, 0x12, 0x05, 0xff, 0xc6, 0xff, 0xc3, 0xe3,
	0xe9, 0x99, 0xb4, 0x2e, 0xd9, 0xb9, 0xe4, 0xde, 0x25, 0x6f, 0x50, 0x17, 0xb3, 0xcb, 0xfa, 0x63,
	0xc4, 0xde, 0x3e, 0x47, 0x61, 0xae, 0xcd, 0x72, 0x9b, 0xc8, 0x14, 0xd7, 0xb0, 0xff, 0x98, 0x3d,
	0x26, 0x94, 0x3d, 0x80, 0x79, 0x2a, 0x15, 0xf5, 0x0d, 0x34, 0xb7, 0xe4, 0xd9, 0x5d, 0xfd, 0x2d,
	0xd8, 0x6b, 0x23, 0x58, 0xdd, 0x08, 0xbe, 0x6b, 0x04, 0xff, 0x6a, 0x04, 0x7f, 0x69, 0x05, 0xdb,
	0xb5, 0x82, 0xbd, 0xb7, 0x82, 0xdd, 0x5f, 0xff, 0x41, 0xba, 0x05, 0x4c, 0x70, 0xb1, 0xd0, 0xa9,
	0x8e, 0x57, 0xb0, 0xdc, 0x26, 0x50, 0x45, 0x53, 0x78, 0xfc, 0x5d, 0x5b, 0x2f, 0x49, 0x0e, 0xfb,
	0x39, 0xaf, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff, 0x98, 0xa4, 0x4f, 0x57, 0x01, 0x00, 0x00,
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDeposit(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovDeposit(uint64(l))
		}
	}
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
