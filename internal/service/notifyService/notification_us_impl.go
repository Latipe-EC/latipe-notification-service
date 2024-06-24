package notifyService

import (
	"context"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2/log"
	"latipe-notification-service/config"
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/domain/entities/notication"
	userDevice "latipe-notification-service/internal/domain/entities/userDevice"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
	"latipe-notification-service/internal/infrastructure/grpcExt/scheduleGrpc"
	"latipe-notification-service/internal/service/utils"
	"latipe-notification-service/pkgUtils/fcm"
	"latipe-notification-service/pkgUtils/util/errorUtils"
	"latipe-notification-service/pkgUtils/util/mapper"
	"time"
)

type notificationService struct {
	config           *config.AppConfig
	notificationRepo notifyRepos.NotificationRepository
	userDeviceRepo   userDeviceRepos.UserDeviceRepository
	fbCloudMessage   fcm.NotificationCloudMessage
	scheduleGrpc     scheduleGrpc.ScheduleServiceClient
}

func NewNotificationService(config *config.AppConfig, notificationRepo notifyRepos.NotificationRepository,
	userDeviceRepo userDeviceRepos.UserDeviceRepository,
	fbCloudMessage fcm.NotificationCloudMessage, scheduleGrpc scheduleGrpc.ScheduleServiceClient) NotificationService {
	return &notificationService{
		config:           config,
		notificationRepo: notificationRepo,
		userDeviceRepo:   userDeviceRepo,
		fbCloudMessage:   fbCloudMessage,
		scheduleGrpc:     scheduleGrpc,
	}
}

func (n notificationService) GetNotificationsOfUser(ctx context.Context, req *dto.GetNotificationsRequest) (*dto.GetNotificationsResponse, error) {
	notifications, total, err := n.notificationRepo.FindUnreadMessageOfUser(ctx, req.UserID, req.Query)
	if err != nil {
		return nil, err
	}

	var responses []dto.GetNotificationDetailResponse

	if err := mapper.BindingStruct(notifications, &responses); err != nil {
		return nil, err
	}

	resp := dto.GetNotificationsResponse{}

	resp.Total = total
	resp.Items = responses
	resp.Size = req.Query.Size
	resp.Page = req.Query.Page
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
		if err := n.notificationRepo.UpdateReadStatusNotification(ctx, noti); err != nil {
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

func (n notificationService) SendCampaignInternalService(ctx context.Context, req *dto.SendCampaignNotificationRequest) (*dto.SendCampaignNotificationResponse, error) {
	schedule, err := utils.RetrieveCurrentDate(req.ScheduleDisplay)
	if err != nil {
		log.Error(err)
		return nil, errorUtils.ErrParseDatetimeParameters
	}

	if schedule.After(time.Now()) {
		noti := notication.NewNotification()

		noti.OwnerID = "all"
		noti.Title = req.Title
		noti.Image = req.Image
		noti.Body = req.Body
		noti.CampaignTopic = req.CampaignTopic
		noti.Type = notication.NOTIFY_CAMPAIGN
		noti.UnRead = true
		noti.CreatedBy = "INTERNAL_SERVICE"
		noti.ScheduleDisplay = req.ScheduleDisplay

		insertedNoti, err := n.notificationRepo.Save(ctx, &noti)
		if err != nil {
			return nil, err
		}

		if err := n.sendCampaignToAllDevice(ctx, insertedNoti); err != nil {
			return nil, err
		}

		resp := dto.SendCampaignNotificationResponse{}
		return &resp, nil
	}

	return nil, errorUtils.ErrInvalidDatetimeParameters
}

func (n notificationService) SendNotification(ctx context.Context, req *dto.SendNotificationRequest) (*dto.SendNotificationResponse, error) {
	noti := notication.NewNotification()

	noti.OwnerID = req.UserID
	noti.Title = req.Title
	noti.Image = req.Image
	noti.Body = req.Body
	noti.Type = notication.NOTIFY_USER

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

func (n notificationService) AdminGetAllCampaigns(ctx context.Context, req *dto.AdminGetAllCampaignRequest) (*dto.AdminGetAllCampaignResponse, error) {
	campaigns, total, err := n.notificationRepo.FindAllCampaigns(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	var responseItems []dto.AdminGetCampaignDetailResponse
	if err := mapper.BindingStruct(campaigns, &responseItems); err != nil {
		return nil, err
	}

	resp := dto.AdminGetAllCampaignResponse{}

	resp.Total = total
	resp.Items = responseItems
	resp.Size = req.Query.Size
	resp.Page = req.Query.Page
	resp.HasMore = req.Query.GetHasMore(total)

	return &resp, nil
}

func (n notificationService) AdminCreateCampaign(ctx context.Context, req *dto.AdminCreateCampaignRequest) (*dto.AdminCreateCampaignResponse, error) {
	schedule, err := utils.RetrieveCurrentDate(req.ScheduleDisplay)
	if err != nil {
		log.Error(err)
		return nil, errorUtils.ErrParseDatetimeParameters
	}

	if schedule.After(time.Now()) {
		noti := notication.NewNotification()
		noti.OwnerID = "all"
		noti.CampaignTopic = req.CampaignTopic
		noti.Title = req.Title
		noti.Image = req.Image
		noti.Body = req.Body
		noti.Type = notication.NOTIFY_CAMPAIGN
		noti.CreatedBy = req.UserID
		noti.ScheduleDisplay = req.ScheduleDisplay

		insertedNoti, err := n.notificationRepo.Save(ctx, &noti)
		if err != nil {
			return nil, err
		}

		if err := n.sendCampaignToAllDevice(ctx, insertedNoti); err != nil {
			return nil, err
		}

		resp := dto.AdminCreateCampaignResponse{
			ID: insertedNoti.ID.Hex(),
		}

		return &resp, nil
	}

	return nil, errorUtils.ErrInvalidDatetimeParameters

}

func (n notificationService) AdminRecallCampaign(ctx context.Context, req *dto.RecallNotificationRequest) (*dto.RecallNotificationRequest, error) {
	noti, err := n.notificationRepo.FindByID(ctx, req.NotificationID)
	if err != nil {
		return nil, err
	}

	scheduleTime, err := noti.ParseScheduleToTime()
	if err != nil {
		log.Error(err)
	}

	if scheduleTime.After(time.Now()) {
		noti.RecallReason = req.Reason
		if err := n.notificationRepo.RecallCampaign(ctx, noti); err != nil {
			return nil, err
		}
		return req, nil
	}

	return nil, errorUtils.ErrInvalidParameters
}
