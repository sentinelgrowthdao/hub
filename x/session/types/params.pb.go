// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params represents the parameters for the sessions module.
type Params struct {
	// Field 1: Duration for status change delay.
	StatusChangeDelay time.Duration `protobuf:"bytes,1,opt,name=status_change_delay,json=statusChangeDelay,proto3,stdduration" json:"status_change_delay"`
	// Field 2: Flag indicating whether proof verification is enabled.
	ProofVerificationEnabled bool `protobuf:"varint,2,opt,name=proof_verification_enabled,json=proofVerificationEnabled,proto3" json:"proof_verification_enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_98343ae164d22c30, []int{0}
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
	proto.RegisterType((*Params)(nil), "sentinel.session.v2.Params")
}

func init() { proto.RegisterFile("sentinel/session/v2/params.proto", fileDescriptor_98343ae164d22c30) }

var fileDescriptor_98343ae164d22c30 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0x84, 0x30,
	0x1c, 0x80, 0x5b, 0x87, 0xcb, 0x05, 0x27, 0x39, 0x07, 0x64, 0xe8, 0x11, 0xa7, 0x5b, 0x6c, 0x23,
	0x3a, 0x3a, 0x9d, 0xe7, 0x6e, 0xce, 0xe4, 0x06, 0x17, 0x52, 0xa0, 0x94, 0x26, 0x1c, 0x25, 0xb4,
	0x10, 0xef, 0x2d, 0x1c, 0x4d, 0x7c, 0x01, 0x1f, 0x85, 0xf1, 0x46, 0x27, 0xff, 0xc0, 0x8b, 0x18,
	0xca, 0x61, 0x6e, 0x6b, 0xfb, 0xfb, 0xbe, 0x2f, 0xe9, 0xcf, 0xf2, 0x14, 0xcb, 0xb5, 0xc8, 0x59,
	0x46, 0x14, 0x53, 0x4a, 0xc8, 0x9c, 0xd4, 0x3e, 0x29, 0x68, 0x49, 0xb7, 0x0a, 0x17, 0xa5, 0xd4,
	0xd2, 0x9e, 0x8d, 0x04, 0x3e, 0x10, 0xb8, 0xf6, 0xdd, 0x73, 0x2e, 0xb9, 0x34, 0x73, 0xd2, 0x9f,
	0x06, 0xd4, 0x45, 0x5c, 0x4a, 0x9e, 0x31, 0x62, 0x6e, 0x61, 0x95, 0x90, 0xb8, 0x2a, 0xa9, 0xee,
	0x15, 0xf3, 0x72, 0xf9, 0x0e, 0xad, 0xc9, 0xa3, 0x69, 0xdb, 0x4f, 0xd6, 0x4c, 0x69, 0xaa, 0x2b,
	0x15, 0x44, 0x29, 0xcd, 0x39, 0x0b, 0x62, 0x96, 0xd1, 0x9d, 0x03, 0x3d, 0xb8, 0x38, 0xf5, 0x2f,
	0xf0, 0x10, 0xc2, 0x63, 0x08, 0xaf, 0x0e, 0xa1, 0xe5, 0xb4, 0xf9, 0x9a, 0x83, 0xb7, 0xef, 0x39,
	0x5c, 0x9f, 0x0d, 0xfe, 0xbd, 0xd1, 0x57, 0xbd, 0x6d, 0xdf, 0x59, 0x6e, 0x51, 0x4a, 0x99, 0x04,
	0x35, 0x2b, 0x45, 0x22, 0x22, 0xa3, 0x04, 0x2c, 0xa7, 0x61, 0xc6, 0x62, 0xe7, 0xc4, 0x83, 0x8b,
	0xe9, 0xda, 0x31, 0xc4, 0xe6, 0x08, 0x78, 0x18, 0xe6, 0xcb, 0x4d, 0xf3, 0x8b, 0xc0, 0x47, 0x8b,
	0x40, 0xd3, 0x22, 0xb8, 0x6f, 0x11, 0xfc, 0x69, 0x11, 0x7c, 0xed, 0x10, 0xd8, 0x77, 0x08, 0x7c,
	0x76, 0x08, 0x3c, 0xdf, 0x72, 0xa1, 0xd3, 0x2a, 0xc4, 0x91, 0xdc, 0x92, 0x71, 0x2b, 0x57, 0x32,
	0x49, 0x44, 0x24, 0x68, 0x46, 0xd2, 0x2a, 0x24, 0xf5, 0xb5, 0x4f, 0x5e, 0xfe, 0x57, 0xa9, 0x77,
	0x05, 0x53, 0xe1, 0xc4, 0xfc, 0xe2, 0xe6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xcb, 0x6b, 0x30,
	0x6b, 0x01, 0x00, 0x00,
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
	if m.ProofVerificationEnabled {
		i--
		if m.ProofVerificationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.StatusChangeDelay, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.StatusChangeDelay):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
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
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.StatusChangeDelay)
	n += 1 + l + sovParams(uint64(l))
	if m.ProofVerificationEnabled {
		n += 2
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field StatusChangeDelay", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.StatusChangeDelay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofVerificationEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProofVerificationEnabled = bool(v != 0)
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
