package notifyService

import "latipe-notification-service/internal/domain/dto"

type NotificationService interface {
	// queries
	GetNotificationsOfUser(req *dto.GetNotificationsRequest) (*dto.GetNotificationsResponse, error)
	GetNotificationDetail(req *dto.GetNotificationDetailRequest) (*dto.GetNotificationDetailResponse, error)
	TotalUnreadNotification(req *dto.TotalUnreadNotificationRequest) (*dto.TotalUnreadNotificationResponse, error)
	// commands
	SendNotification(req *dto.SendNotificationRequest) (*dto.SendNotificationResponse, error)
	MarkAsRead(req *dto.MarkAsReadRequest) (*dto.MarkAsReadResponse, error)
	ClearAllNotification(req *dto.ClearNotificationRequest) (*dto.ClearNotificationResponse, error)
}
