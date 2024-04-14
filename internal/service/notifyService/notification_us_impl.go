package notifyService

import (
	"context"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2/log"
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/domain/entities/notication"
	userDevice "latipe-notification-service/internal/domain/entities/userDevice"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
	"latipe-notification-service/pkgUtils/fcm"
	"latipe-notification-service/pkgUtils/util/mapper"
	"time"
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

func (n notificationService) GetNotificationsOfUser(ctx context.Context, req *dto.GetNotificationsRequest) (*dto.GetNotificationsResponse, error) {
	notifications, total, err := n.notificationRepo.FindUnreadMessageOfUser(ctx, req.UserID, req.Query)
	if err != nil {
		return nil, err
	}

	resp := dto.GetNotificationsResponse{}

	resp.Total = total
	resp.Items = notifications
	resp.Size = req.Query.Size
	resp.HasMore = req.Query.GetHasMore(total)

	return &resp, nil
}

func (n notificationService) GetNotificationDetail(ctx context.Context, req *dto.GetNotificationDetailRequest) (*dto.GetNotificationDetailResponse, error) {
	noti, err := n.notificationRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if noti.Type == notication.NOTIFY_USER {
		noti.UnRead = false
		if err := n.notificationRepo.Update(ctx, noti); err != nil {
			return nil, err
		}
	}

	resp := dto.GetNotificationDetailResponse{}

	if err := mapper.BindingStruct(noti, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, err
}

func (n notificationService) TotalUnreadNotification(ctx context.Context, req *dto.TotalUnreadNotificationRequest) (*dto.TotalUnreadNotificationResponse, error) {
	total, err := n.notificationRepo.TotalUnreadMessageOfUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	resp := dto.TotalUnreadNotificationResponse{Total: int(total)}

	return &resp, nil
}

func (n notificationService) SendCampaignNotification(ctx context.Context, req *dto.SendCampaignNotificationRequest) (*dto.SendCampaignNotificationResponse, error) {
	noti := notication.Notification{
		Owner:  "all",
		Title:  req.Title,
		Image:  req.Image,
		Body:   req.Body,
		Type:   notication.NOTIFY_CAMPAIGN,
		UnRead: true,
	}

	entity, err := n.notificationRepo.Save(ctx, &noti)
	if err != nil {
		return nil, err
	}

	deviceTokens, err := n.userDeviceRepo.GetAllActiveDeviceToken(ctx)
	if err != nil {
		return nil, err
	}

	if len(deviceTokens) != 0 {
		if err := n.fbCloudMessage.SubscribeToTopic(ctx, deviceTokens, req.CampaignName); err != nil {
			return nil, err
		}

		message := messaging.Message{
			Notification: &messaging.Notification{
				Title:    req.Title,
				Body:     req.Body,
				ImageURL: req.Image,
			},
			Topic: req.CampaignName,
		}

		if err := n.fbCloudMessage.SendToTopic(ctx, &message); err != nil {
			return nil, err
		}

		if err := n.fbCloudMessage.UnsubscribeFromTopic(ctx, deviceTokens, req.CampaignName); err != nil {
			return nil, err
		}
	}

	resp := dto.SendCampaignNotificationResponse{}
	if err := mapper.BindingStruct(entity, &resp); err != nil {
		log.Error(err)
		return nil, err
	}
	return &resp, nil
}

func (n notificationService) SendNotification(ctx context.Context, req *dto.SendNotificationRequest) (*dto.SendNotificationResponse, error) {
	noti := notication.Notification{
		Owner:  req.UserID,
		Title:  req.Title,
		Image:  req.Image,
		Body:   req.Body,
		Type:   notication.NOTIFY_USER,
		UnRead: true,
		//ScheduleDisplay: time.Time{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	entity, err := n.notificationRepo.Save(ctx, &noti)
	if err != nil {
		return nil, err
	}

	if req.PushToDevice {
		devices, err := n.userDeviceRepo.FindActiveDeviceByUserID(ctx, req.UserID)
		if err != nil {
			return nil, err
		}

		var deviceTokens []string
		for _, device := range devices {
			if device.DeviceToken != "" {
				deviceTokens = append(deviceTokens, device.DeviceToken)
			}
		}

		if len(deviceTokens) != 0 {
			message := messaging.MulticastMessage{
				Notification: &messaging.Notification{
					Title:    req.Title,
					Body:     req.Body,
					ImageURL: req.Image,
				},
				Tokens: deviceTokens,
			}

			if err := n.fbCloudMessage.SendToMultipleDevices(ctx, &message); err != nil {
				return nil, err
			}
		}

	}

	resp := dto.SendNotificationResponse{}
	if err := mapper.BindingStruct(entity, &resp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &resp, nil
}

func (n notificationService) MarkAllRead(ctx context.Context, req *dto.MarkAsReadRequest) (*dto.MarkAsReadResponse, error) {
	err := n.notificationRepo.UpdateAllReadMessageOfUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (n notificationService) ClearAllNotification(ctx context.Context, req *dto.ClearNotificationRequest) (*dto.ClearNotificationResponse, error) {
	if err := n.notificationRepo.DeleteManyNotificationOfUser(ctx, req.UserID); err != nil {
		return nil, err
	}

	return nil, nil
}

func (n notificationService) RegisterNewUserDevice(ctx context.Context, req *dto.RegisterNewDevice) (*dto.RegisterNewDeviceResponse, error) {
	newDevice := userDevice.UserDevice{
		UserID:      req.UserID,
		DeviceInfo:  req.DeviceInfo,
		DeviceToken: req.DeviceToken,
		DeviceType:  req.GetDeviceType(),
		IsActive:    true,
		LoggedDate:  time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	entity, err := n.userDeviceRepo.Save(ctx, &newDevice)
	if err != nil {
		return nil, err
	}

	resp := dto.RegisterNewDeviceResponse{}

	if err := mapper.BindingStruct(entity, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
