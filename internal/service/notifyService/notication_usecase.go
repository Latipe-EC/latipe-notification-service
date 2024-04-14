package notifyService

import (
	"context"
	"latipe-notification-service/internal/domain/dto"
)

type NotificationService interface {
	// queries
	GetNotificationsOfUser(ctx context.Context, req *dto.GetNotificationsRequest) (*dto.GetNotificationsResponse, error)
	GetNotificationDetail(ctx context.Context, req *dto.GetNotificationDetailRequest) (*dto.GetNotificationDetailResponse, error)
	TotalUnreadNotification(ctx context.Context, req *dto.TotalUnreadNotificationRequest) (*dto.TotalUnreadNotificationResponse, error)
	SendCampaignNotification(ctx context.Context, req *dto.SendCampaignNotificationRequest) (*dto.SendCampaignNotificationResponse, error)
	// commands
	RegisterNewUserDevice(ctx context.Context, req *dto.RegisterNewDevice) (*dto.RegisterNewDeviceResponse, error)
	SendNotification(ctx context.Context, req *dto.SendNotificationRequest) (*dto.SendNotificationResponse, error)
	MarkAllRead(ctx context.Context, req *dto.MarkAsReadRequest) (*dto.MarkAsReadResponse, error)
	ClearAllNotification(ctx context.Context, req *dto.ClearNotificationRequest) (*dto.ClearNotificationResponse, error)
}
