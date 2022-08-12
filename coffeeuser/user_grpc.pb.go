// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package coffeeuser

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoffeeUserServicesClient is the client API for CoffeeUserServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoffeeUserServicesClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	AddCoffee(ctx context.Context, in *AddCoffeeToCart, opts ...grpc.CallOption) (*Response, error)
}

type coffeeUserServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewCoffeeUserServicesClient(cc grpc.ClientConnInterface) CoffeeUserServicesClient {
	return &coffeeUserServicesClient{cc}
}

func (c *coffeeUserServicesClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CoffeeUserServices/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeUserServicesClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CoffeeUserServices/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeUserServicesClient) AddCoffee(ctx context.Context, in *AddCoffeeToCart, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CoffeeUserServices/AddCoffee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoffeeUserServicesServer is the server API for CoffeeUserServices service.
// All implementations must embed UnimplementedCoffeeUserServicesServer
// for forward compatibility
type CoffeeUserServicesServer interface {
	Register(context.Context, *RegisterRequest) (*Response, error)
	Login(context.Context, *LoginRequest) (*Response, error)
	AddCoffee(context.Context, *AddCoffeeToCart) (*Response, error)
	mustEmbedUnimplementedCoffeeUserServicesServer()
}

// UnimplementedCoffeeUserServicesServer must be embedded to have forward compatible implementations.
type UnimplementedCoffeeUserServicesServer struct {
}

func (UnimplementedCoffeeUserServicesServer) Register(context.Context, *RegisterRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCoffeeUserServicesServer) Login(context.Context, *LoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCoffeeUserServicesServer) AddCoffee(context.Context, *AddCoffeeToCart) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoffee not implemented")
}
func (UnimplementedCoffeeUserServicesServer) mustEmbedUnimplementedCoffeeUserServicesServer() {}

// UnsafeCoffeeUserServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoffeeUserServicesServer will
// result in compilation errors.
type UnsafeCoffeeUserServicesServer interface {
	mustEmbedUnimplementedCoffeeUserServicesServer()
}

func RegisterCoffeeUserServicesServer(s grpc.ServiceRegistrar, srv CoffeeUserServicesServer) {
	s.RegisterService(&CoffeeUserServices_ServiceDesc, srv)
}

func _CoffeeUserServices_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeUserServicesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeUserServices/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeUserServicesServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeUserServices_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeUserServicesServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeUserServices/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeUserServicesServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeUserServices_AddCoffee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoffeeToCart)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeUserServicesServer).AddCoffee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeUserServices/AddCoffee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeUserServicesServer).AddCoffee(ctx, req.(*AddCoffeeToCart))
	}
	return interceptor(ctx, in, info, handler)
}

// CoffeeUserServices_ServiceDesc is the grpc.ServiceDesc for CoffeeUserServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoffeeUserServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoffeeUserServices",
	HandlerType: (*CoffeeUserServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _CoffeeUserServices_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CoffeeUserServices_Login_Handler,
		},
		{
			MethodName: "AddCoffee",
			Handler:    _CoffeeUserServices_AddCoffee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
