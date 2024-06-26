package notifyHandler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/middleware"
	"latipe-notification-service/internal/service/notifyService"
	"latipe-notification-service/pkgUtils/util/errorUtils"
	"latipe-notification-service/pkgUtils/util/pagable"
	responses "latipe-notification-service/pkgUtils/util/response"
	"latipe-notification-service/pkgUtils/util/valid"
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
		switch {
		case errors.Is(err, mongo.ErrNoDocuments):
			return errorUtils.ErrNotFound
		}
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

	if err := valid.GetValidator().Validate(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.SendCampaignInternalService(context, &req)
	if err != nil {
		switch {
		case errors.Is(err, errorUtils.ErrParseDatetimeParameters):
			return errorUtils.ErrParseDatetimeParameters
		case errors.Is(err, errorUtils.ErrInvalidParameters):
			return errorUtils.ErrInvalidParameters
		case errors.Is(err, errorUtils.ErrInvalidDatetimeParameters):
			return errorUtils.ErrInvalidDatetimeParameters
		}

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

	if err := valid.GetValidator().Validate(&req); err != nil {
		return errorUtils.ErrInvalidParameters
	}

	_, err := n.notifyService.SendNotification(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess

	return response.JSON(ctx)
}

func (n notifyHandler) MarkAllRead(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userID := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userID == "" {
		return errorUtils.ErrUnauthenticated
	}

	req := dto.MarkAsReadRequest{userID}

	data, err := n.notifyService.MarkAllRead(context, &req)
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

func (n notifyHandler) RegisterNewUserDevice(ctx *fiber.Ctx) error {
	context := ctx.Context()

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		return errorUtils.ErrUnauthenticated
	}

	req := dto.RegisterNewDevice{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	req.UserID = userId

	if err := valid.GetValidator().Validate(&req); err != nil {
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.RegisterNewUserDevice(context, &req)
	if err != nil {
		switch {
		case errors.Is(err, errorUtils.ErrDeviceAlreadyRegistered):
			return errorUtils.ErrDeviceAlreadyRegistered
		}

		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = data

	return response.JSON(ctx)
}

func (n notifyHandler) AdminGetAllCampaigns(ctx *fiber.Ctx) error {
	context := ctx.Context()

	query, err := pagable.GetQueryFromFiberCtx(ctx)
	if err != nil {
		return errorUtils.ErrInvalidParameters
	}

	req := dto.AdminGetAllCampaignRequest{
		Query: query,
	}

	if err := valid.GetValidator().Validate(&req); err != nil {
		return errorUtils.ErrInvalidParameters
	}

	data, err := n.notifyService.AdminGetAllCampaigns(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = data

	return response.JSON(ctx)
}

func (n notifyHandler) AdminCreateCampaign(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.AdminCreateCampaignRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	userId := fmt.Sprintf("%s", ctx.Locals(middleware.USER_ID))
	if userId == "" {
		return errorUtils.ErrUnauthenticated
	}
	req.UserID = userId

	if err := valid.GetValidator().Validate(&req); err != nil {
		return errorUtils.ErrInvalidParameters
	}

	id, err := n.notifyService.AdminCreateCampaign(context, &req)
	if err != nil {
		switch {
		case errors.Is(err, errorUtils.ErrParseDatetimeParameters):
			return errorUtils.ErrParseDatetimeParameters
		case errors.Is(err, errorUtils.ErrInvalidParameters):
			return errorUtils.ErrInvalidParameters
		case errors.Is(err, errorUtils.ErrInvalidDatetimeParameters):
			return errorUtils.ErrInvalidDatetimeParameters
		}

		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess
	response.Data = id

	return response.JSON(ctx)
}

func (n notifyHandler) AdminRecallCampaign(ctx *fiber.Ctx) error {
	context := ctx.Context()

	req := dto.RecallNotificationRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		log.Error(err)
		return errorUtils.ErrInvalidParameters
	}

	if err := valid.GetValidator().Validate(&req); err != nil {
		return errorUtils.ErrInvalidParameters
	}

	_, err := n.notifyService.AdminRecallCampaign(context, &req)
	if err != nil {
		return errorUtils.ErrInternalServer
	}

	response := responses.DefaultSuccess

	return response.JSON(ctx)
}
