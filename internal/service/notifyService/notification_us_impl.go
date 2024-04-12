package notifyService

import (
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
	"latipe-notification-service/pkgUtils/fcm"
)

type notificationService struct {
	notificationRepo notifyRepos.NotificationRepository
	userDeviceRepo   userDeviceRepos.UserDeviceRepository
	fbCloudMessage   *fcm.FirebaseCloudMessage
}

func NewNotificationService(notificationRepo notifyRepos.NotificationRepository,
	userDeviceRepo userDeviceRepos.UserDeviceRepository,
	fbCloudMessage *fcm.FirebaseCloudMessage) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		userDeviceRepo:   userDeviceRepo,
		fbCloudMessage:   fbCloudMessage,
	}
}

func (n notificationService) GetNotificationsOfUser(req *dto.GetNotificationsRequest) (*dto.GetNotificationsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationService) GetNotificationDetail(req *dto.GetNotificationDetailRequest) (*dto.GetNotificationDetailResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationService) TotalUnreadNotification(req *dto.TotalUnreadNotificationRequest) (*dto.TotalUnreadNotificationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationService) SendNotification(req *dto.SendNotificationRequest) (*dto.SendNotificationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationService) MarkAsRead(req *dto.MarkAsReadRequest) (*dto.MarkAsReadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationService) ClearAllNotification(req *dto.ClearNotificationRequest) (*dto.ClearNotificationResponse, error) {
	//TODO implement me
	panic("implement me")
}
