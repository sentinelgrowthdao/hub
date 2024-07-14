// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/node.proto

package v2

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	v1 "github.com/sentinel-official/hub/v12/types/v1"
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

type Node struct {
	Address        string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	GigabytePrices github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=gigabyte_prices,json=gigabytePrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"gigabyte_prices"`
	HourlyPrices   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=hourly_prices,json=hourlyPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"hourly_prices"`
	RemoteURL      string                                   `protobuf:"bytes,4,opt,name=remote_url,json=remoteUrl,proto3" json:"remote_url,omitempty"`
	InactiveAt     time.Time                                `protobuf:"bytes,5,opt,name=inactive_at,json=inactiveAt,proto3,stdtime" json:"inactive_at"`
	Status         v1.Status                                `protobuf:"varint,6,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
	StatusAt       time.Time                                `protobuf:"bytes,7,opt,name=status_at,json=statusAt,proto3,stdtime" json:"status_at"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_f03f9fca82881dc6, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Node)(nil), "sentinel.node.v2.Node")
}

func init() { proto.RegisterFile("sentinel/node/v2/node.proto", fileDescriptor_f03f9fca82881dc6) }

var fileDescriptor_f03f9fca82881dc6 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x13, 0xa6, 0x4c, 0x3b, 0x1e, 0x5a, 0x50, 0xc4, 0x22, 0x1d, 0x24, 0x67, 0xc4, 0x2a,
	0x0b, 0x6a, 0x93, 0xc0, 0x05, 0x26, 0x88, 0x1d, 0x42, 0x28, 0x50, 0x16, 0x6c, 0x46, 0x4e, 0xe2,
	0xc9, 0x58, 0x24, 0x71, 0x14, 0x3b, 0x11, 0x73, 0x8b, 0x1e, 0x03, 0x71, 0x92, 0x59, 0x76, 0xc9,
	0xaa, 0x85, 0xcc, 0x11, 0xb8, 0x00, 0x8a, 0x1d, 0x57, 0x6c, 0x91, 0xba, 0xf2, 0xf3, 0xf3, 0x7b,
	0xfa, 0x7e, 0xbf, 0xff, 0x81, 0x67, 0x82, 0x56, 0x92, 0x55, 0xb4, 0xc0, 0x15, 0xcf, 0x28, 0xee,
	0x42, 0x75, 0xa2, 0xba, 0xe1, 0x92, 0x3b, 0x4f, 0xcc, 0x23, 0x52, 0xc9, 0x2e, 0x5c, 0xc0, 0x94,
	0x8b, 0x92, 0x0b, 0x9c, 0x10, 0x41, 0x71, 0x17, 0x24, 0x54, 0x92, 0x00, 0xa7, 0x9c, 0x55, 0xba,
	0x63, 0xf1, 0x34, 0xe7, 0x39, 0x57, 0x21, 0x1e, 0xa2, 0x31, 0xeb, 0xe5, 0x9c, 0xe7, 0x05, 0xc5,
	0xea, 0x96, 0xb4, 0x1b, 0x2c, 0x59, 0x49, 0x85, 0x24, 0x65, 0x3d, 0x16, 0xc0, 0x3b, 0x15, 0x72,
	0x57, 0x53, 0x81, 0xbb, 0x00, 0x0b, 0x49, 0x64, 0x2b, 0xf4, 0xfb, 0xf3, 0x3f, 0x13, 0x70, 0xf4,
	0x9e, 0x67, 0xd4, 0x71, 0xc1, 0x31, 0xc9, 0xb2, 0x86, 0x0a, 0xe1, 0xda, 0x4b, 0xdb, 0x9f, 0xc5,
	0xe6, 0xea, 0x48, 0xf0, 0x38, 0x67, 0x39, 0x49, 0x76, 0x92, 0xae, 0xeb, 0x86, 0xa5, 0x54, 0xb8,
	0x0f, 0x96, 0x13, 0x7f, 0x1e, 0x9e, 0x23, 0xad, 0x19, 0x0d, 0x9a, 0xd1, 0xa8, 0x19, 0xbd, 0xe1,
	0xac, 0x8a, 0x5e, 0xee, 0x6f, 0x3c, 0xeb, 0xc7, 0xad, 0xe7, 0xe7, 0x4c, 0x6e, 0xdb, 0x04, 0xa5,
	0xbc, 0xc4, 0xe3, 0x07, 0xf5, 0x71, 0x21, 0xb2, 0xaf, 0x5a, 0x92, 0x6a, 0x10, 0xf1, 0x99, 0x61,
	0x7c, 0x50, 0x08, 0xa7, 0x06, 0xa7, 0x5b, 0xde, 0x36, 0xc5, 0xce, 0x30, 0x27, 0xf7, 0xcf, 0x7c,
	0xa4, 0x09, 0x23, 0xf1, 0x05, 0x00, 0x0d, 0x2d, 0xb9, 0xa4, 0xeb, 0xb6, 0x29, 0xdc, 0xa3, 0x61,
	0x08, 0xd1, 0x69, 0x7f, 0xe3, 0xcd, 0x62, 0x95, 0xbd, 0x8c, 0xdf, 0xc5, 0x33, 0x5d, 0x70, 0xd9,
	0x14, 0xce, 0x5b, 0x30, 0x67, 0x15, 0x49, 0x25, 0xeb, 0xe8, 0x9a, 0x48, 0xf7, 0xe1, 0xd2, 0xf6,
	0xe7, 0xe1, 0x02, 0x69, 0x3f, 0x90, 0xf1, 0x03, 0x7d, 0x32, 0x7e, 0x44, 0x27, 0x83, 0xbc, 0xab,
	0x5b, 0xcf, 0x8e, 0x81, 0x69, 0x5c, 0x49, 0x27, 0x00, 0x53, 0xed, 0x87, 0x3b, 0x5d, 0xda, 0xfe,
	0x59, 0x78, 0x8e, 0xee, 0x36, 0x43, 0x2b, 0xed, 0x02, 0xf4, 0x51, 0x15, 0xc4, 0x63, 0xa1, 0xb3,
	0x02, 0x33, 0x1d, 0x0d, 0xdc, 0xe3, 0xff, 0xe0, 0x9e, 0xe8, 0xb6, 0x95, 0x8c, 0x3e, 0xef, 0x7f,
	0x43, 0xeb, 0x7b, 0x0f, 0xad, 0x7d, 0x0f, 0xed, 0xeb, 0x1e, 0xda, 0xbf, 0x7a, 0x68, 0x5f, 0x1d,
	0xa0, 0x75, 0x7d, 0x80, 0xd6, 0xcf, 0x03, 0xb4, 0xbe, 0xbc, 0xfe, 0x67, 0x88, 0x46, 0xd1, 0x05,
	0xdf, 0x6c, 0x58, 0xca, 0x48, 0x81, 0xb7, 0x6d, 0x82, 0xbb, 0x20, 0xc4, 0xdf, 0xf4, 0x6e, 0x8f,
	0xab, 0x15, 0x26, 0x53, 0xc5, 0x7f, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xee, 0x05, 0xe4,
	0xfc, 0x02, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StatusAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintNode(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if m.Status != 0 {
		i = encodeVarintNode(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.InactiveAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintNode(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if len(m.RemoteURL) > 0 {
		i -= len(m.RemoteURL)
		copy(dAtA[i:], m.RemoteURL)
		i = encodeVarintNode(dAtA, i, uint64(len(m.RemoteURL)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HourlyPrices) > 0 {
		for iNdEx := len(m.HourlyPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HourlyPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.GigabytePrices) > 0 {
		for iNdEx := len(m.GigabytePrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GigabytePrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNode(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if len(m.GigabytePrices) > 0 {
		for _, e := range m.GigabytePrices {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	if len(m.HourlyPrices) > 0 {
		for _, e := range m.HourlyPrices {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	l = len(m.RemoteURL)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.InactiveAt)
	n += 1 + l + sovNode(uint64(l))
	if m.Status != 0 {
		n += 1 + sovNode(uint64(m.Status))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StatusAt)
	n += 1 + l + sovNode(uint64(l))
	return n
}

func sovNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GigabytePrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GigabytePrices = append(m.GigabytePrices, types.Coin{})
			if err := m.GigabytePrices[len(m.GigabytePrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HourlyPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HourlyPrices = append(m.HourlyPrices, types.Coin{})
			if err := m.HourlyPrices[len(m.HourlyPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemoteURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemoteURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactiveAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.InactiveAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
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
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNode = fmt.Errorf("proto: unexpected end of group")
)
