// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	QueryRequest
	QueryResponse
	ExecRequest
	ExecResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// The query statement
type QueryRequest struct {
	Stmt string `protobuf:"bytes,1,opt,name=stmt" json:"stmt,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto1.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryRequest) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

// The query response
type QueryResponse struct {
	Columns []string `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
	Types   []string `protobuf:"bytes,2,rep,name=types" json:"types,omitempty"`
	Values  [][]byte `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto1.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *QueryResponse) GetTypes() []string {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *QueryResponse) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

// The exec statement
type ExecRequest struct {
	Stmt string `protobuf:"bytes,1,opt,name=stmt" json:"stmt,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto1.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExecRequest) GetStmt() string {
	if m != nil {
		return m.Stmt
	}
	return ""
}

// The exec response
type ExecResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto1.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExecResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*QueryRequest)(nil), "proto.QueryRequest")
	proto1.RegisterType((*QueryResponse)(nil), "proto.QueryResponse")
	proto1.RegisterType((*ExecRequest)(nil), "proto.ExecRequest")
	proto1.RegisterType((*ExecResponse)(nil), "proto.ExecResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DBProvider service

type DBProviderClient interface {
	// Query executes a statement that reads data.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	// Exec executes a statement that writes data.
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type dBProviderClient struct {
	cc *grpc.ClientConn
}

func NewDBProviderClient(cc *grpc.ClientConn) DBProviderClient {
	return &dBProviderClient{cc}
}

func (c *dBProviderClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := grpc.Invoke(ctx, "/proto.DBProvider/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBProviderClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := grpc.Invoke(ctx, "/proto.DBProvider/Exec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DBProvider service

type DBProviderServer interface {
	// Query executes a statement that reads data.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	// Exec executes a statement that writes data.
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
}

func RegisterDBProviderServer(s *grpc.Server, srv DBProviderServer) {
	s.RegisterService(&_DBProvider_serviceDesc, srv)
}

func _DBProvider_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBProviderServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBProvider/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBProviderServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBProvider_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBProviderServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DBProvider/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBProviderServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DBProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DBProvider",
	HandlerType: (*DBProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DBProvider_Query_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _DBProvider_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x09, 0x69, 0x8a, 0x7a, 0xa4, 0xcb, 0xb5, 0x42, 0x56, 0xa7, 0xe0, 0xc9, 0x53, 0x25,
	0xfe, 0x7c, 0x02, 0x04, 0x3b, 0x78, 0x61, 0x2e, 0xe1, 0x84, 0x2a, 0x35, 0x71, 0xf0, 0xd9, 0x11,
	0xf9, 0xf6, 0x08, 0xdb, 0x11, 0xc9, 0xd2, 0xc9, 0x7e, 0x3f, 0x3f, 0xdf, 0xbb, 0x07, 0x6b, 0x26,
	0xdb, 0x1f, 0x6b, 0xda, 0x77, 0xd6, 0x38, 0x83, 0x45, 0x38, 0xa4, 0x84, 0xf2, 0xcd, 0x93, 0x1d,
	0x34, 0x7d, 0x7b, 0x62, 0x87, 0x08, 0x0b, 0x76, 0x8d, 0x13, 0x59, 0x95, 0xa9, 0x95, 0x0e, 0x77,
	0xf9, 0x0e, 0xeb, 0xe4, 0xe1, 0xce, 0xb4, 0x4c, 0x28, 0xe0, 0xaa, 0x36, 0x27, 0xdf, 0xb4, 0x2c,
	0xb2, 0x2a, 0x57, 0x2b, 0x3d, 0x4a, 0xdc, 0x42, 0xe1, 0x86, 0x8e, 0x58, 0x5c, 0x06, 0x1e, 0x05,
	0xde, 0xc0, 0xb2, 0x3f, 0x9c, 0x3c, 0xb1, 0xc8, 0xab, 0x5c, 0x95, 0x3a, 0x29, 0x79, 0x0b, 0xd7,
	0x2f, 0x3f, 0x54, 0x9f, 0xcb, 0x56, 0x50, 0x46, 0xcb, 0x7f, 0x74, 0x43, 0xcc, 0x87, 0x2f, 0x4a,
	0xb6, 0x51, 0xde, 0x7b, 0x80, 0xe7, 0xa7, 0x57, 0x6b, 0xfa, 0xe3, 0x27, 0x59, 0x7c, 0x84, 0x22,
	0xec, 0x8c, 0x9b, 0xd8, 0x77, 0x3f, 0x6d, 0xb9, 0xdb, 0xce, 0x61, 0x9c, 0x2d, 0x2f, 0xf0, 0x0e,
	0x16, 0x7f, 0x69, 0x88, 0xe9, 0x7d, 0xb2, 0xdd, 0x6e, 0x33, 0x63, 0xe3, 0x97, 0x8f, 0x65, 0xa0,
	0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0x0b, 0xb5, 0x39, 0x5f, 0x01, 0x00, 0x00,
}
