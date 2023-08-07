// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Params struct {
	StatusChangeDelay        time.Duration `protobuf:"bytes,1,opt,name=status_change_delay,json=statusChangeDelay,proto3,stdduration" json:"status_change_delay"`
	ProofVerificationEnabled bool          `protobuf:"varint,2,opt,name=proof_verification_enabled,json=proofVerificationEnabled,proto3" json:"proof_verification_enabled,omitempty"`
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
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0x84, 0x30,
	0x1c, 0x80, 0x5b, 0x87, 0xcb, 0x05, 0x27, 0x39, 0x07, 0x64, 0xe8, 0x11, 0xa7, 0x5b, 0x6c, 0x0d,
	0xae, 0x4e, 0xe7, 0xb9, 0x1b, 0x4c, 0x1c, 0x5c, 0x48, 0x81, 0x52, 0x9a, 0x70, 0x94, 0xd0, 0x42,
	0xbc, 0xb7, 0x70, 0x34, 0xf1, 0x05, 0x7c, 0x14, 0xc6, 0x1b, 0x9d, 0xfc, 0x03, 0x2f, 0x62, 0x28,
	0x87, 0xb9, 0xad, 0xed, 0xef, 0xfb, 0xbe, 0xa4, 0x3f, 0xcb, 0x53, 0xac, 0xd0, 0xa2, 0x60, 0x39,
	0x51, 0x4c, 0x29, 0x21, 0x0b, 0xd2, 0xf8, 0xa4, 0xa4, 0x15, 0xdd, 0x2a, 0x5c, 0x56, 0x52, 0x4b,
	0x7b, 0x31, 0x11, 0xf8, 0x40, 0xe0, 0xc6, 0x77, 0xcf, 0xb9, 0xe4, 0xd2, 0xcc, 0xc9, 0x70, 0x1a,
	0x51, 0x17, 0x71, 0x29, 0x79, 0xce, 0x88, 0xb9, 0x45, 0x75, 0x4a, 0x92, 0xba, 0xa2, 0x7a, 0x50,
	0xcc, 0xcb, 0xe5, 0x3b, 0xb4, 0x66, 0x0f, 0xa6, 0x6d, 0x3f, 0x5a, 0x0b, 0xa5, 0xa9, 0xae, 0x55,
	0x18, 0x67, 0xb4, 0xe0, 0x2c, 0x4c, 0x58, 0x4e, 0x77, 0x0e, 0xf4, 0xe0, 0xea, 0xd4, 0xbf, 0xc0,
	0x63, 0x08, 0x4f, 0x21, 0xbc, 0x39, 0x84, 0xd6, 0xf3, 0xf6, 0x6b, 0x09, 0xde, 0xbe, 0x97, 0x30,
	0x38, 0x1b, 0xfd, 0x3b, 0xa3, 0x6f, 0x06, 0xdb, 0xbe, 0xb5, 0xdc, 0xb2, 0x92, 0x32, 0x0d, 0x1b,
	0x56, 0x89, 0x54, 0xc4, 0x46, 0x09, 0x59, 0x41, 0xa3, 0x9c, 0x25, 0xce, 0x89, 0x07, 0x57, 0xf3,
	0xc0, 0x31, 0xc4, 0xd3, 0x11, 0x70, 0x3f, 0xce, 0xd7, 0x41, 0xfb, 0x8b, 0xc0, 0x47, 0x87, 0x40,
	0xdb, 0x21, 0xb8, 0xef, 0x10, 0xfc, 0xe9, 0x10, 0x7c, 0xed, 0x11, 0xd8, 0xf7, 0x08, 0x7c, 0xf6,
	0x08, 0x3c, 0x5f, 0x73, 0xa1, 0xb3, 0x3a, 0xc2, 0xb1, 0xdc, 0x92, 0x69, 0x2b, 0x57, 0x32, 0x4d,
	0x45, 0x2c, 0x68, 0x4e, 0xb2, 0x3a, 0x22, 0x2f, 0xff, 0x6b, 0xd4, 0xbb, 0x92, 0xa9, 0x68, 0x66,
	0x7e, 0x70, 0xf3, 0x17, 0x00, 0x00, 0xff, 0xff, 0x08, 0x90, 0xb8, 0xeb, 0x67, 0x01, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.StatusChangeDelay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.StatusChangeDelay):])
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
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.StatusChangeDelay)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.StatusChangeDelay, dAtA[iNdEx:postIndex]); err != nil {
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
