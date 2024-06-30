// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/session.proto

package v2

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	v1 "github.com/sentinel-official/hub/v12/types/v1"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Session represents a session.
type Session struct {
	// Field 1: Session ID.
	ID uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Field 2: Subscription ID.
	SubscriptionID uint64 `protobuf:"varint,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	// Field 3: Node address.
	NodeAddress string `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	// Field 4: Account address.
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// Field 5: Bandwidth details.
	Bandwidth v1.Bandwidth `protobuf:"bytes,5,opt,name=bandwidth,proto3" json:"bandwidth"`
	// Field 6: Session duration.
	Duration time.Duration `protobuf:"bytes,6,opt,name=duration,proto3,stdduration" json:"duration"`
	// Field 7: Inactive timestamp.
	InactiveAt time.Time `protobuf:"bytes,7,opt,name=inactive_at,json=inactiveAt,proto3,stdtime" json:"inactive_at"`
	// Field 8: Session status.
	Status v1.Status `protobuf:"varint,8,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	// Field 9: Status timestamp.
	StatusAt time.Time `protobuf:"bytes,9,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa3a03f78d221e53, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Session)(nil), "sentinel.session.v2.Session")
}

func init() { proto.RegisterFile("sentinel/session/v2/session.proto", fileDescriptor_aa3a03f78d221e53) }

