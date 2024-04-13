// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: internal/grpc-service/notificationGrpc/notification_service.proto

package notificationGrpc

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

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	CreateOrderNotification(ctx context.Context, in *CreateOrderNotificationRequest, opts ...grpc.CallOption) (*CreateOrderNotificationResponse, error)
	GetNotificationById(ctx context.Context, in *GetNotificationByIdRequest, opts ...grpc.CallOption) (*GetNotificationByIdResponse, error)
	GetNotificationByUserId(ctx context.Context, in *GetNotificationByUserIdRequest, opts ...grpc.CallOption) (*GetNotificationByUserIdResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) CreateOrderNotification(ctx context.Context, in *CreateOrderNotificationRequest, opts ...grpc.CallOption) (*CreateOrderNotificationResponse, error) {
	out := new(CreateOrderNotificationResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/CreateOrderNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotificationById(ctx context.Context, in *GetNotificationByIdRequest, opts ...grpc.CallOption) (*GetNotificationByIdResponse, error) {
	out := new(GetNotificationByIdResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/GetNotificationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotificationByUserId(ctx context.Context, in *GetNotificationByUserIdRequest, opts ...grpc.CallOption) (*GetNotificationByUserIdResponse, error) {
	out := new(GetNotificationByUserIdResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/GetNotificationByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	CreateOrderNotification(context.Context, *CreateOrderNotificationRequest) (*CreateOrderNotificationResponse, error)
	GetNotificationById(context.Context, *GetNotificationByIdRequest) (*GetNotificationByIdResponse, error)
	GetNotificationByUserId(context.Context, *GetNotificationByUserIdRequest) (*GetNotificationByUserIdResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) CreateOrderNotification(context.Context, *CreateOrderNotificationRequest) (*CreateOrderNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderNotification not implemented")
}
func (UnimplementedNotificationServiceServer) GetNotificationById(context.Context, *GetNotificationByIdRequest) (*GetNotificationByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationById not implemented")
}
func (UnimplementedNotificationServiceServer) GetNotificationByUserId(context.Context, *GetNotificationByUserIdRequest) (*GetNotificationByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationByUserId not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_CreateOrderNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateOrderNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/CreateOrderNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateOrderNotification(ctx, req.(*CreateOrderNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetNotificationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotificationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/GetNotificationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotificationById(ctx, req.(*GetNotificationByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetNotificationByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotificationByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/GetNotificationByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotificationByUserId(ctx, req.(*GetNotificationByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrderNotification",
			Handler:    _NotificationService_CreateOrderNotification_Handler,
		},
		{
			MethodName: "GetNotificationById",
			Handler:    _NotificationService_GetNotificationById_Handler,
		},
		{
			MethodName: "GetNotificationByUserId",
			Handler:    _NotificationService_GetNotificationByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpc-service/notificationGrpc/notification_service.proto",
}