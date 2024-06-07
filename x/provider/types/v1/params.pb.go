// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/provider/v1/params.proto

package v1types

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

type Params struct {
	Deposit      types.Coin                             `protobuf:"bytes,1,opt,name=deposit,proto3" json:"deposit"`
	StakingShare github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=staking_share,json=stakingShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_share"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4f14af8b7b64362, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "sentinel.provider.v1.Params")
}

func init() { proto.RegisterFile("sentinel/provider/v1/params.proto", fileDescriptor_f4f14af8b7b64362) }

var fileDescriptor_f4f14af8b7b64362 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0x42, 0x31,
	0x14, 0xc6, 0x5b, 0x63, 0x30, 0x5e, 0x75, 0x21, 0x0c, 0xc8, 0x50, 0xd0, 0xc1, 0xb0, 0xd0, 0x93,
	0xe2, 0x64, 0x5c, 0x0c, 0xfa, 0x00, 0x06, 0x36, 0x17, 0xd3, 0x5e, 0xca, 0xa5, 0x01, 0xee, 0xb9,
	0xb9, 0x2d, 0x8d, 0xbe, 0x85, 0xa3, 0x8f, 0xe0, 0xa3, 0x30, 0x32, 0x1a, 0x07, 0xa2, 0x97, 0x17,
	0x31, 0xf7, 0x5f, 0x74, 0xea, 0x97, 0xe6, 0x77, 0x7e, 0xe7, 0xe4, 0x0b, 0x2e, 0xac, 0x8e, 0x9d,
	0x89, 0xf5, 0x12, 0x92, 0x14, 0xbd, 0x99, 0xea, 0x14, 0xbc, 0x80, 0x44, 0xa6, 0x72, 0x65, 0x79,
	0x92, 0xa2, 0xc3, 0x66, 0xab, 0x46, 0x78, 0x8d, 0x70, 0x2f, 0x3a, 0x2c, 0x44, 0xbb, 0x42, 0x0b,
	0x4a, 0x5a, 0x0d, 0x5e, 0x28, 0xed, 0xa4, 0x80, 0x10, 0x4d, 0x5c, 0x4e, 0x75, 0x5a, 0x11, 0x46,
	0x58, 0x44, 0xc8, 0x53, 0xf9, 0x7b, 0xf9, 0x4e, 0x83, 0xc6, 0x63, 0x21, 0x6f, 0xde, 0x04, 0x47,
	0x53, 0x9d, 0xa0, 0x35, 0xae, 0x4d, 0x7b, 0xb4, 0x7f, 0x32, 0x3c, 0xe7, 0xa5, 0x92, 0xe7, 0x4a,
	0x5e, 0x29, 0xf9, 0x3d, 0x9a, 0x78, 0x74, 0xb8, 0xd9, 0x75, 0xc9, 0xb8, 0xe6, 0x9b, 0x93, 0xe0,
	0xcc, 0x3a, 0xb9, 0x30, 0x71, 0xf4, 0x6c, 0xe7, 0x32, 0xd5, 0xed, 0x83, 0x1e, 0xed, 0x1f, 0x8f,
	0x78, 0x4e, 0x7d, 0xed, 0xba, 0x57, 0x91, 0x71, 0xf3, 0xb5, 0xe2, 0x21, 0xae, 0xa0, 0xba, 0xb2,
	0x7c, 0x06, 0x76, 0xba, 0x00, 0xf7, 0x9a, 0x68, 0xcb, 0x1f, 0x74, 0x38, 0x3e, 0xad, 0x24, 0x93,
	0xdc, 0x31, 0x52, 0x9b, 0x1f, 0x46, 0x3e, 0x32, 0x46, 0x36, 0x19, 0xa3, 0xdb, 0x8c, 0xd1, 0xef,
	0x8c, 0xd1, 0xb7, 0x3d, 0x23, 0xdb, 0x3d, 0x23, 0x9f, 0x7b, 0x46, 0x9e, 0xee, 0xfe, 0x79, 0xeb,
	0x4e, 0x06, 0x38, 0x9b, 0x99, 0xd0, 0xc8, 0x25, 0xcc, 0xd7, 0x0a, 0xbc, 0x18, 0xc2, 0xcb, 0x5f,
	0x93, 0xc5, 0x2a, 0xf0, 0xe2, 0xd6, 0x8b, 0x22, 0xaa, 0x46, 0xd1, 0xc2, 0xf5, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0x4e, 0x3a, 0x52, 0x76, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.StakingShare.Size()
		i -= size
		if _, err := m.StakingShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Deposit.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.StakingShare.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StakingShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
