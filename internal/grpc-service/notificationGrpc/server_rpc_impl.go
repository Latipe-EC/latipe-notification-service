package notificationGrpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	notifyDTO "latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/service/notifyService"
	"latipe-notification-service/pkgUtils/util/mapper"
	"latipe-notification-service/pkgUtils/util/valid"
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

	dto := notifyDTO.SendNotificationRequest{}
	if err := mapper.BindingStruct(request, &dto); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	if err := valid.GetValidator().Validate(&dto); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint(err))
	}

	data, err := n.notifyService.SendNotification(ctx, &dto)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	resp := CreateNotificationResponse{NotificationId: data.ID}

	return &resp, nil
}

func (n notificationGrpcServer) SendCampaign(ctx context.Context, request *CreateCampaignRequest) (*CreateCampaignResponse, error) {
	dto := notifyDTO.SendCampaignNotificationRequest{}
	if err := mapper.BindingStruct(request, &dto); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	if err := valid.GetValidator().Validate(&dto); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint(err))
	}

	data, err := n.notifyService.SendCampaignInternalService(ctx, &dto)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	resp := CreateCampaignResponse{NotificationId: data.ID}

	return &resp, nil
}

func (n notificationGrpcServer) GetNotificationById(ctx context.Context, request *GetNotificationByIdRequest) (*GetNotificationResponse, error) {
	dto := notifyDTO.GetNotificationDetailRequest{}
	dto.ID = request.NotificationId

	//if err := valid.GetValidator().Validate(&dto); err != nil {
	//
	//return nil, status.Error(codes.InvalidArgument, fmt.Sprint(err))
	//}

	data, err := n.notifyService.GetNotificationDetail(ctx, &dto)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	resp := GetNotificationResponse{}

	if err := mapper.BindingStruct(data, &resp); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &resp, nil
}

func (n notificationGrpcServer) GetNotificationUserId(ctx context.Context, request *GetNotificationByUserRequest) (*GetNotificationByUserResponse, error) {
	dto := notifyDTO.GetNotificationsRequest{}
	dto.UserID = request.UserId
	dto.Query.Size = int(request.Size)
	dto.Query.Page = int(request.Page)

	//if err := valid.GetValidator().Validate(&dto); err != nil {
	//
	//return nil, status.Error(codes.InvalidArgument, fmt.Sprint(err))
	//}

	data, err := n.notifyService.GetNotificationsOfUser(ctx, &dto)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	resp := GetNotificationByUserResponse{}

	if err := mapper.BindingStruct(data, &resp); err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &resp, nil
}

func (n notificationGrpcServer) mustEmbedUnimplementedNotificationServiceServer() {
	//TODO implement me
	panic("implement me")
}
