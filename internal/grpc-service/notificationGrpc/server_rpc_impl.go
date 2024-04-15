package notificationGrpc

import (
	"context"
	"latipe-notification-service/internal/service/notifyService"
)

type notificationGrpcServer struct {
	UnimplementedNotificationServiceServer
	notifyService notifyService.NotificationService
}

func NewNotificationGrpcServer(notifyService notifyService.NotificationService) NotificationServiceServer {
	return &notificationGrpcServer{
		notifyService: notifyService,
	}
}

func (n notificationGrpcServer) SendNotificationToUser(ctx context.Context, request *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationGrpcServer) SendCampaign(ctx context.Context, request *CreateCampaignRequest) (*CreateCampaignResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationGrpcServer) GetNotificationById(ctx context.Context, request *GetNotificationByIdRequest) (*GetNotificationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationGrpcServer) GetNotificationUserId(ctx context.Context, request *GetNotificationByUserRequest) (*GetNotificationByUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationGrpcServer) mustEmbedUnimplementedNotificationServiceServer() {
	//TODO implement me
	panic("implement me")
}
