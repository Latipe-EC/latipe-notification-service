package notifyService

import (
	"context"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	dataJson "latipe-notification-service/internal/domain/dto/schedule_callback_dto"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/internal/infrastructure/grpcExt/scheduleGrpc"
	"time"
)

func (n notificationService) sendCampaignToAllDevice(ctx context.Context, noti *notication.Notification) error {

	scheduleTime, err := noti.ParseScheduleToTime()
	if err != nil {
		log.Error(err)
	}

	if scheduleTime.After(time.Now()) {
		jsonData := dataJson.ScheduleJsonData{
			NotificationID: noti.ID.Hex(),
		}

		data := scheduleGrpc.ConvertStructToJSONString(jsonData)
		if len(data) == 0 {
			return fmt.Errorf("failed to convert struct to json string")
		}

		scheduleRequest := scheduleGrpc.CreateScheduleRequest{
			From:     n.config.Server.Name,
			Type:     scheduleGrpc.ONETIME,
			Deadline: noti.ScheduleDisplay,
			ReplyOn:  n.config.GrpcInfrastructure.ScheduleGRPC.CallbackURL,
			XApiKey:  n.config.Server.APIKey,
			Data:     data,
		}

		if _, err := n.scheduleGrpc.CreateSchedule(ctx, &scheduleRequest); err != nil {
			return err
		}

	} else {
		if err := n.sendCampaignToUserDevice(ctx, noti); err != nil {
			return err
		}
	}

	return nil

}

func (n notificationService) sendCampaignToUserDevice(ctx context.Context, noti *notication.Notification) error {
	deviceTokens, err := n.userDeviceRepo.GetAllActiveDeviceToken(ctx)
	if err != nil {
		return err
	}

	if len(deviceTokens) != 0 {
		if err := n.fbCloudMessage.SubscribeToTopic(ctx, deviceTokens, noti.CampaignTopic); err != nil {
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

		if err := n.fbCloudMessage.SendToTopic(ctx, &message); err != nil {
			return err
		}

		if err := n.fbCloudMessage.UnsubscribeFromTopic(ctx, deviceTokens, noti.CampaignTopic); err != nil {
			return err
		}
	}
	return nil
}
