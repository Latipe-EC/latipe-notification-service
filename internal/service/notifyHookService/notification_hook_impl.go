package notifyHookService

import (
	"context"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2/log"
	"latipe-notification-service/config"
	dto "latipe-notification-service/internal/domain/dto/schedule_callback_dto"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
	"latipe-notification-service/pkgUtils/fcm"
)

type notifyHookService struct {
	config          *config.AppConfig
	notiRepos       notifyRepos.NotificationRepository
	userDeviceRepos userDeviceRepos.UserDeviceRepository
	fcmClient       fcm.NotificationCloudMessage
}

func NewNotifyHookService(config *config.AppConfig, notiRepos notifyRepos.NotificationRepository,
	userDeviceRepos userDeviceRepos.UserDeviceRepository, fcmClient fcm.NotificationCloudMessage) NotifyHookService {
	return &notifyHookService{
		config:          config,
		notiRepos:       notiRepos,
		userDeviceRepos: userDeviceRepos,
		fcmClient:       fcmClient,
	}
}

func (n notifyHookService) CallBackFromScheduleService(ctx context.Context, req *dto.ScheduleCallbackRequest) error {
	jsonData := req.ParseDataToStruct()
	if jsonData != nil {
		// get notification by id
		notification, err := n.notiRepos.FindByID(ctx, jsonData.NotificationID)
		if err != nil {
			return err
		}

		if notification.IsActive {
			err := n.sendCampaignToUserDevice(ctx, notification)
			if err != nil {
				log.Error(err)
				return err
			}
		}
	}
	return nil
}

func (n notifyHookService) sendCampaignToUserDevice(ctx context.Context, noti *notication.Notification) error {
	deviceTokens, err := n.userDeviceRepos.GetAllActiveDeviceToken(ctx)
	if err != nil {
		return err
	}

	if len(deviceTokens) != 0 {
		if err := n.fcmClient.SubscribeToTopic(ctx, deviceTokens, noti.CampaignTopic); err != nil {
			return err
		}

		message := messaging.Message{
			Notification: &messaging.Notification{
				Title:    noti.Title,
				Body:     noti.Body,
				ImageURL: noti.Image,
			},
			Topic: noti.CampaignTopic,
		}

		if err := n.fcmClient.SendToTopic(ctx, &message); err != nil {
			return err
		}

		if err := n.fcmClient.UnsubscribeFromTopic(ctx, deviceTokens, noti.CampaignTopic); err != nil {
			return err
		}
	}
	return nil
}
