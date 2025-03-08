// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: imageconverter/image_converter_service.proto

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	model "shared/model"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Converter_ConvertToWebP_FullMethodName = "/Converter/ConvertToWebP"
)

// ConverterClient is the client API for Converter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Image Converter GRPC Operations
type ConverterClient interface {
	// Convert Image to WebP format
	ConvertToWebP(ctx context.Context, in *model.ImageRequest, opts ...grpc.CallOption) (*model.ImageRequest, error)
}

type converterClient struct {
	cc grpc.ClientConnInterface
}

func NewConverterClient(cc grpc.ClientConnInterface) ConverterClient {
	return &converterClient{cc}
}

func (c *converterClient) ConvertToWebP(ctx context.Context, in *model.ImageRequest, opts ...grpc.CallOption) (*model.ImageRequest, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(model.ImageRequest)
	err := c.cc.Invoke(ctx, Converter_ConvertToWebP_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConverterServer is the server API for Converter service.
// All implementations must embed UnimplementedConverterServer
// for forward compatibility.
//
// Image Converter GRPC Operations
type ConverterServer interface {
	// Convert Image to WebP format
	ConvertToWebP(context.Context, *model.ImageRequest) (*model.ImageRequest, error)
	mustEmbedUnimplementedConverterServer()
}

// UnimplementedConverterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConverterServer struct{}

func (UnimplementedConverterServer) ConvertToWebP(context.Context, *model.ImageRequest) (*model.ImageRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertToWebP not implemented")
}
func (UnimplementedConverterServer) mustEmbedUnimplementedConverterServer() {}
func (UnimplementedConverterServer) testEmbeddedByValue()                   {}

// UnsafeConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConverterServer will
// result in compilation errors.
type UnsafeConverterServer interface {
	mustEmbedUnimplementedConverterServer()
}

func RegisterConverterServer(s grpc.ServiceRegistrar, srv ConverterServer) {
	// If the following call pancis, it indicates UnimplementedConverterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Converter_ServiceDesc, srv)
}

func _Converter_ConvertToWebP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).ConvertToWebP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Converter_ConvertToWebP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).ConvertToWebP(ctx, req.(*model.ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Converter_ServiceDesc is the grpc.ServiceDesc for Converter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Converter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Converter",
	HandlerType: (*ConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertToWebP",
			Handler:    _Converter_ConvertToWebP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imageconverter/image_converter_service.proto",
}
