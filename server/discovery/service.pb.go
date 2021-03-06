// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Void
	DescribeResponse
*/
package discovery

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google/descriptor"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Void is an empty message.
type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// DescribeResponse describes services.
type DescribeResponse struct {
	// Description contains descriptions of all services, their types and all
	// transitive dependencies.
	Description *google_protobuf.FileDescriptorSet `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// Services are service names provided by a server.
	Services []string `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DescribeResponse) GetDescription() *google_protobuf.FileDescriptorSet {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*Void)(nil), "discovery.Void")
	proto.RegisterType((*DescribeResponse)(nil), "discovery.DescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Discovery service

type DiscoveryClient interface {
	// Describe returns a list of services and a descriptor.FileDescriptorSet
	// that covers them all.
	Describe(ctx context.Context, in *Void, opts ...grpc.CallOption) (*DescribeResponse, error)
}
type discoveryPRPCClient struct {
	client *prpccommon.Client
}

func NewDiscoveryPRPCClient(client *prpccommon.Client) DiscoveryClient {
	return &discoveryPRPCClient{client}
}

func (c *discoveryPRPCClient) Describe(ctx context.Context, in *Void, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := c.client.Call(ctx, "discovery.Discovery", "Describe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Describe(ctx context.Context, in *Void, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	// Describe returns a list of services and a descriptor.FileDescriptorSet
	// that covers them all.
	Describe(context.Context, *Void) (*DescribeResponse, error)
}

func RegisterDiscoveryServer(s prpc.Registrar, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Describe(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Discovery_Describe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0xcb, 0x82, 0x40,
	0x10, 0x86, 0x3f, 0xbf, 0x42, 0x74, 0x24, 0x8a, 0x3d, 0x89, 0x5d, 0x64, 0x4f, 0x9e, 0x56, 0xb0,
	0x5b, 0xe7, 0xa5, 0xee, 0x06, 0xfd, 0x00, 0x75, 0x92, 0x05, 0x71, 0x64, 0x77, 0x13, 0xfa, 0xf7,
	0x91, 0xb6, 0x22, 0x1d, 0x67, 0xde, 0x77, 0x1e, 0x9e, 0x81, 0x9d, 0x41, 0x3d, 0xaa, 0x1a, 0xc5,
	0xa0, 0xc9, 0x12, 0x0b, 0x1b, 0x65, 0x6a, 0x1a, 0x51, 0xbf, 0x92, 0xb4, 0x25, 0x6a, 0x3b, 0xcc,
	0xa7, 0xa0, 0x7a, 0x3e, 0xf2, 0x06, 0x4d, 0xad, 0xd5, 0x60, 0x49, 0xcf, 0x65, 0xee, 0xc3, 0xf6,
	0x4e, 0xaa, 0xe1, 0x16, 0x0e, 0x72, 0xca, 0x2a, 0x2c, 0xd1, 0x0c, 0xd4, 0x1b, 0x64, 0x12, 0x22,
	0xd7, 0x57, 0xd4, 0xc7, 0x5e, 0xea, 0x65, 0x51, 0xc1, 0xc5, 0xcc, 0x14, 0x8e, 0x29, 0x2e, 0xaa,
	0x43, 0xb9, 0x70, 0x6f, 0x68, 0xcb, 0xf5, 0x19, 0x4b, 0x20, 0xf8, 0xfa, 0x99, 0xf8, 0x3f, 0xdd,
	0x64, 0x61, 0xb9, 0xcc, 0xc5, 0x15, 0x42, 0xe9, 0x64, 0xd9, 0x19, 0x02, 0xa7, 0xc0, 0xf6, 0x62,
	0x79, 0x42, 0x7c, 0xfc, 0x92, 0xe3, 0x6a, 0xf1, 0x2b, 0xca, 0xff, 0x2a, 0x7f, 0xb2, 0x39, 0xbd,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x30, 0x21, 0xa9, 0x08, 0x0b, 0x01, 0x00, 0x00,
}
