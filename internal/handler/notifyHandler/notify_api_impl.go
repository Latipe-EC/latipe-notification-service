package notifyHandler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/middleware"
	"latipe-notification-service/internal/service/notifyService"
	"latipe-notification-service/pkgUtils/util/errorUtils"
	"latipe-notification-service/pkgUtils/util/pagable"
	responses "latipe-notification-service/pkgUtils/util/response"
)

type notifyHandler struct {
	notifyService notifyService.NotificationService
}

func NewNotifyHandler(notifyService notifyService.NotificationService) NotifyHandler {
	return &notifyHandler{
		notifyService: notifyService,
	}
}

func (n notifyHandler) GetNotificationsOfUser(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		return errorUtils.ErrUnauthenticated
	}

	query, err := pagable.GetQueryFromFiberCtx(ctx)
	if err != nil {
		return errorUtils.ErrInvalidParameters
	}

	req := dto.GetNotificationsRequest{
		UserID: userId,
		Query:  query,
	}

	data, err := n.notifyService.GetNotificationsOfUser(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	resp := responses.DefaultSuccess
	resp.Data = data
	return resp.JSON(ctx)
}

func (n notifyHandler) GetNotificationDetail(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.GetNotificationDetailRequest{}
	if err := ctx.ParamsParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.GetNotificationDetail(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	resp := responses.DefaultSuccess
	resp.Data = data
	return resp.JSON(ctx)
}

func (n notifyHandler) TotalUnreadNotification(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		return errorUtils.ErrUnauthenticated
	}

	req := dto.TotalUnreadNotificationRequest{
		UserID: userId,
	}

	data, err := n.notifyService.TotalUnreadNotification(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	resp := responses.DefaultSuccess
	resp.Data = data

	return resp.JSON(ctx)

}

func (n notifyHandler) SendCampaignNotification(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.SendCampaignNotificationRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.SendCampaignNotification(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = data

	return response.JSON(ctx)
}

func (n notifyHandler) SendNotification(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.SendNotificationRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.SendNotification(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = data

	return response.JSON(ctx)
}

func (n notifyHandler) MarkAsRead(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.MarkAsReadRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.MarkAsRead(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = data

	return response.JSON(ctx)
}

func (n notifyHandler) ClearAllNotification(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		return errorUtils.ErrUnauthenticated
	}

	req := dto.ClearNotificationRequest{UserID: userId}

	data, err := n.notifyService.ClearAllNotification(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = data

	return response.JSON(ctx)
}
