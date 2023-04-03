// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/provider/v2/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryProvidersRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryProvidersRequest) Reset()         { *m = QueryProvidersRequest{} }
func (m *QueryProvidersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProvidersRequest) ProtoMessage()    {}
func (*QueryProvidersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb0e6f64d0260de0, []int{0}
}
func (m *QueryProvidersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProvidersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProvidersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProvidersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProvidersRequest.Merge(m, src)
}
func (m *QueryProvidersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProvidersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProvidersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProvidersRequest proto.InternalMessageInfo

type QueryProviderRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryProviderRequest) Reset()         { *m = QueryProviderRequest{} }
func (m *QueryProviderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProviderRequest) ProtoMessage()    {}
func (*QueryProviderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb0e6f64d0260de0, []int{1}
}
func (m *QueryProviderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProviderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProviderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProviderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProviderRequest.Merge(m, src)
}
func (m *QueryProviderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProviderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProviderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProviderRequest proto.InternalMessageInfo

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb0e6f64d0260de0, []int{2}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryProvidersResponse struct {
	Providers  []Provider          `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryProvidersResponse) Reset()         { *m = QueryProvidersResponse{} }
func (m *QueryProvidersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProvidersResponse) ProtoMessage()    {}
func (*QueryProvidersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb0e6f64d0260de0, []int{3}
}
func (m *QueryProvidersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProvidersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProvidersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProvidersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProvidersResponse.Merge(m, src)
}
func (m *QueryProvidersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProvidersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProvidersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProvidersResponse proto.InternalMessageInfo

type QueryProviderResponse struct {
	Provider Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider"`
}

func (m *QueryProviderResponse) Reset()         { *m = QueryProviderResponse{} }
func (m *QueryProviderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProviderResponse) ProtoMessage()    {}
func (*QueryProviderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb0e6f64d0260de0, []int{4}
}
func (m *QueryProviderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProviderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProviderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProviderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProviderResponse.Merge(m, src)
}
func (m *QueryProviderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProviderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProviderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProviderResponse proto.InternalMessageInfo

type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb0e6f64d0260de0, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryProvidersRequest)(nil), "sentinel.provider.v2.QueryProvidersRequest")
	proto.RegisterType((*QueryProviderRequest)(nil), "sentinel.provider.v2.QueryProviderRequest")
	proto.RegisterType((*QueryParamsRequest)(nil), "sentinel.provider.v2.QueryParamsRequest")
	proto.RegisterType((*QueryProvidersResponse)(nil), "sentinel.provider.v2.QueryProvidersResponse")
	proto.RegisterType((*QueryProviderResponse)(nil), "sentinel.provider.v2.QueryProviderResponse")
	proto.RegisterType((*QueryParamsResponse)(nil), "sentinel.provider.v2.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("sentinel/provider/v2/querier.proto", fileDescriptor_fb0e6f64d0260de0)
}

