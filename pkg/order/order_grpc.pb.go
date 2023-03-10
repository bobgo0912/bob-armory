// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	GetOrdersByIds(ctx context.Context, in *OrderIdsReq, opts ...grpc.CallOption) (*OrderResp, error)
	GetOrdersByPeriod(ctx context.Context, in *OrderPeriodReq, opts ...grpc.CallOption) (*OrderResp, error)
	GetCardsByPeriod(ctx context.Context, in *CardPeriodReq, opts ...grpc.CallOption) (*CardResp, error)
	GetUnSettleCardsIdsByPeriod(ctx context.Context, in *CardsIdsReq, opts ...grpc.CallOption) (*CardsIdsResp, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) GetOrdersByIds(ctx context.Context, in *OrderIdsReq, opts ...grpc.CallOption) (*OrderResp, error) {
	out := new(OrderResp)
	err := c.cc.Invoke(ctx, "/order.Order/GetOrdersByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrdersByPeriod(ctx context.Context, in *OrderPeriodReq, opts ...grpc.CallOption) (*OrderResp, error) {
	out := new(OrderResp)
	err := c.cc.Invoke(ctx, "/order.Order/GetOrdersByPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetCardsByPeriod(ctx context.Context, in *CardPeriodReq, opts ...grpc.CallOption) (*CardResp, error) {
	out := new(CardResp)
	err := c.cc.Invoke(ctx, "/order.Order/GetCardsByPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetUnSettleCardsIdsByPeriod(ctx context.Context, in *CardsIdsReq, opts ...grpc.CallOption) (*CardsIdsResp, error) {
	out := new(CardsIdsResp)
	err := c.cc.Invoke(ctx, "/order.Order/GetUnSettleCardsIdsByPeriod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	GetOrdersByIds(context.Context, *OrderIdsReq) (*OrderResp, error)
	GetOrdersByPeriod(context.Context, *OrderPeriodReq) (*OrderResp, error)
	GetCardsByPeriod(context.Context, *CardPeriodReq) (*CardResp, error)
	GetUnSettleCardsIdsByPeriod(context.Context, *CardsIdsReq) (*CardsIdsResp, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) GetOrdersByIds(context.Context, *OrderIdsReq) (*OrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByIds not implemented")
}
func (UnimplementedOrderServer) GetOrdersByPeriod(context.Context, *OrderPeriodReq) (*OrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrdersByPeriod not implemented")
}
func (UnimplementedOrderServer) GetCardsByPeriod(context.Context, *CardPeriodReq) (*CardResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardsByPeriod not implemented")
}
func (UnimplementedOrderServer) GetUnSettleCardsIdsByPeriod(context.Context, *CardsIdsReq) (*CardsIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnSettleCardsIdsByPeriod not implemented")
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

func _Order_GetOrdersByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrdersByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/GetOrdersByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrdersByIds(ctx, req.(*OrderIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrdersByPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPeriodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrdersByPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/GetOrdersByPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrdersByPeriod(ctx, req.(*OrderPeriodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetCardsByPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardPeriodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetCardsByPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/GetCardsByPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetCardsByPeriod(ctx, req.(*CardPeriodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetUnSettleCardsIdsByPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardsIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetUnSettleCardsIdsByPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/GetUnSettleCardsIdsByPeriod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetUnSettleCardsIdsByPeriod(ctx, req.(*CardsIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrdersByIds",
			Handler:    _Order_GetOrdersByIds_Handler,
		},
		{
			MethodName: "GetOrdersByPeriod",
			Handler:    _Order_GetOrdersByPeriod_Handler,
		},
		{
			MethodName: "GetCardsByPeriod",
			Handler:    _Order_GetCardsByPeriod_Handler,
		},
		{
			MethodName: "GetUnSettleCardsIdsByPeriod",
			Handler:    _Order_GetUnSettleCardsIdsByPeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
