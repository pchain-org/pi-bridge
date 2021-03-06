// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proxy/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// this line is used by starport scaffolding # 3
type QueryGetProxyRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetProxyRequest) Reset()         { *m = QueryGetProxyRequest{} }
func (m *QueryGetProxyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetProxyRequest) ProtoMessage()    {}
func (*QueryGetProxyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5027a8bbaf1cb4eb, []int{0}
}
func (m *QueryGetProxyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProxyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProxyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProxyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProxyRequest.Merge(m, src)
}
func (m *QueryGetProxyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProxyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProxyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProxyRequest proto.InternalMessageInfo

func (m *QueryGetProxyRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetProxyResponse struct {
	Proxy *Proxy `protobuf:"bytes,1,opt,name=Proxy,proto3" json:"Proxy,omitempty"`
}

func (m *QueryGetProxyResponse) Reset()         { *m = QueryGetProxyResponse{} }
func (m *QueryGetProxyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetProxyResponse) ProtoMessage()    {}
func (*QueryGetProxyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5027a8bbaf1cb4eb, []int{1}
}
func (m *QueryGetProxyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetProxyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetProxyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetProxyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetProxyResponse.Merge(m, src)
}
func (m *QueryGetProxyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetProxyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetProxyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetProxyResponse proto.InternalMessageInfo

func (m *QueryGetProxyResponse) GetProxy() *Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

type QueryAllProxyRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllProxyRequest) Reset()         { *m = QueryAllProxyRequest{} }
func (m *QueryAllProxyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllProxyRequest) ProtoMessage()    {}
func (*QueryAllProxyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5027a8bbaf1cb4eb, []int{2}
}
func (m *QueryAllProxyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllProxyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllProxyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllProxyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllProxyRequest.Merge(m, src)
}
func (m *QueryAllProxyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllProxyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllProxyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllProxyRequest proto.InternalMessageInfo

func (m *QueryAllProxyRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllProxyResponse struct {
	Proxy      []*Proxy            `protobuf:"bytes,1,rep,name=Proxy,proto3" json:"Proxy,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllProxyResponse) Reset()         { *m = QueryAllProxyResponse{} }
func (m *QueryAllProxyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllProxyResponse) ProtoMessage()    {}
func (*QueryAllProxyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5027a8bbaf1cb4eb, []int{3}
}
func (m *QueryAllProxyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllProxyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllProxyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllProxyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllProxyResponse.Merge(m, src)
}
func (m *QueryAllProxyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllProxyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllProxyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllProxyResponse proto.InternalMessageInfo

func (m *QueryAllProxyResponse) GetProxy() []*Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *QueryAllProxyResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetProxyRequest)(nil), "pchainorg.pibridge.proxy.QueryGetProxyRequest")
	proto.RegisterType((*QueryGetProxyResponse)(nil), "pchainorg.pibridge.proxy.QueryGetProxyResponse")
	proto.RegisterType((*QueryAllProxyRequest)(nil), "pchainorg.pibridge.proxy.QueryAllProxyRequest")
	proto.RegisterType((*QueryAllProxyResponse)(nil), "pchainorg.pibridge.proxy.QueryAllProxyResponse")
}

func init() { proto.RegisterFile("proxy/query.proto", fileDescriptor_5027a8bbaf1cb4eb) }

var fileDescriptor_5027a8bbaf1cb4eb = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6a, 0x1a, 0x41,
	0x1c, 0xc7, 0x1d, 0x8b, 0xa5, 0x9d, 0x9e, 0x3a, 0x58, 0x10, 0x29, 0x5b, 0xd9, 0x43, 0x2b, 0x52,
	0x67, 0xaa, 0xa5, 0x0f, 0x60, 0x0f, 0xf1, 0x16, 0x8c, 0xc7, 0x1c, 0x02, 0xb3, 0x3a, 0x8c, 0x03,
	0xeb, 0xce, 0xb8, 0x33, 0x06, 0x25, 0xe4, 0x92, 0x27, 0x08, 0x84, 0x10, 0xc8, 0xb3, 0xe4, 0x01,
	0x72, 0x14, 0x72, 0xc9, 0x31, 0x68, 0x1e, 0x24, 0x38, 0x33, 0xc1, 0x3f, 0x89, 0xac, 0x17, 0xc1,
	0x1f, 0xdf, 0x3f, 0x9f, 0xdf, 0xfe, 0x76, 0xe1, 0x57, 0x95, 0xca, 0xc9, 0x94, 0x8c, 0xc6, 0x2c,
	0x9d, 0x62, 0x95, 0x4a, 0x23, 0x51, 0x49, 0xf5, 0x06, 0x54, 0x24, 0x32, 0xe5, 0x58, 0x89, 0x28,
	0x15, 0x7d, 0xce, 0xb0, 0x55, 0x95, 0xbf, 0x73, 0x29, 0x79, 0xcc, 0x08, 0x55, 0x82, 0xd0, 0x24,
	0x91, 0x86, 0x1a, 0x21, 0x13, 0xed, 0x7c, 0xe5, 0x5a, 0x4f, 0xea, 0xa1, 0xd4, 0x24, 0xa2, 0x9a,
	0xb9, 0x40, 0x72, 0xda, 0x88, 0x98, 0xa1, 0x0d, 0xa2, 0x28, 0x17, 0x89, 0x15, 0x7b, 0xad, 0xaf,
	0xb5, 0xbf, 0x6e, 0x14, 0xfe, 0x86, 0xc5, 0xa3, 0xa5, 0xa9, 0xcd, 0x4c, 0x67, 0x39, 0xee, 0xb2,
	0xd1, 0x98, 0x69, 0x83, 0x8a, 0xb0, 0x20, 0x92, 0x3e, 0x9b, 0x94, 0x40, 0x05, 0x54, 0x3f, 0x77,
	0xdd, 0x9f, 0xf0, 0x10, 0x7e, 0xdb, 0x52, 0x6b, 0x25, 0x13, 0xcd, 0xd0, 0x3f, 0x58, 0xb0, 0x03,
	0x2b, 0xff, 0xd2, 0xfc, 0x81, 0x77, 0x6d, 0x83, 0x9d, 0xcf, 0xa9, 0xc3, 0x13, 0xdf, 0xde, 0x8a,
	0xe3, 0x8d, 0xf6, 0x03, 0x08, 0x57, 0xf0, 0x3e, 0xf3, 0x27, 0x76, 0x9b, 0xe2, 0xe5, 0xa6, 0xd8,
	0x3d, 0x3a, 0xbf, 0x29, 0xee, 0x50, 0xce, 0xbc, 0xb7, 0xbb, 0xe6, 0x0c, 0x6f, 0x80, 0x07, 0x5e,
	0x15, 0xbc, 0x05, 0xfe, 0xb0, 0x3f, 0x30, 0x6a, 0x6f, 0x80, 0xe5, 0x2d, 0xd8, 0xaf, 0x4c, 0x30,
	0xd7, 0xb9, 0x4e, 0xd6, 0xbc, 0xcb, 0xc3, 0x82, 0x25, 0x43, 0xb7, 0xc0, 0xa3, 0x20, 0xbc, 0x1b,
	0xe2, 0xbd, 0x1b, 0x95, 0xc9, 0xde, 0x7a, 0x07, 0x10, 0xfe, 0xb9, 0x78, 0x78, 0xbe, 0xca, 0xd7,
	0x50, 0x95, 0x38, 0x63, 0x5d, 0xa6, 0x9c, 0xbc, 0x3a, 0xc9, 0xda, 0xcb, 0x41, 0xce, 0xec, 0xbd,
	0xcf, 0xd1, 0x35, 0x80, 0x9f, 0x6c, 0x46, 0x2b, 0x8e, 0x33, 0xf9, 0xb6, 0xae, 0x98, 0xc9, 0xb7,
	0x7d, 0x94, 0xb0, 0x6a, 0xf9, 0x42, 0x54, 0xc9, 0xe2, 0xfb, 0xdf, 0xbe, 0x9f, 0x07, 0x60, 0x36,
	0x0f, 0xc0, 0xd3, 0x3c, 0x00, 0x97, 0x8b, 0x20, 0x37, 0x5b, 0x04, 0xb9, 0xc7, 0x45, 0x90, 0x3b,
	0xae, 0x73, 0x61, 0x06, 0xe3, 0x08, 0xf7, 0xe4, 0x70, 0x33, 0xa5, 0xee, 0x63, 0x26, 0x3e, 0xc8,
	0x4c, 0x15, 0xd3, 0xd1, 0x47, 0xfb, 0x19, 0xfc, 0x7d, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x79,
	0x87, 0x26, 0x92, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a proxy by index.
	Proxy(ctx context.Context, in *QueryGetProxyRequest, opts ...grpc.CallOption) (*QueryGetProxyResponse, error)
	// Queries a list of proxy items.
	ProxyAll(ctx context.Context, in *QueryAllProxyRequest, opts ...grpc.CallOption) (*QueryAllProxyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Proxy(ctx context.Context, in *QueryGetProxyRequest, opts ...grpc.CallOption) (*QueryGetProxyResponse, error) {
	out := new(QueryGetProxyResponse)
	err := c.cc.Invoke(ctx, "/pchainorg.pibridge.proxy.Query/Proxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProxyAll(ctx context.Context, in *QueryAllProxyRequest, opts ...grpc.CallOption) (*QueryAllProxyResponse, error) {
	out := new(QueryAllProxyResponse)
	err := c.cc.Invoke(ctx, "/pchainorg.pibridge.proxy.Query/ProxyAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a proxy by index.
	Proxy(context.Context, *QueryGetProxyRequest) (*QueryGetProxyResponse, error)
	// Queries a list of proxy items.
	ProxyAll(context.Context, *QueryAllProxyRequest) (*QueryAllProxyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Proxy(ctx context.Context, req *QueryGetProxyRequest) (*QueryGetProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proxy not implemented")
}
func (*UnimplementedQueryServer) ProxyAll(ctx context.Context, req *QueryAllProxyRequest) (*QueryAllProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Proxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Proxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pchainorg.pibridge.proxy.Query/Proxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Proxy(ctx, req.(*QueryGetProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProxyAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProxyAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pchainorg.pibridge.proxy.Query/ProxyAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProxyAll(ctx, req.(*QueryAllProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pchainorg.pibridge.proxy.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proxy",
			Handler:    _Query_Proxy_Handler,
		},
		{
			MethodName: "ProxyAll",
			Handler:    _Query_ProxyAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proxy/query.proto",
}

func (m *QueryGetProxyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProxyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProxyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetProxyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetProxyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetProxyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proxy != nil {
		{
			size, err := m.Proxy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllProxyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllProxyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllProxyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllProxyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllProxyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllProxyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Proxy) > 0 {
		for iNdEx := len(m.Proxy) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proxy[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetProxyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetProxyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Proxy != nil {
		l = m.Proxy.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllProxyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllProxyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Proxy) > 0 {
		for _, e := range m.Proxy {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetProxyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetProxyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProxyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetProxyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetProxyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetProxyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proxy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proxy == nil {
				m.Proxy = &Proxy{}
			}
			if err := m.Proxy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllProxyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllProxyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllProxyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllProxyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllProxyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllProxyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proxy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proxy = append(m.Proxy, &Proxy{})
			if err := m.Proxy[len(m.Proxy)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
