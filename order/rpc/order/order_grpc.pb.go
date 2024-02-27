// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.11.4
// source: order.proto

package order

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

const (
	Order_Orders_FullMethodName               = "/order.order/Orders"
	Order_CreateOrder_FullMethodName          = "/order.order/CreateOrder"
	Order_CreateOrderCheck_FullMethodName     = "/order.order/CreateOrderCheck"
	Order_RollbackOrder_FullMethodName        = "/order.order/RollbackOrder"
	Order_CreateOrderDTM_FullMethodName       = "/order.order/CreateOrderDTM"
	Order_CreateOrderDTMRevert_FullMethodName = "/order.order/CreateOrderDTMRevert"
	Order_GetOrderByOrdernum_FullMethodName   = "/order.order/GetOrderByOrdernum"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	// 个人订单分页
	Orders(ctx context.Context, in *OrdersReq, opts ...grpc.CallOption) (*OrdersRsp, error)
	// 创建订单
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error)
	// 订单创建验证
	CreateOrderCheck(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error)
	// 回滚订单
	RollbackOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error)
	// 创建订单try
	CreateOrderDTM(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderRsp, error)
	// 回撤
	CreateOrderDTMRevert(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderRsp, error)
	GetOrderByOrdernum(ctx context.Context, in *GetOrderByOrdernumReq, opts ...grpc.CallOption) (*GetOrderByOrdernumRsp, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) Orders(ctx context.Context, in *OrdersReq, opts ...grpc.CallOption) (*OrdersRsp, error) {
	out := new(OrdersRsp)
	err := c.cc.Invoke(ctx, Order_Orders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error) {
	out := new(CreateOrderRsp)
	err := c.cc.Invoke(ctx, Order_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrderCheck(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error) {
	out := new(CreateOrderRsp)
	err := c.cc.Invoke(ctx, Order_CreateOrderCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) RollbackOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error) {
	out := new(CreateOrderRsp)
	err := c.cc.Invoke(ctx, Order_RollbackOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrderDTM(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderRsp, error) {
	out := new(AddOrderRsp)
	err := c.cc.Invoke(ctx, Order_CreateOrderDTM_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrderDTMRevert(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderRsp, error) {
	out := new(AddOrderRsp)
	err := c.cc.Invoke(ctx, Order_CreateOrderDTMRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderByOrdernum(ctx context.Context, in *GetOrderByOrdernumReq, opts ...grpc.CallOption) (*GetOrderByOrdernumRsp, error) {
	out := new(GetOrderByOrdernumRsp)
	err := c.cc.Invoke(ctx, Order_GetOrderByOrdernum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	// 个人订单分页
	Orders(context.Context, *OrdersReq) (*OrdersRsp, error)
	// 创建订单
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRsp, error)
	// 订单创建验证
	CreateOrderCheck(context.Context, *CreateOrderReq) (*CreateOrderRsp, error)
	// 回滚订单
	RollbackOrder(context.Context, *CreateOrderReq) (*CreateOrderRsp, error)
	// 创建订单try
	CreateOrderDTM(context.Context, *AddOrderReq) (*AddOrderRsp, error)
	// 回撤
	CreateOrderDTMRevert(context.Context, *AddOrderReq) (*AddOrderRsp, error)
	GetOrderByOrdernum(context.Context, *GetOrderByOrdernumReq) (*GetOrderByOrdernumRsp, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) Orders(context.Context, *OrdersReq) (*OrdersRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Orders not implemented")
}
func (UnimplementedOrderServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) CreateOrderCheck(context.Context, *CreateOrderReq) (*CreateOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderCheck not implemented")
}
func (UnimplementedOrderServer) RollbackOrder(context.Context, *CreateOrderReq) (*CreateOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackOrder not implemented")
}
func (UnimplementedOrderServer) CreateOrderDTM(context.Context, *AddOrderReq) (*AddOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderDTM not implemented")
}
func (UnimplementedOrderServer) CreateOrderDTMRevert(context.Context, *AddOrderReq) (*AddOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderDTMRevert not implemented")
}
func (UnimplementedOrderServer) GetOrderByOrdernum(context.Context, *GetOrderByOrdernumReq) (*GetOrderByOrdernumRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByOrdernum not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_Orders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Orders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Orders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Orders(ctx, req.(*OrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrderCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrderCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrderCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrderCheck(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_RollbackOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).RollbackOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_RollbackOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).RollbackOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrderDTM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrderDTM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrderDTM_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrderDTM(ctx, req.(*AddOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrderDTMRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrderDTMRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrderDTMRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrderDTMRevert(ctx, req.(*AddOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderByOrdernum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByOrdernumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderByOrdernum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetOrderByOrdernum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderByOrdernum(ctx, req.(*GetOrderByOrdernumReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Orders",
			Handler:    _Order_Orders_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "CreateOrderCheck",
			Handler:    _Order_CreateOrderCheck_Handler,
		},
		{
			MethodName: "RollbackOrder",
			Handler:    _Order_RollbackOrder_Handler,
		},
		{
			MethodName: "CreateOrderDTM",
			Handler:    _Order_CreateOrderDTM_Handler,
		},
		{
			MethodName: "CreateOrderDTMRevert",
			Handler:    _Order_CreateOrderDTMRevert_Handler,
		},
		{
			MethodName: "GetOrderByOrdernum",
			Handler:    _Order_GetOrderByOrdernum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
