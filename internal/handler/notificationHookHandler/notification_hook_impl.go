package notificationHookHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	dto "latipe-notification-service/internal/domain/dto/schedule_callback_dto"
	"latipe-notification-service/internal/service/notifyHookService"
	"latipe-notification-service/pkgUtils/util/errorUtils"
	responses "latipe-notification-service/pkgUtils/util/response"
	"latipe-notification-service/pkgUtils/util/valid"
)

type notificationHookHandler struct {
	notificationServ notifyHookService.NotifyHookService
}

func NewNotificationHookHandler(notificationServ notifyHookService.NotifyHookService) NotificationHookHandler {
	return &notificationHookHandler{
		notificationServ: notificationServ,
	}
}

func (n notificationHookHandler) HandleScheduleCallback(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.ScheduleCallbackRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Errorf("Error parsing request: %s", err)
		return errorUtils.ErrInternalServer
	}

	if err := valid.GetValidator().Validate(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	err := n.notificationServ.CallBackFromScheduleService(context, &req)
	if err != nil {
		log.Error(err)
		return errorUtils.ErrInternalServer
	}

	return responses.DefaultSuccess.JSON(ctx)
}
