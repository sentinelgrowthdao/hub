// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/oracle/v1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("sentinel/oracle/v1/msg.proto", fileDescriptor_b67b95f425ef3bb9) }

var fileDescriptor_b67b95f425ef3bb9 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4b, 0xcd, 0xd1, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x33, 0xd4, 0xcf,
	0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xc9, 0xea, 0x41, 0x64, 0xf5,
	0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xd2, 0xfa, 0x20, 0x16, 0x44, 0xa5, 0x11,
	0x0f, 0x17, 0x97, 0x6f, 0x71, 0x7a, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x53, 0xe8, 0x89,
	0x87, 0x72, 0x0c, 0x2b, 0x1e, 0xc9, 0x31, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3,
	0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c,
	0x43, 0x94, 0x71, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xcc, 0x12,
	0xdd, 0xfc, 0xb4, 0xb4, 0xcc, 0xe4, 0xcc, 0xc4, 0x1c, 0xfd, 0x8c, 0xd2, 0x24, 0xfd, 0x32, 0x43,
	0x23, 0xfd, 0x0a, 0x98, 0xab, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x76, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x5f, 0x44, 0x97, 0xb5, 0x00, 0x00, 0x00,
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
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.oracle.v1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "sentinel/oracle/v1/msg.proto",
}