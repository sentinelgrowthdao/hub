// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/sentinel-official/hub/types"
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

type EventRegister struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventRegister) Reset()         { *m = EventRegister{} }
func (m *EventRegister) String() string { return proto.CompactTextString(m) }
func (*EventRegister) ProtoMessage()    {}
func (*EventRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{0}
}
func (m *EventRegister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegister.Merge(m, src)
}
func (m *EventRegister) XXX_Size() int {
	return m.Size()
}
func (m *EventRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegister.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegister proto.InternalMessageInfo

type EventUpdateDetails struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventUpdateDetails) Reset()         { *m = EventUpdateDetails{} }
func (m *EventUpdateDetails) String() string { return proto.CompactTextString(m) }
func (*EventUpdateDetails) ProtoMessage()    {}
func (*EventUpdateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{1}
}
func (m *EventUpdateDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateDetails.Merge(m, src)
}
func (m *EventUpdateDetails) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateDetails proto.InternalMessageInfo

type EventUpdateStatus struct {
	Address string       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Status  types.Status `protobuf:"varint,2,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
}

func (m *EventUpdateStatus) Reset()         { *m = EventUpdateStatus{} }
func (m *EventUpdateStatus) String() string { return proto.CompactTextString(m) }
func (*EventUpdateStatus) ProtoMessage()    {}
func (*EventUpdateStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{2}
}
func (m *EventUpdateStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateStatus.Merge(m, src)
}
func (m *EventUpdateStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateStatus proto.InternalMessageInfo

type EventLease struct {
	ID     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	Lessor string `protobuf:"bytes,2,opt,name=lessor,proto3" json:"lessor,omitempty" yaml:"lessor"`
	Lessee string `protobuf:"bytes,3,opt,name=lessee,proto3" json:"lessee,omitempty" yaml:"lessee"`
}

func (m *EventLease) Reset()         { *m = EventLease{} }
func (m *EventLease) String() string { return proto.CompactTextString(m) }
func (*EventLease) ProtoMessage()    {}
func (*EventLease) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{3}
}
func (m *EventLease) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventLease) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventLease.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventLease) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLease.Merge(m, src)
}
func (m *EventLease) XXX_Size() int {
	return m.Size()
}
func (m *EventLease) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLease.DiscardUnknown(m)
}

var xxx_messageInfo_EventLease proto.InternalMessageInfo

type EventCreateSubscription struct {
	ID          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	NodeAddress string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
}

func (m *EventCreateSubscription) Reset()         { *m = EventCreateSubscription{} }
func (m *EventCreateSubscription) String() string { return proto.CompactTextString(m) }
func (*EventCreateSubscription) ProtoMessage()    {}
func (*EventCreateSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{4}
}
func (m *EventCreateSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateSubscription.Merge(m, src)
}
func (m *EventCreateSubscription) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateSubscription proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRegister)(nil), "sentinel.node.v2.EventRegister")
	proto.RegisterType((*EventUpdateDetails)(nil), "sentinel.node.v2.EventUpdateDetails")
	proto.RegisterType((*EventUpdateStatus)(nil), "sentinel.node.v2.EventUpdateStatus")
	proto.RegisterType((*EventLease)(nil), "sentinel.node.v2.EventLease")
	proto.RegisterType((*EventCreateSubscription)(nil), "sentinel.node.v2.EventCreateSubscription")
}

func init() { proto.RegisterFile("sentinel/node/v2/events.proto", fileDescriptor_d02d789f2b133ff1) }

var fileDescriptor_d02d789f2b133ff1 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8a, 0xd3, 0x40,
	0x1c, 0xc6, 0x93, 0x28, 0x95, 0x8e, 0xee, 0xe2, 0x66, 0x85, 0xad, 0x0b, 0x4e, 0x96, 0xf1, 0xb2,
	0x82, 0xcd, 0xd0, 0x7a, 0x2b, 0x78, 0x30, 0xd6, 0x83, 0x20, 0x08, 0x11, 0x2f, 0x5e, 0x24, 0x69,
	0xfe, 0x4d, 0x07, 0xd2, 0x4c, 0xc8, 0x4c, 0x82, 0xf5, 0x05, 0xbc, 0xfa, 0x18, 0x3e, 0x4a, 0x8f,
	0x3d, 0x7a, 0x0a, 0x9a, 0xbe, 0x41, 0x9e, 0x40, 0x32, 0x93, 0x94, 0x1e, 0x3c, 0x6c, 0x6f, 0x99,
	0x7c, 0xdf, 0xff, 0x37, 0xdf, 0x7c, 0xfc, 0xd1, 0x33, 0x01, 0xa9, 0x64, 0x29, 0x24, 0x34, 0xe5,
	0x11, 0xd0, 0x72, 0x4a, 0xa1, 0x84, 0x54, 0x0a, 0x37, 0xcb, 0xb9, 0xe4, 0xf6, 0xe3, 0x5e, 0x76,
	0x5b, 0xd9, 0x2d, 0xa7, 0xd7, 0x4f, 0x62, 0x1e, 0x73, 0x25, 0xd2, 0xf6, 0x4b, 0xfb, 0xae, 0xf1,
	0x01, 0x23, 0x37, 0x19, 0x08, 0x5a, 0x4e, 0xa8, 0x90, 0x81, 0x2c, 0x3a, 0x0e, 0x79, 0x8d, 0xce,
	0xde, 0xb5, 0x5c, 0x1f, 0x62, 0x26, 0x24, 0xe4, 0xf6, 0x4b, 0xf4, 0x20, 0x88, 0xa2, 0x1c, 0x84,
	0x18, 0x99, 0x37, 0xe6, 0xed, 0xd0, 0xb3, 0x9b, 0xca, 0x39, 0xdf, 0x04, 0xeb, 0x64, 0x46, 0x3a,
	0x81, 0xf8, 0xbd, 0x85, 0x78, 0xc8, 0x56, 0xe3, 0x9f, 0xb3, 0x28, 0x90, 0x30, 0x07, 0x19, 0xb0,
	0x44, 0x9c, 0xc8, 0xf8, 0x61, 0xa2, 0x8b, 0x23, 0xc8, 0x27, 0x15, 0xef, 0x34, 0x86, 0x3d, 0x47,
	0x03, 0xfd, 0xac, 0x91, 0x75, 0x63, 0xde, 0x9e, 0x4f, 0x9f, 0xba, 0x87, 0x7e, 0xd4, 0xbb, 0xdd,
	0x72, 0xe2, 0x6a, 0xb0, 0x77, 0xd1, 0x54, 0xce, 0x99, 0xe6, 0xe8, 0x11, 0xe2, 0x77, 0xb3, 0x6d,
	0x12, 0xa4, 0x92, 0x7c, 0x80, 0x40, 0x80, 0xfd, 0x1c, 0x59, 0x2c, 0x52, 0xb7, 0xdf, 0xf7, 0x2e,
	0xeb, 0xca, 0xb1, 0xde, 0xcf, 0x9b, 0xca, 0x19, 0xea, 0x59, 0x16, 0x11, 0xdf, 0x62, 0x91, 0xfd,
	0x02, 0x0d, 0x12, 0x10, 0x82, 0xe7, 0xea, 0xe6, 0xe1, 0x31, 0x5e, 0xff, 0x27, 0x7e, 0x67, 0xe8,
	0xad, 0x00, 0xa3, 0x7b, 0xff, 0xb3, 0x02, 0x74, 0x56, 0x00, 0xf2, 0x1d, 0x5d, 0xa9, 0x20, 0x6f,
	0x73, 0x68, 0x2b, 0x29, 0x42, 0xb1, 0xc8, 0x59, 0x26, 0x19, 0x4f, 0xef, 0x96, 0x6a, 0x86, 0x1e,
	0xb5, 0x7b, 0xf1, 0xb5, 0xaf, 0x50, 0x67, 0xbb, 0x6a, 0x2a, 0xe7, 0x52, 0x1b, 0x8f, 0x55, 0xe2,
	0x3f, 0x6c, 0x8f, 0x6f, 0xf4, 0xc9, 0xfb, 0xb8, 0xfd, 0x8b, 0x8d, 0x5f, 0x35, 0x36, 0xb6, 0x35,
	0x36, 0x77, 0x35, 0x36, 0xff, 0xd4, 0xd8, 0xfc, 0xb9, 0xc7, 0xc6, 0x6e, 0x8f, 0x8d, 0xdf, 0x7b,
	0x6c, 0x7c, 0x19, 0xc7, 0x4c, 0xae, 0x8a, 0xd0, 0x5d, 0xf0, 0x35, 0xed, 0x7b, 0x1e, 0xf3, 0xe5,
	0x92, 0x2d, 0x58, 0x90, 0xd0, 0x55, 0x11, 0xd2, 0x6f, 0x7a, 0x6b, 0x55, 0xf7, 0xe1, 0x40, 0xad,
	0xda, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xd5, 0xfa, 0xb1, 0xd3, 0x02, 0x00, 0x00,
}

func (m *EventRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventLease) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventLease) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventLease) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Lessee) > 0 {
		i -= len(m.Lessee)
		copy(dAtA[i:], m.Lessee)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Lessee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Lessor) > 0 {
		i -= len(m.Lessor)
		copy(dAtA[i:], m.Lessor)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Lessor)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRegister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	return n
}

func (m *EventLease) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	l = len(m.Lessor)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Lessee)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCreateSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRegister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRegister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventLease) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventLease: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventLease: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return fmt.Errorf("proto: wrong wireType = %d for field Lessor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lessor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lessee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lessee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCreateSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
