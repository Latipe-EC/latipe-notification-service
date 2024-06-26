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
	// commands
	SendNotificationToUser(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	SendCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error)
	// query
	GetNotificationById(ctx context.Context, in *GetNotificationByIdRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error)
	GetNotificationUserId(ctx context.Context, in *GetNotificationByUserRequest, opts ...grpc.CallOption) (*GetNotificationByUserResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) SendNotificationToUser(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/SendNotificationToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendCampaign(ctx context.Context, in *CreateCampaignRequest, opts ...grpc.CallOption) (*CreateCampaignResponse, error) {
	out := new(CreateCampaignResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/SendCampaign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotificationById(ctx context.Context, in *GetNotificationByIdRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error) {
	out := new(GetNotificationResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/GetNotificationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetNotificationUserId(ctx context.Context, in *GetNotificationByUserRequest, opts ...grpc.CallOption) (*GetNotificationByUserResponse, error) {
	out := new(GetNotificationByUserResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/GetNotificationUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	// commands
	SendNotificationToUser(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	SendCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error)
	// query
	GetNotificationById(context.Context, *GetNotificationByIdRequest) (*GetNotificationResponse, error)
	GetNotificationUserId(context.Context, *GetNotificationByUserRequest) (*GetNotificationByUserResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) SendNotificationToUser(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotificationToUser not implemented")
}
func (UnimplementedNotificationServiceServer) SendCampaign(context.Context, *CreateCampaignRequest) (*CreateCampaignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCampaign not implemented")
}
func (UnimplementedNotificationServiceServer) GetNotificationById(context.Context, *GetNotificationByIdRequest) (*GetNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationById not implemented")
}
func (UnimplementedNotificationServiceServer) GetNotificationUserId(context.Context, *GetNotificationByUserRequest) (*GetNotificationByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationUserId not implemented")
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

func _NotificationService_SendNotificationToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendNotificationToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/SendNotificationToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendNotificationToUser(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/SendCampaign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendCampaign(ctx, req.(*CreateCampaignRequest))
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

func _NotificationService_GetNotificationUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetNotificationUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/GetNotificationUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetNotificationUserId(ctx, req.(*GetNotificationByUserRequest))
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
			MethodName: "SendNotificationToUser",
			Handler:    _NotificationService_SendNotificationToUser_Handler,
		},
		{
			MethodName: "SendCampaign",
			Handler:    _NotificationService_SendCampaign_Handler,
		},
		{
			MethodName: "GetNotificationById",
			Handler:    _NotificationService_GetNotificationById_Handler,
		},
		{
			MethodName: "GetNotificationUserId",
			Handler:    _NotificationService_GetNotificationUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/grpc-service/notificationGrpc/notification_service.proto",
}