var fileDescriptor_aa3a03f78d221e53 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0xc7, 0x93, 0x5e, 0xe9, 0x8b, 0x8b, 0x8a, 0x64, 0x10, 0xca, 0x55, 0xc8, 0xe9, 0x31, 0x75,
	0xc1, 0x56, 0xc3, 0xc0, 0xc0, 0x00, 0x8d, 0xca, 0xd0, 0x35, 0x65, 0x81, 0xa5, 0x4a, 0x6a, 0x37,
	0xb5, 0xd4, 0xc6, 0x55, 0xed, 0x04, 0xf8, 0x16, 0x37, 0xf2, 0x11, 0xf8, 0x28, 0x1d, 0x6f, 0x44,
	0x42, 0x2a, 0x90, 0x7e, 0x11, 0x14, 0xbb, 0xce, 0x9d, 0x38, 0x16, 0xb6, 0x27, 0xcf, 0xff, 0xf7,
	0xfc, 0x9f, 0x97, 0x18, 0x5c, 0x49, 0x96, 0x29, 0x9e, 0xb1, 0x0d, 0x91, 0x4c, 0x4a, 0x2e, 0x32,
	0x52, 0x04, 0x36, 0xc4, 0xbb, 0xbd, 0x50, 0x02, 0x3e, 0xb6, 0x08, 0xb6, 0xf9, 0x22, 0x18, 0x3c,
	0x49, 0x45, 0x2a, 0xb4, 0x4e, 0xaa, 0xc8, 0xa0, 0x03, 0x94, 0x0a, 0x91, 0x6e, 0x18, 0xd1, 0x5f,
	0x49, 0xbe, 0x22, 0x34, 0xdf, 0xc7, 0xaa, 0xb6, 0x1a, 0xf8, 0x7f, 0xeb, 0x8a, 0x6f, 0x99, 0x54,
	0xf1, 0x76, 0x77, 0x06, 0x6e, 0xc7, 0x51, 0x5f, 0x76, 0x4c, 0x92, 0x62, 0x4c, 0x92, 0x38, 0xa3,
	0x9f, 0x38, 0x55, 0x6b, 0xdb, 0xe3, 0x3e, 0x22, 0x55, 0xac, 0x72, 0x69, 0xf4, 0xe7, 0x3f, 0x2e,
	0x40, 0x7b, 0x6e, 0x06, 0x85, 0x4f, 0x41, 0x83, 0x53, 0xcf, 0x1d, 0xba, 0xa3, 0x66, 0xd8, 0x2a,
	0x8f, 0x7e, 0x63, 0x36, 0x8d, 0x1a, 0x9c, 0xc2, 0xd7, 0xe0, 0x91, 0xcc, 0x13, 0xb9, 0xdc, 0xf3,
	0x5d, 0x35, 0xdd, 0x82, 0x53, 0xaf, 0xa1, 0x21, 0x58, 0x1e, 0xfd, 0xfe, 0xfc, 0x8e, 0x34, 0x9b,
	0x46, 0xfd, 0xbb, 0xe8, 0x8c, 0xc2, 0x2b, 0xf0, 0x30, 0x13, 0x94, 0x2d, 0x62, 0x4a, 0xf7, 0x4c,
	0x4a, 0xef, 0x62, 0xe8, 0x8e, 0xba, 0x51, 0xaf, 0xca, 0x4d, 0x4c, 0x0a, 0x7a, 0xa0, 0x6d, 0xd5,
	0xa6, 0x56, 0xed, 0x27, 0x7c, 0x0b, 0xba, 0xf5, 0x42, 0xde, 0x83, 0xa1, 0x3b, 0xea, 0x05, 0xcf,
	0x70, 0x7d, 0x60, 0xbd, 0x11, 0x2e, 0xc6, 0x38, 0xb4, 0x4c, 0xd8, 0x3c, 0x1c, 0x7d, 0x27, 0xba,
	0x2d, 0x82, 0x6f, 0x40, 0xc7, 0x5e, 0xd5, 0x6b, 0x69, 0x83, 0x4b, 0x6c, 0xce, 0x8a, 0xed, 0x59,
	0xf1, 0xf4, 0x0c, 0x84, 0x9d, 0xaa, 0xfa, 0xeb, 0x4f, 0xdf, 0x8d, 0xea, 0x22, 0xf8, 0x0e, 0xf4,
	0x78, 0x16, 0x2f, 0x15, 0x2f, 0xd8, 0x22, 0x56, 0x5e, 0x5b, 0x7b, 0x0c, 0xee, 0x79, 0xbc, 0xb7,
	0xbf, 0xc6, 0x98, 0x5c, 0x57, 0x26, 0xc0, 0x16, 0x4e, 0x14, 0x1c, 0x83, 0x96, 0xb9, 0xbb, 0xd7,
	0x19, 0xba, 0xa3, 0x7e, 0x70, 0xf9, 0x8f, 0x35, 0xe6, 0x1a, 0x88, 0xce, 0x20, 0x9c, 0x80, 0xae,
	0x89, 0xaa, 0xbe, 0xdd, 0xff, 0xe8, 0xdb, 0x31, 0x65, 0x13, 0x15, 0x7e, 0x38, 0xfc, 0x46, 0xce,
	0xb7, 0x12, 0x39, 0x87, 0x12, 0xb9, 0x37, 0x25, 0x72, 0x7f, 0x95, 0xc8, 0xbd, 0x3e, 0x21, 0xe7,
	0xe6, 0x84, 0x9c, 0xef, 0x27, 0xe4, 0x7c, 0x7c, 0x95, 0x72, 0xb5, 0xce, 0x13, 0xbc, 0x14, 0x5b,
	0x62, 0x27, 0x7a, 0x21, 0x56, 0x2b, 0xbe, 0xe4, 0xf1, 0x86, 0xac, 0xf3, 0x84, 0x14, 0xe3, 0x80,
	0x7c, 0xae, 0xdf, 0xfb, 0xf9, 0x15, 0x05, 0x49, 0x4b, 0x8f, 0xf0, 0xf2, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb1, 0x0c, 0xb2, 0xa0, 0x13, 0x03, 0x00, 0x00,
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSession(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	if m.Status != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.InactiveAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintSession(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x3a
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintSession(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Bandwidth.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSession(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSession(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintSession(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SubscriptionID != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x10
	}
	if m.ID != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	offset -= sovSession(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovSession(uint64(m.ID))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovSession(uint64(m.SubscriptionID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	l = m.Bandwidth.Size()
	n += 1 + l + sovSession(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovSession(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt)
	n += 1 + l + sovSession(uint64(l))
	if m.Status != 0 {
		n += 1 + sovSession(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovSession(uint64(l))
	return n
}

func sovSession(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriptionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bandwidth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bandwidth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.InactiveAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= v1.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StatusAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
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
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
				return 0, ErrInvalidLengthSession
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSession
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSession
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSession        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSession = fmt.Errorf("proto: unexpected end of group")
)