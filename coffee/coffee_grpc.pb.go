// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package coffee

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

// CoffeeServiceClient is the client API for CoffeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoffeeServiceClient interface {
	AddCoffee(ctx context.Context, in *AddCoffeeRequest, opts ...grpc.CallOption) (*AddCoffeeResponse, error)
	FindCoffeeByID(ctx context.Context, in *FindCoffeeById, opts ...grpc.CallOption) (*AddCoffeeResponse, error)
	UpdateCoffeeInfo(ctx context.Context, in *UpdateCoffee, opts ...grpc.CallOption) (*AddCoffeeResponse, error)
	DeleteCoffeeInfo(ctx context.Context, in *DeleteCoffeeOrder, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
	FindAllCoffee(ctx context.Context, in *GetAllCoffee, opts ...grpc.CallOption) (*AddCoffeeResponse, error)
}

type coffeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoffeeServiceClient(cc grpc.ClientConnInterface) CoffeeServiceClient {
	return &coffeeServiceClient{cc}
}

func (c *coffeeServiceClient) AddCoffee(ctx context.Context, in *AddCoffeeRequest, opts ...grpc.CallOption) (*AddCoffeeResponse, error) {
	out := new(AddCoffeeResponse)
	err := c.cc.Invoke(ctx, "/CoffeeService/AddCoffee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeServiceClient) FindCoffeeByID(ctx context.Context, in *FindCoffeeById, opts ...grpc.CallOption) (*AddCoffeeResponse, error) {
	out := new(AddCoffeeResponse)
	err := c.cc.Invoke(ctx, "/CoffeeService/FindCoffeeByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeServiceClient) UpdateCoffeeInfo(ctx context.Context, in *UpdateCoffee, opts ...grpc.CallOption) (*AddCoffeeResponse, error) {
	out := new(AddCoffeeResponse)
	err := c.cc.Invoke(ctx, "/CoffeeService/UpdateCoffeeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeServiceClient) DeleteCoffeeInfo(ctx context.Context, in *DeleteCoffeeOrder, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, "/CoffeeService/DeleteCoffeeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coffeeServiceClient) FindAllCoffee(ctx context.Context, in *GetAllCoffee, opts ...grpc.CallOption) (*AddCoffeeResponse, error) {
	out := new(AddCoffeeResponse)
	err := c.cc.Invoke(ctx, "/CoffeeService/FindAllCoffee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoffeeServiceServer is the server API for CoffeeService service.
// All implementations must embed UnimplementedCoffeeServiceServer
// for forward compatibility
type CoffeeServiceServer interface {
	AddCoffee(context.Context, *AddCoffeeRequest) (*AddCoffeeResponse, error)
	FindCoffeeByID(context.Context, *FindCoffeeById) (*AddCoffeeResponse, error)
	UpdateCoffeeInfo(context.Context, *UpdateCoffee) (*AddCoffeeResponse, error)
	DeleteCoffeeInfo(context.Context, *DeleteCoffeeOrder) (*DeleteOrderResponse, error)
	FindAllCoffee(context.Context, *GetAllCoffee) (*AddCoffeeResponse, error)
	mustEmbedUnimplementedCoffeeServiceServer()
}

// UnimplementedCoffeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoffeeServiceServer struct {
}

func (UnimplementedCoffeeServiceServer) AddCoffee(context.Context, *AddCoffeeRequest) (*AddCoffeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCoffee not implemented")
}
func (UnimplementedCoffeeServiceServer) FindCoffeeByID(context.Context, *FindCoffeeById) (*AddCoffeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindCoffeeByID not implemented")
}
func (UnimplementedCoffeeServiceServer) UpdateCoffeeInfo(context.Context, *UpdateCoffee) (*AddCoffeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoffeeInfo not implemented")
}
func (UnimplementedCoffeeServiceServer) DeleteCoffeeInfo(context.Context, *DeleteCoffeeOrder) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoffeeInfo not implemented")
}
func (UnimplementedCoffeeServiceServer) FindAllCoffee(context.Context, *GetAllCoffee) (*AddCoffeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllCoffee not implemented")
}
func (UnimplementedCoffeeServiceServer) mustEmbedUnimplementedCoffeeServiceServer() {}

// UnsafeCoffeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoffeeServiceServer will
// result in compilation errors.
type UnsafeCoffeeServiceServer interface {
	mustEmbedUnimplementedCoffeeServiceServer()
}

func RegisterCoffeeServiceServer(s grpc.ServiceRegistrar, srv CoffeeServiceServer) {
	s.RegisterService(&CoffeeService_ServiceDesc, srv)
}

func _CoffeeService_AddCoffee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCoffeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeServiceServer).AddCoffee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeService/AddCoffee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeServiceServer).AddCoffee(ctx, req.(*AddCoffeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeService_FindCoffeeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCoffeeById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeServiceServer).FindCoffeeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeService/FindCoffeeByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeServiceServer).FindCoffeeByID(ctx, req.(*FindCoffeeById))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeService_UpdateCoffeeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoffee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeServiceServer).UpdateCoffeeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeService/UpdateCoffeeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeServiceServer).UpdateCoffeeInfo(ctx, req.(*UpdateCoffee))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeService_DeleteCoffeeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCoffeeOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeServiceServer).DeleteCoffeeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeService/DeleteCoffeeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeServiceServer).DeleteCoffeeInfo(ctx, req.(*DeleteCoffeeOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoffeeService_FindAllCoffee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCoffee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoffeeServiceServer).FindAllCoffee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoffeeService/FindAllCoffee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoffeeServiceServer).FindAllCoffee(ctx, req.(*GetAllCoffee))
	}
	return interceptor(ctx, in, info, handler)
}

// CoffeeService_ServiceDesc is the grpc.ServiceDesc for CoffeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoffeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoffeeService",
	HandlerType: (*CoffeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCoffee",
			Handler:    _CoffeeService_AddCoffee_Handler,
		},
		{
			MethodName: "FindCoffeeByID",
			Handler:    _CoffeeService_FindCoffeeByID_Handler,
		},
		{
			MethodName: "UpdateCoffeeInfo",
			Handler:    _CoffeeService_UpdateCoffeeInfo_Handler,
		},
		{
			MethodName: "DeleteCoffeeInfo",
			Handler:    _CoffeeService_DeleteCoffeeInfo_Handler,
		},
		{
			MethodName: "FindAllCoffee",
			Handler:    _CoffeeService_FindAllCoffee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coffee.proto",
}
