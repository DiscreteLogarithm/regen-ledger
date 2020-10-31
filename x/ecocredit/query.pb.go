// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1alpha1/query.proto

package ecocredit

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type QueryClassInfoRequest struct {
	ClassId string `protobuf:"bytes,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (m *QueryClassInfoRequest) Reset()         { *m = QueryClassInfoRequest{} }
func (m *QueryClassInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryClassInfoRequest) ProtoMessage()    {}
func (*QueryClassInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a16cc4c1db940dc, []int{0}
}
func (m *QueryClassInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClassInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClassInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClassInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClassInfoRequest.Merge(m, src)
}
func (m *QueryClassInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryClassInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClassInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClassInfoRequest proto.InternalMessageInfo

func (m *QueryClassInfoRequest) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

type QueryClassInfoResponse struct {
	Info *ClassInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *QueryClassInfoResponse) Reset()         { *m = QueryClassInfoResponse{} }
func (m *QueryClassInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClassInfoResponse) ProtoMessage()    {}
func (*QueryClassInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a16cc4c1db940dc, []int{1}
}
func (m *QueryClassInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClassInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClassInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClassInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClassInfoResponse.Merge(m, src)
}
func (m *QueryClassInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClassInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClassInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClassInfoResponse proto.InternalMessageInfo

func (m *QueryClassInfoResponse) GetInfo() *ClassInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type QueryBatchInfoRequest struct {
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
}

func (m *QueryBatchInfoRequest) Reset()         { *m = QueryBatchInfoRequest{} }
func (m *QueryBatchInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBatchInfoRequest) ProtoMessage()    {}
func (*QueryBatchInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a16cc4c1db940dc, []int{2}
}
func (m *QueryBatchInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBatchInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBatchInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBatchInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBatchInfoRequest.Merge(m, src)
}
func (m *QueryBatchInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBatchInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBatchInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBatchInfoRequest proto.InternalMessageInfo

func (m *QueryBatchInfoRequest) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

type QueryBatchInfoResponse struct {
	Info          *BatchInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	RetiredSupply string     `protobuf:"bytes,2,opt,name=retired_supply,json=retiredSupply,proto3" json:"retired_supply,omitempty"`
}

func (m *QueryBatchInfoResponse) Reset()         { *m = QueryBatchInfoResponse{} }
func (m *QueryBatchInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBatchInfoResponse) ProtoMessage()    {}
func (*QueryBatchInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a16cc4c1db940dc, []int{3}
}
func (m *QueryBatchInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBatchInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBatchInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBatchInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBatchInfoResponse.Merge(m, src)
}
func (m *QueryBatchInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBatchInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBatchInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBatchInfoResponse proto.InternalMessageInfo

func (m *QueryBatchInfoResponse) GetInfo() *BatchInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *QueryBatchInfoResponse) GetRetiredSupply() string {
	if m != nil {
		return m.RetiredSupply
	}
	return ""
}

type QueryBalanceRequest struct {
	Account    string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
}

func (m *QueryBalanceRequest) Reset()         { *m = QueryBalanceRequest{} }
func (m *QueryBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceRequest) ProtoMessage()    {}
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a16cc4c1db940dc, []int{4}
}
func (m *QueryBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceRequest.Merge(m, src)
}
func (m *QueryBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceRequest proto.InternalMessageInfo

func (m *QueryBalanceRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *QueryBalanceRequest) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

type QueryBalanceResponse struct {
	TradeableUnits string `protobuf:"bytes,1,opt,name=tradeable_units,json=tradeableUnits,proto3" json:"tradeable_units,omitempty"`
	RetiredUnits   string `protobuf:"bytes,2,opt,name=retired_units,json=retiredUnits,proto3" json:"retired_units,omitempty"`
}

func (m *QueryBalanceResponse) Reset()         { *m = QueryBalanceResponse{} }
func (m *QueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceResponse) ProtoMessage()    {}
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a16cc4c1db940dc, []int{5}
}
func (m *QueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceResponse.Merge(m, src)
}
func (m *QueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceResponse proto.InternalMessageInfo

func (m *QueryBalanceResponse) GetTradeableUnits() string {
	if m != nil {
		return m.TradeableUnits
	}
	return ""
}

func (m *QueryBalanceResponse) GetRetiredUnits() string {
	if m != nil {
		return m.RetiredUnits
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryClassInfoRequest)(nil), "regen.data.v1alpha1.QueryClassInfoRequest")
	proto.RegisterType((*QueryClassInfoResponse)(nil), "regen.data.v1alpha1.QueryClassInfoResponse")
	proto.RegisterType((*QueryBatchInfoRequest)(nil), "regen.data.v1alpha1.QueryBatchInfoRequest")
	proto.RegisterType((*QueryBatchInfoResponse)(nil), "regen.data.v1alpha1.QueryBatchInfoResponse")
	proto.RegisterType((*QueryBalanceRequest)(nil), "regen.data.v1alpha1.QueryBalanceRequest")
	proto.RegisterType((*QueryBalanceResponse)(nil), "regen.data.v1alpha1.QueryBalanceResponse")
}

func init() {
	proto.RegisterFile("regen/ecocredit/v1alpha1/query.proto", fileDescriptor_6a16cc4c1db940dc)
}

var fileDescriptor_6a16cc4c1db940dc = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x6d, 0x2a, 0xa0, 0x74, 0x16, 0x16, 0xc9, 0x0b, 0xa8, 0xf4, 0x10, 0x50, 0x00, 0xb1, 0x80,
	0x48, 0xd4, 0x70, 0xe1, 0xbc, 0x70, 0x59, 0xc1, 0x01, 0x16, 0x71, 0xe1, 0x40, 0xe4, 0xd8, 0xd3,
	0x36, 0x22, 0xb5, 0x53, 0xdb, 0x01, 0xfa, 0x17, 0xfc, 0x0b, 0x3f, 0xc1, 0xb1, 0x47, 0x8e, 0xa8,
	0xfd, 0x11, 0x14, 0xd7, 0x69, 0x4b, 0x5b, 0xaa, 0x72, 0x9c, 0xa7, 0xf7, 0xe6, 0xbd, 0x99, 0xb1,
	0xe1, 0x81, 0xc2, 0x01, 0x8a, 0x08, 0x99, 0x64, 0x0a, 0x79, 0x66, 0xa2, 0x2f, 0x3d, 0x9a, 0x17,
	0x43, 0xda, 0x8b, 0xc6, 0x25, 0xaa, 0x49, 0x58, 0x28, 0x69, 0x24, 0x39, 0xb1, 0xac, 0x90, 0x53,
	0x43, 0xc3, 0x9a, 0xd0, 0xfd, 0xb7, 0xd4, 0x4c, 0x0a, 0xd4, 0x0b, 0x69, 0x10, 0xc3, 0xad, 0x77,
	0x55, 0xa7, 0x97, 0x39, 0xd5, 0xfa, 0x5c, 0xf4, 0xe5, 0x05, 0x8e, 0x4b, 0xd4, 0x86, 0xdc, 0x81,
	0xab, 0xac, 0xc2, 0x92, 0x8c, 0x77, 0xbc, 0x7b, 0xde, 0x69, 0xfb, 0xa2, 0x65, 0xeb, 0x73, 0x1e,
	0xbc, 0x81, 0xdb, 0x9b, 0x1a, 0x5d, 0x48, 0xa1, 0x91, 0xc4, 0x70, 0x29, 0x13, 0x7d, 0x69, 0x05,
	0x47, 0xb1, 0x1f, 0xee, 0xc8, 0x15, 0xae, 0x54, 0x96, 0x1b, 0xbc, 0x70, 0x09, 0xce, 0xa8, 0x61,
	0xc3, 0xf5, 0x04, 0x77, 0xe1, 0x28, 0xad, 0xb0, 0x84, 0xa3, 0x90, 0x23, 0x17, 0x02, 0x2c, 0xf4,
	0xaa, 0x42, 0x02, 0xed, 0x72, 0xac, 0x29, 0xff, 0x23, 0xc7, 0x4a, 0x65, 0xb9, 0xe4, 0x21, 0x1c,
	0x2b, 0x34, 0x99, 0x42, 0x9e, 0xe8, 0xb2, 0x28, 0xf2, 0x49, 0xa7, 0x69, 0x1d, 0xaf, 0x3b, 0xf4,
	0xbd, 0x05, 0x83, 0xb7, 0x70, 0xe2, 0x4c, 0x73, 0x2a, 0x18, 0xd6, 0x61, 0x3b, 0xd0, 0xa2, 0x8c,
	0xc9, 0x52, 0x98, 0x7a, 0x5b, 0xae, 0xdc, 0x1c, 0xa3, 0xb9, 0x35, 0x06, 0x87, 0x9b, 0x7f, 0x77,
	0x74, 0x43, 0x3c, 0x82, 0x1b, 0x46, 0x51, 0x8e, 0x34, 0xcd, 0x31, 0x29, 0x45, 0x66, 0xb4, 0x6b,
	0x7d, 0xbc, 0x84, 0x3f, 0x54, 0x28, 0xb9, 0x0f, 0x75, 0x46, 0x47, 0x5b, 0x78, 0x5c, 0x73, 0xa0,
	0x25, 0xc5, 0x3f, 0x9a, 0x70, 0xd9, 0xda, 0x10, 0x0e, 0xed, 0xe5, 0x0d, 0xc8, 0x93, 0x9d, 0xbb,
	0xd9, 0xf9, 0x24, 0xba, 0x4f, 0x0f, 0xe2, 0xba, 0xf4, 0x1c, 0xda, 0xcb, 0x0d, 0xef, 0x73, 0xd9,
	0x3c, 0xfb, 0x3e, 0x97, 0xed, 0x43, 0x7f, 0x82, 0x96, 0x5b, 0x1b, 0x39, 0xdd, 0xa7, 0x5b, 0xbf,
	0x55, 0xf7, 0xf1, 0x01, 0xcc, 0x45, 0xff, 0xb3, 0xd7, 0x3f, 0x67, 0xbe, 0x37, 0x9d, 0xf9, 0xde,
	0xef, 0x99, 0xef, 0x7d, 0x9f, 0xfb, 0x8d, 0xe9, 0xdc, 0x6f, 0xfc, 0x9a, 0xfb, 0x8d, 0x8f, 0xbd,
	0x41, 0x66, 0x86, 0x65, 0x1a, 0x32, 0x39, 0x8a, 0x6c, 0xbb, 0x67, 0x02, 0xcd, 0x57, 0xa9, 0x3e,
	0xbb, 0x2a, 0x47, 0x3e, 0x40, 0x15, 0x7d, 0x5b, 0x7d, 0xc0, 0xf4, 0x8a, 0xfd, 0x72, 0xcf, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xf1, 0x36, 0x1f, 0xd5, 0x03, 0x00, 0x00,
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
	ClassInfo(ctx context.Context, in *QueryClassInfoRequest, opts ...grpc.CallOption) (*QueryClassInfoResponse, error)
	BatchInfo(ctx context.Context, in *QueryBatchInfoRequest, opts ...grpc.CallOption) (*QueryBatchInfoResponse, error)
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ClassInfo(ctx context.Context, in *QueryClassInfoRequest, opts ...grpc.CallOption) (*QueryClassInfoResponse, error) {
	out := new(QueryClassInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Query/ClassInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BatchInfo(ctx context.Context, in *QueryBatchInfoRequest, opts ...grpc.CallOption) (*QueryBatchInfoResponse, error) {
	out := new(QueryBatchInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Query/BatchInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	ClassInfo(context.Context, *QueryClassInfoRequest) (*QueryClassInfoResponse, error)
	BatchInfo(context.Context, *QueryBatchInfoRequest) (*QueryBatchInfoResponse, error)
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ClassInfo(ctx context.Context, req *QueryClassInfoRequest) (*QueryClassInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassInfo not implemented")
}
func (*UnimplementedQueryServer) BatchInfo(ctx context.Context, req *QueryBatchInfoRequest) (*QueryBatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchInfo not implemented")
}
func (*UnimplementedQueryServer) Balance(ctx context.Context, req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ClassInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClassInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Query/ClassInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClassInfo(ctx, req.(*QueryClassInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Query/BatchInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BatchInfo(ctx, req.(*QueryBatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "regen.data.v1alpha1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClassInfo",
			Handler:    _Query_ClassInfo_Handler,
		},
		{
			MethodName: "BatchInfo",
			Handler:    _Query_BatchInfo_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/v1alpha1/query.proto",
}

func (m *QueryClassInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClassInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClassInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryClassInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClassInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClassInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryBatchInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBatchInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBatchInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBatchInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBatchInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBatchInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RetiredSupply) > 0 {
		i -= len(m.RetiredSupply)
		copy(dAtA[i:], m.RetiredSupply)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RetiredSupply)))
		i--
		dAtA[i] = 0x12
	}
	if m.Info != nil {
		{
			size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RetiredUnits) > 0 {
		i -= len(m.RetiredUnits)
		copy(dAtA[i:], m.RetiredUnits)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RetiredUnits)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TradeableUnits) > 0 {
		i -= len(m.TradeableUnits)
		copy(dAtA[i:], m.TradeableUnits)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TradeableUnits)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryClassInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryClassInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBatchInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBatchInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.RetiredSupply)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TradeableUnits)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.RetiredUnits)
	if l > 0 {
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
func (m *QueryClassInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClassInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClassInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
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
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryClassInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClassInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClassInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
			if m.Info == nil {
				m.Info = &ClassInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryBatchInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBatchInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBatchInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
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
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryBatchInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBatchInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBatchInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
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
			if m.Info == nil {
				m.Info = &BatchInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetiredSupply", wireType)
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
			m.RetiredSupply = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryBalanceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
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
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryBalanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeableUnits", wireType)
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
			m.TradeableUnits = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetiredUnits", wireType)
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
			m.RetiredUnits = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
