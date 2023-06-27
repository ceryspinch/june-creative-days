// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: ticket_manager/ticket.proto

package ticket_manager

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

// TicketManagerClient is the client API for TicketManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketManagerClient interface {
	// Create
	BuyTicket(ctx context.Context, in *BuyTicketRequest, opts ...grpc.CallOption) (*BuyTicketResponse, error)
	// Read
	ListTicket(ctx context.Context, in *ListTicketRequest, opts ...grpc.CallOption) (*ListTicketResponse, error)
	// Read
	ListAllTickets(ctx context.Context, in *ListAllTicketsRequest, opts ...grpc.CallOption) (*ListAllTicketsResponse, error)
	// Update
	UpdateTicketInformation(ctx context.Context, in *UpdateTicketInformationRequest, opts ...grpc.CallOption) (*UpdateTicketInformationResponse, error)
	// Delete
	DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error)
}

type ticketManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketManagerClient(cc grpc.ClientConnInterface) TicketManagerClient {
	return &ticketManagerClient{cc}
}

func (c *ticketManagerClient) BuyTicket(ctx context.Context, in *BuyTicketRequest, opts ...grpc.CallOption) (*BuyTicketResponse, error) {
	out := new(BuyTicketResponse)
	err := c.cc.Invoke(ctx, "/ticket_manager.TicketManager/BuyTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketManagerClient) ListTicket(ctx context.Context, in *ListTicketRequest, opts ...grpc.CallOption) (*ListTicketResponse, error) {
	out := new(ListTicketResponse)
	err := c.cc.Invoke(ctx, "/ticket_manager.TicketManager/ListTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketManagerClient) ListAllTickets(ctx context.Context, in *ListAllTicketsRequest, opts ...grpc.CallOption) (*ListAllTicketsResponse, error) {
	out := new(ListAllTicketsResponse)
	err := c.cc.Invoke(ctx, "/ticket_manager.TicketManager/ListAllTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketManagerClient) UpdateTicketInformation(ctx context.Context, in *UpdateTicketInformationRequest, opts ...grpc.CallOption) (*UpdateTicketInformationResponse, error) {
	out := new(UpdateTicketInformationResponse)
	err := c.cc.Invoke(ctx, "/ticket_manager.TicketManager/UpdateTicketInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketManagerClient) DeleteTicket(ctx context.Context, in *DeleteTicketRequest, opts ...grpc.CallOption) (*DeleteTicketResponse, error) {
	out := new(DeleteTicketResponse)
	err := c.cc.Invoke(ctx, "/ticket_manager.TicketManager/DeleteTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketManagerServer is the server API for TicketManager service.
// All implementations must embed UnimplementedTicketManagerServer
// for forward compatibility
type TicketManagerServer interface {
	// Create
	BuyTicket(context.Context, *BuyTicketRequest) (*BuyTicketResponse, error)
	// Read
	ListTicket(context.Context, *ListTicketRequest) (*ListTicketResponse, error)
	// Read
	ListAllTickets(context.Context, *ListAllTicketsRequest) (*ListAllTicketsResponse, error)
	// Update
	UpdateTicketInformation(context.Context, *UpdateTicketInformationRequest) (*UpdateTicketInformationResponse, error)
	// Delete
	DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error)
	mustEmbedUnimplementedTicketManagerServer()
}

// UnimplementedTicketManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTicketManagerServer struct {
}

func (UnimplementedTicketManagerServer) BuyTicket(context.Context, *BuyTicketRequest) (*BuyTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyTicket not implemented")
}
func (UnimplementedTicketManagerServer) ListTicket(context.Context, *ListTicketRequest) (*ListTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTicket not implemented")
}
func (UnimplementedTicketManagerServer) ListAllTickets(context.Context, *ListAllTicketsRequest) (*ListAllTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllTickets not implemented")
}
func (UnimplementedTicketManagerServer) UpdateTicketInformation(context.Context, *UpdateTicketInformationRequest) (*UpdateTicketInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTicketInformation not implemented")
}
func (UnimplementedTicketManagerServer) DeleteTicket(context.Context, *DeleteTicketRequest) (*DeleteTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTicket not implemented")
}
func (UnimplementedTicketManagerServer) mustEmbedUnimplementedTicketManagerServer() {}

// UnsafeTicketManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketManagerServer will
// result in compilation errors.
type UnsafeTicketManagerServer interface {
	mustEmbedUnimplementedTicketManagerServer()
}

func RegisterTicketManagerServer(s grpc.ServiceRegistrar, srv TicketManagerServer) {
	s.RegisterService(&TicketManager_ServiceDesc, srv)
}

func _TicketManager_BuyTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketManagerServer).BuyTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_manager.TicketManager/BuyTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketManagerServer).BuyTicket(ctx, req.(*BuyTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketManager_ListTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketManagerServer).ListTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_manager.TicketManager/ListTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketManagerServer).ListTicket(ctx, req.(*ListTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketManager_ListAllTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketManagerServer).ListAllTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_manager.TicketManager/ListAllTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketManagerServer).ListAllTickets(ctx, req.(*ListAllTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketManager_UpdateTicketInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTicketInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketManagerServer).UpdateTicketInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_manager.TicketManager/UpdateTicketInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketManagerServer).UpdateTicketInformation(ctx, req.(*UpdateTicketInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketManager_DeleteTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketManagerServer).DeleteTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_manager.TicketManager/DeleteTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketManagerServer).DeleteTicket(ctx, req.(*DeleteTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketManager_ServiceDesc is the grpc.ServiceDesc for TicketManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket_manager.TicketManager",
	HandlerType: (*TicketManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyTicket",
			Handler:    _TicketManager_BuyTicket_Handler,
		},
		{
			MethodName: "ListTicket",
			Handler:    _TicketManager_ListTicket_Handler,
		},
		{
			MethodName: "ListAllTickets",
			Handler:    _TicketManager_ListAllTickets_Handler,
		},
		{
			MethodName: "UpdateTicketInformation",
			Handler:    _TicketManager_UpdateTicketInformation_Handler,
		},
		{
			MethodName: "DeleteTicket",
			Handler:    _TicketManager_DeleteTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket_manager/ticket.proto",
}
