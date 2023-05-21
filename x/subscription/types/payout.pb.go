// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v2/payout.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Payout struct {
	ID        uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hours     int64      `protobuf:"varint,2,opt,name=hours,proto3" json:"hours,omitempty"`
	Price     types.Coin `protobuf:"bytes,3,opt,name=price,proto3" json:"price"`
	Timestamp time.Time  `protobuf:"bytes,4,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
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
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x80, 0x33, 0xe9, 0x0f, 0x1a, 0x77, 0xa1, 0x48, 0xec, 0x62, 0x52, 0x04, 0xa1, 0x1b, 0x67,
	0x68, 0xa5, 0x17, 0x88, 0x6e, 0xdc, 0x49, 0x10, 0x04, 0x77, 0x99, 0x74, 0x9a, 0x3e, 0x68, 0xf2,
	0x42, 0x66, 0x52, 0xec, 0x2d, 0x7a, 0x0c, 0xb7, 0xde, 0xa2, 0xcb, 0x2e, 0x5d, 0x55, 0x4d, 0x2f,
	0x22, 0xcd, 0x34, 0xfe, 0xec, 0xde, 0x1b, 0xbe, 0x0f, 0x3e, 0xe6, 0x39, 0x57, 0x4a, 0x66, 0x1a,
	0x32, 0xb9, 0xe0, 0xaa, 0x14, 0x2a, 0x2e, 0x20, 0xd7, 0x80, 0x19, 0x5f, 0x8e, 0x79, 0x1e, 0xad,
	0xb0, 0xd4, 0x2c, 0x2f, 0x50, 0xa3, 0xeb, 0x35, 0x18, 0xfb, 0x8b, 0xb1, 0xe5, 0xb8, 0x4f, 0x63,
	0x54, 0x29, 0x2a, 0x2e, 0x22, 0x25, 0xf9, 0x72, 0x24, 0xa4, 0x8e, 0x46, 0x3c, 0x46, 0xc8, 0x8c,
	0xd9, 0xef, 0x25, 0x98, 0x60, 0x3d, 0xf2, 0xc3, 0x74, 0x7c, 0xf5, 0x13, 0xc4, 0x64, 0x21, 0x79,
	0xbd, 0x89, 0x72, 0xc6, 0x35, 0xa4, 0x52, 0xe9, 0x28, 0xcd, 0x0d, 0x70, 0xf9, 0x46, 0x9c, 0xee,
	0x43, 0x5d, 0xe0, 0x9e, 0x3b, 0x36, 0x4c, 0x3d, 0x32, 0x20, 0xc3, 0x76, 0xd0, 0xad, 0x76, 0xbe,
	0x7d, 0x7f, 0x17, 0xda, 0x30, 0x75, 0x7b, 0x4e, 0x67, 0x8e, 0x65, 0xa1, 0x3c, 0x7b, 0x40, 0x86,
	0xad, 0xd0, 0x2c, 0xee, 0xc4, 0xe9, 0xe4, 0x05, 0xc4, 0xd2, 0x6b, 0x0d, 0xc8, 0xf0, 0x6c, 0x7c,
	0xc1, 0x4c, 0x1f, 0x3b, 0xf4, 0xb1, 0x63, 0x1f, 0xbb, 0x45, 0xc8, 0x82, 0xf6, 0x66, 0xe7, 0x5b,
	0xa1, 0xa1, 0xdd, 0xc0, 0x39, 0xfd, 0x49, 0xf0, 0xda, 0xb5, 0xda, 0x67, 0x26, 0x92, 0x35, 0x91,
	0xec, 0xb1, 0x21, 0x82, 0x93, 0x83, 0xbb, 0xfe, 0xf0, 0x49, 0xf8, 0xab, 0x05, 0x4f, 0x9b, 0x2f,
	0x6a, 0xbd, 0x56, 0xd4, 0xda, 0x54, 0x94, 0x6c, 0x2b, 0x4a, 0x3e, 0x2b, 0x4a, 0xd6, 0x7b, 0x6a,
	0x6d, 0xf7, 0xd4, 0x7a, 0xdf, 0x53, 0xeb, 0x79, 0x92, 0x80, 0x9e, 0x97, 0x82, 0xc5, 0x98, 0xf2,
	0xe6, 0x47, 0xaf, 0x71, 0x36, 0x83, 0x18, 0xa2, 0x05, 0x9f, 0x97, 0x82, 0xbf, 0xfc, 0xbf, 0x83,
	0x5e, 0xe5, 0x52, 0x89, 0x6e, 0x5d, 0x70, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x97, 0x4c, 0x39,
	0x55, 0xad, 0x01, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPayout(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPayout(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Hours != 0 {
		i = encodeVarintPayout(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x10
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
	if m.Hours != 0 {
		n += 1 + sovPayout(uint64(m.Hours))
	}
	l = m.Price.Size()
	n += 1 + l + sovPayout(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
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
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
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