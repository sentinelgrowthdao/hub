// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/subscription/v3/msg.proto

package v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgStartSessionRequest struct {
	From           string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	SubscriptionID uint64 `protobuf:"varint,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	NodeAddress    string `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
}

func (m *MsgStartSessionRequest) Reset()         { *m = MsgStartSessionRequest{} }
func (m *MsgStartSessionRequest) String() string { return proto.CompactTextString(m) }
func (*MsgStartSessionRequest) ProtoMessage()    {}
func (*MsgStartSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b821cd74f392424, []int{0}
}
func (m *MsgStartSessionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartSessionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartSessionRequest.Merge(m, src)
}
func (m *MsgStartSessionRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartSessionRequest proto.InternalMessageInfo

type MsgStartSessionResponse struct {
}

func (m *MsgStartSessionResponse) Reset()         { *m = MsgStartSessionResponse{} }
func (m *MsgStartSessionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStartSessionResponse) ProtoMessage()    {}
func (*MsgStartSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b821cd74f392424, []int{1}
}
func (m *MsgStartSessionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartSessionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartSessionResponse.Merge(m, src)
}
func (m *MsgStartSessionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartSessionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStartSessionRequest)(nil), "sentinel.subscription.v3.MsgStartSessionRequest")
	proto.RegisterType((*MsgStartSessionResponse)(nil), "sentinel.subscription.v3.MsgStartSessionResponse")
}

func init() {
	proto.RegisterFile("sentinel/subscription/v3/msg.proto", fileDescriptor_4b821cd74f392424)
}

var fileDescriptor_4b821cd74f392424 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4e, 0x32, 0x41,
	0x14, 0xc5, 0x77, 0xbe, 0x8f, 0x98, 0x38, 0x1a, 0x48, 0x26, 0x46, 0x57, 0x8a, 0x11, 0xa9, 0x68,
	0x9c, 0x11, 0xb6, 0xd4, 0x46, 0x62, 0x43, 0x61, 0x03, 0x9d, 0x85, 0x84, 0xdd, 0x1d, 0x96, 0x49,
	0x60, 0x66, 0x9d, 0x3b, 0x3b, 0xd1, 0xde, 0xd2, 0xc2, 0xc7, 0xf0, 0x51, 0x28, 0x29, 0xad, 0x8c,
	0x2e, 0x2f, 0x62, 0x58, 0x42, 0xb2, 0xfe, 0x2b, 0xec, 0x6e, 0xee, 0xfd, 0xdd, 0x93, 0x73, 0x72,
	0x70, 0x13, 0x84, 0xb2, 0x52, 0x89, 0x29, 0x87, 0x2c, 0x84, 0xc8, 0xc8, 0xd4, 0x4a, 0xad, 0xb8,
	0x0b, 0xf8, 0x0c, 0x12, 0x96, 0x1a, 0x6d, 0x35, 0xf1, 0x37, 0x0c, 0x2b, 0x33, 0xcc, 0x05, 0xf5,
	0xbd, 0x44, 0x27, 0xba, 0x80, 0xf8, 0x6a, 0x5a, 0xf3, 0xcd, 0x47, 0x84, 0xf7, 0xaf, 0x20, 0x19,
	0xd8, 0x91, 0xb1, 0x03, 0x01, 0x20, 0xb5, 0xea, 0x8b, 0xdb, 0x4c, 0x80, 0x25, 0x04, 0x57, 0xc6,
	0x46, 0xcf, 0x7c, 0xd4, 0x40, 0xad, 0xed, 0x7e, 0x31, 0x93, 0x33, 0x5c, 0x2b, 0xeb, 0x0e, 0x65,
	0xec, 0xff, 0x6b, 0xa0, 0x56, 0xa5, 0x4b, 0xf2, 0xd7, 0xa3, 0xea, 0xa0, 0x74, 0xea, 0x5d, 0xf6,
	0xab, 0x65, 0xb4, 0x17, 0x93, 0x63, 0xbc, 0xab, 0x74, 0x2c, 0x86, 0xa3, 0x38, 0x36, 0x02, 0xc0,
	0xff, 0x5f, 0x08, 0xef, 0xac, 0x76, 0x17, 0xeb, 0x55, 0xf3, 0x10, 0x1f, 0x7c, 0x73, 0x03, 0xa9,
	0x56, 0x20, 0x3a, 0x0f, 0x08, 0xe3, 0xd5, 0x4d, 0x18, 0x27, 0x23, 0x41, 0x1c, 0xae, 0x7d, 0x21,
	0xc9, 0x29, 0xfb, 0x2d, 0x3c, 0xfb, 0x39, 0x62, 0xbd, 0xfd, 0x87, 0x8f, 0xb5, 0x8d, 0xee, 0xcd,
	0xfc, 0x9d, 0x7a, 0xcf, 0x39, 0xf5, 0xe6, 0x39, 0x45, 0x8b, 0x9c, 0xa2, 0xb7, 0x9c, 0xa2, 0xa7,
	0x25, 0xf5, 0x16, 0x4b, 0xea, 0xbd, 0x2c, 0xa9, 0x77, 0x7d, 0x9e, 0x48, 0x3b, 0xc9, 0x42, 0x16,
	0xe9, 0x19, 0xdf, 0xc8, 0x9f, 0xe8, 0xf1, 0x58, 0x46, 0x72, 0x34, 0xe5, 0x93, 0x2c, 0xe4, 0xae,
	0xdd, 0xe1, 0x77, 0x9f, 0x4b, 0xb4, 0xf7, 0xa9, 0x00, 0xee, 0x82, 0x70, 0xab, 0xe8, 0x25, 0xf8,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x01, 0x3b, 0x92, 0x9b, 0xed, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgStartSession(ctx context.Context, in *MsgStartSessionRequest, opts ...grpc.CallOption) (*MsgStartSessionResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgStartSession(ctx context.Context, in *MsgStartSessionRequest, opts ...grpc.CallOption) (*MsgStartSessionResponse, error) {
	out := new(MsgStartSessionResponse)
	err := c.cc.Invoke(ctx, "/sentinel.subscription.v3.MsgService/MsgStartSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgStartSession(context.Context, *MsgStartSessionRequest) (*MsgStartSessionResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgStartSession(ctx context.Context, req *MsgStartSessionRequest) (*MsgStartSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgStartSession not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgStartSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgStartSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.subscription.v3.MsgService/MsgStartSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgStartSession(ctx, req.(*MsgStartSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.subscription.v3.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgStartSession",
			Handler:    _MsgService_MsgStartSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/subscription/v3/msg.proto",
}

func (m *MsgStartSessionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartSessionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartSessionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SubscriptionID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.SubscriptionID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStartSessionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartSessionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartSessionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStartSessionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovMsg(uint64(m.SubscriptionID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgStartSessionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStartSessionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgStartSessionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartSessionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgStartSessionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgStartSessionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartSessionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)