var fileDescriptor_fb0e6f64d0260de0 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x3c,
	0x1c, 0xc6, 0xe3, 0xed, 0xd5, 0x5e, 0xe6, 0x02, 0x07, 0xaf, 0x43, 0x53, 0x19, 0x1e, 0xcb, 0x04,
	0x74, 0x1b, 0xd8, 0x34, 0xdc, 0x38, 0xa1, 0x1e, 0xe0, 0xba, 0x75, 0x27, 0xb8, 0x20, 0xa7, 0xf5,
	0x32, 0x4b, 0x6d, 0x9c, 0xc5, 0x49, 0xc4, 0x84, 0xb8, 0x70, 0xe2, 0x82, 0x40, 0xe2, 0x33, 0x20,
	0xf1, 0x1d, 0xf8, 0x02, 0x3d, 0x4e, 0xe2, 0xc2, 0x09, 0x41, 0xcb, 0x07, 0x41, 0x75, 0xec, 0xa6,
	0x19, 0x01, 0x72, 0x4b, 0xd3, 0xe7, 0x79, 0xfc, 0x7b, 0xfc, 0xff, 0xb7, 0xd0, 0x55, 0x3c, 0x4c,
	0x44, 0xc8, 0x87, 0x34, 0x8a, 0x65, 0x26, 0x06, 0x3c, 0xa6, 0x99, 0x47, 0x4f, 0x53, 0x1e, 0x0b,
	0x1e, 0x93, 0x28, 0x96, 0x89, 0x44, 0x4d, 0xab, 0x21, 0x56, 0x43, 0x32, 0xaf, 0xb5, 0xd7, 0x97,
	0x6a, 0x24, 0x15, 0xf5, 0x99, 0xe2, 0xda, 0x70, 0x46, 0xb3, 0x8e, 0xcf, 0x13, 0xd6, 0xa1, 0x11,
	0x0b, 0x44, 0xc8, 0x12, 0x21, 0xc3, 0x3c, 0xa1, 0xd5, 0x0c, 0x64, 0x20, 0xf5, 0x23, 0x9d, 0x3d,
	0x99, 0xb7, 0x9b, 0x81, 0x94, 0xc1, 0x90, 0x53, 0x16, 0x09, 0xca, 0xc2, 0x50, 0x26, 0xda, 0xa2,
	0xcc, 0xb7, 0xdb, 0x95, 0x64, 0x11, 0x8b, 0xd9, 0xc8, 0x4a, 0x76, 0xaa, 0x25, 0x16, 0x52, 0x8b,
	0xdc, 0xe7, 0x70, 0xfd, 0x70, 0x46, 0x77, 0x60, 0x5e, 0xab, 0x1e, 0x3f, 0x4d, 0xb9, 0x4a, 0xd0,
	0x63, 0x08, 0x0b, 0xd0, 0x0d, 0x70, 0x13, 0xb4, 0x1b, 0xde, 0x6d, 0x92, 0xb7, 0x22, 0xb3, 0x56,
	0x44, 0xb7, 0x22, 0xa6, 0x15, 0x39, 0x60, 0x01, 0x37, 0xde, 0xde, 0x82, 0xd3, 0xbd, 0x0f, 0x9b,
	0xa5, 0x03, 0x6c, 0xfe, 0x06, 0xfc, 0x9f, 0x0d, 0x06, 0x31, 0x57, 0x4a, 0x87, 0xaf, 0xf6, 0xec,
	0x47, 0xb7, 0x09, 0x51, 0xee, 0xd0, 0x65, 0x8c, 0xde, 0xfd, 0x08, 0xe0, 0xb5, 0x8b, 0xa4, 0x2a,
	0x92, 0xa1, 0xe2, 0xa8, 0x0b, 0x57, 0x6d, 0xab, 0x59, 0xd8, 0x72, 0xbb, 0xe1, 0x61, 0x52, 0x35,
	0x15, 0x62, 0xbd, 0xdd, 0xff, 0xc6, 0xdf, 0xb6, 0x9c, 0x5e, 0x61, 0x43, 0x4f, 0x4a, 0x75, 0x97,
	0x74, 0xdd, 0x3b, 0xff, 0xac, 0x9b, 0x03, 0x94, 0xfa, 0x3e, 0xbd, 0x70, 0xa1, 0x73, 0xca, 0x47,
	0xf0, 0x92, 0x3d, 0xce, 0x5c, 0x67, 0x3d, 0xc8, 0xb9, 0xcb, 0x3d, 0x84, 0x6b, 0xa5, 0x8b, 0x31,
	0xc1, 0x0f, 0xe1, 0x4a, 0x3e, 0x77, 0x13, 0xbb, 0xf9, 0x87, 0x58, 0xad, 0x31, 0xa1, 0xc6, 0xe1,
	0x7d, 0x5e, 0x86, 0x97, 0x75, 0xe6, 0x11, 0x8f, 0x33, 0xd1, 0xe7, 0xe8, 0x0d, 0x80, 0x57, 0xcb,
	0xd7, 0x8c, 0xf6, 0xab, 0xf3, 0x2a, 0xd7, 0xa6, 0x75, 0xb7, 0x9e, 0x38, 0x47, 0x77, 0xaf, 0xbf,
	0xfe, 0xf2, 0xf3, 0xc3, 0xd2, 0x3a, 0x5a, 0xa3, 0xbf, 0xed, 0xaa, 0x42, 0xef, 0x00, 0xbc, 0x52,
	0xf2, 0xa1, 0xbd, 0x1a, 0xe1, 0x16, 0x64, 0xbf, 0x96, 0xd6, 0x70, 0xdc, 0xd2, 0x1c, 0x5b, 0xe8,
	0x46, 0x05, 0x07, 0x7d, 0x69, 0x16, 0xf3, 0x15, 0x7a, 0x0b, 0x60, 0x63, 0x61, 0x02, 0xa8, 0xfd,
	0xb7, 0x33, 0x16, 0xb7, 0xb7, 0xb5, 0x5b, 0x43, 0x69, 0x58, 0x76, 0x35, 0xcb, 0x0e, 0xda, 0x2e,
	0x58, 0x46, 0x72, 0x90, 0x0e, 0xb9, 0x2a, 0x7e, 0xc7, 0xf9, 0xf4, 0xba, 0x47, 0xe3, 0x1f, 0xd8,
	0xf9, 0x34, 0xc1, 0xce, 0x78, 0x82, 0xc1, 0xf9, 0x04, 0x83, 0xef, 0x13, 0x0c, 0xde, 0x4f, 0xb1,
	0x73, 0x3e, 0xc5, 0xce, 0xd7, 0x29, 0x76, 0x9e, 0x75, 0x02, 0x91, 0x9c, 0xa4, 0x3e, 0xe9, 0xcb,
	0xd1, 0x3c, 0xee, 0x9e, 0x3c, 0x3e, 0x16, 0x7d, 0xc1, 0x86, 0xf4, 0x24, 0xf5, 0xe9, 0x8b, 0x22,
	0x35, 0x39, 0x8b, 0xb8, 0xf2, 0x57, 0xf4, 0x1f, 0xc3, 0x83, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x1c, 0x6b, 0x66, 0xfc, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	QueryProviders(ctx context.Context, in *QueryProvidersRequest, opts ...grpc.CallOption) (*QueryProvidersResponse, error)
	QueryProvider(ctx context.Context, in *QueryProviderRequest, opts ...grpc.CallOption) (*QueryProviderResponse, error)
	QueryParams(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryProviders(ctx context.Context, in *QueryProvidersRequest, opts ...grpc.CallOption) (*QueryProvidersResponse, error) {
	out := new(QueryProvidersResponse)
	err := c.cc.Invoke(ctx, "/sentinel.provider.v2.QueryService/QueryProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryProvider(ctx context.Context, in *QueryProviderRequest, opts ...grpc.CallOption) (*QueryProviderResponse, error) {
	out := new(QueryProviderResponse)
	err := c.cc.Invoke(ctx, "/sentinel.provider.v2.QueryService/QueryProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryParams(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/sentinel.provider.v2.QueryService/QueryParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	QueryProviders(context.Context, *QueryProvidersRequest) (*QueryProvidersResponse, error)
	QueryProvider(context.Context, *QueryProviderRequest) (*QueryProviderResponse, error)
	QueryParams(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryProviders(ctx context.Context, req *QueryProvidersRequest) (*QueryProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProviders not implemented")
}
func (*UnimplementedQueryServiceServer) QueryProvider(ctx context.Context, req *QueryProviderRequest) (*QueryProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProvider not implemented")
}
func (*UnimplementedQueryServiceServer) QueryParams(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryParams not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.provider.v2.QueryService/QueryProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryProviders(ctx, req.(*QueryProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.provider.v2.QueryService/QueryProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryProvider(ctx, req.(*QueryProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.provider.v2.QueryService/QueryParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryParams(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.provider.v2.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryProviders",
			Handler:    _QueryService_QueryProviders_Handler,
		},
		{
			MethodName: "QueryProvider",
			Handler:    _QueryService_QueryProvider_Handler,
		},
		{
			MethodName: "QueryParams",
			Handler:    _QueryService_QueryParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/provider/v2/querier.proto",
}

func (m *QueryProvidersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProvidersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProvidersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProviderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProviderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProviderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuerier(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryProvidersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProvidersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProvidersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuerier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Providers) > 0 {
		for iNdEx := len(m.Providers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Providers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuerier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryProviderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProviderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProviderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Provider.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuerier(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuerier(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuerier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryProvidersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryProviderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryProvidersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Providers) > 0 {
		for _, e := range m.Providers {
			l = e.Size()
			n += 1 + l + sovQuerier(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuerier(uint64(l))
	}
	return n
}

func (m *QueryProviderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Provider.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuerier(uint64(l))
	return n
}

func sovQuerier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuerier(x uint64) (n int) {
	return sovQuerier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProvidersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryProvidersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProvidersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryProviderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryProviderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProviderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryProvidersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryProvidersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProvidersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Providers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Providers = append(m.Providers, Provider{})
			if err := m.Providers[len(m.Providers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryProviderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryProviderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProviderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Provider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuerier
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuerier
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
				return ErrInvalidLengthQuerier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuerier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuerier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuerier
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
func skipQuerier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
					return 0, ErrIntOverflowQuerier
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
				return 0, ErrInvalidLengthQuerier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuerier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuerier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuerier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuerier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuerier = fmt.Errorf("proto: unexpected end of group")
)
