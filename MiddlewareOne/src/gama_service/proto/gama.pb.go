// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gama.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	gama.proto

It has these top-level messages:
	ExecuteGamaRequest
	ExecuteGamaResponse
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

type ExecuteGamaRequest struct {
	Variable         *int32 `protobuf:"varint,1,req,name=variable" json:"variable,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ExecuteGamaRequest) Reset()                    { *m = ExecuteGamaRequest{} }
func (m *ExecuteGamaRequest) String() string            { return proto1.CompactTextString(m) }
func (*ExecuteGamaRequest) ProtoMessage()               {}
func (*ExecuteGamaRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExecuteGamaRequest) GetVariable() int32 {
	if m != nil && m.Variable != nil {
		return *m.Variable
	}
	return 0
}

type ExecuteGamaResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ExecuteGamaResponse) Reset()                    { *m = ExecuteGamaResponse{} }
func (m *ExecuteGamaResponse) String() string            { return proto1.CompactTextString(m) }
func (*ExecuteGamaResponse) ProtoMessage()               {}
func (*ExecuteGamaResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto1.RegisterType((*ExecuteGamaRequest)(nil), "proto.ExecuteGamaRequest")
	proto1.RegisterType((*ExecuteGamaResponse)(nil), "proto.ExecuteGamaResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GamaService service

type GamaServiceClient interface {
	ExecuteGama(ctx context.Context, in *ExecuteGamaRequest, opts ...grpc.CallOption) (*ExecuteGamaResponse, error)
}

type gamaServiceClient struct {
	cc *grpc.ClientConn
}

func NewGamaServiceClient(cc *grpc.ClientConn) GamaServiceClient {
	return &gamaServiceClient{cc}
}

func (c *gamaServiceClient) ExecuteGama(ctx context.Context, in *ExecuteGamaRequest, opts ...grpc.CallOption) (*ExecuteGamaResponse, error) {
	out := new(ExecuteGamaResponse)
	err := grpc.Invoke(ctx, "/proto.GamaService/executeGama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GamaService service

type GamaServiceServer interface {
	ExecuteGama(context.Context, *ExecuteGamaRequest) (*ExecuteGamaResponse, error)
}

func RegisterGamaServiceServer(s *grpc.Server, srv GamaServiceServer) {
	s.RegisterService(&_GamaService_serviceDesc, srv)
}

func _GamaService_ExecuteGama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteGamaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamaServiceServer).ExecuteGama(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GamaService/ExecuteGama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamaServiceServer).ExecuteGama(ctx, req.(*ExecuteGamaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GamaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GamaService",
	HandlerType: (*GamaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "executeGama",
			Handler:    _GamaService_ExecuteGama_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gama.proto",
}

func init() { proto1.RegisterFile("gama.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6a, 0x5c, 0x42, 0xae, 0x15,
	0xa9, 0xc9, 0xa5, 0x25, 0xa9, 0xee, 0x89, 0xb9, 0x89, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0x02, 0x5c, 0x1c, 0x65, 0x89, 0x45, 0x99, 0x89, 0x49, 0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x4c,
	0x1a, 0xac, 0x4a, 0xa2, 0x5c, 0xc2, 0x28, 0xea, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x8d, 0x82,
	0xb9, 0xb8, 0x41, 0xfc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x17, 0x2e, 0xee, 0x54,
	0x84, 0x2a, 0x21, 0x49, 0x88, 0x5d, 0x7a, 0x98, 0x36, 0x48, 0x49, 0x61, 0x93, 0x82, 0x18, 0x0a,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x84, 0x3b, 0xd9, 0xe2, 0xa7, 0x00, 0x00, 0x00,
}
