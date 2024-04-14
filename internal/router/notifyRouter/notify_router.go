package notifyRouter

import (
	"github.com/gofiber/fiber/v2"
	"latipe-notification-service/internal/handler/notifyHandler"
	"latipe-notification-service/internal/middleware"
)

type NotificationRouter interface {
	Init(root *fiber.Router)
}

type notificationRouter struct {
	notifyHandler  notifyHandler.NotifyHandler
	authMiddleware *middleware.AuthMiddleware
}

func NewNotificationRouter(notifyHandler notifyHandler.NotifyHandler, authMiddleware *middleware.AuthMiddleware) NotificationRouter {
	return &notificationRouter{
		notifyHandler:  notifyHandler,
		authMiddleware: authMiddleware,
	}
}

func (n notificationRouter) Init(root *fiber.Router) {
	notify := (*root).Group("/notifications")

	user := notify.Group("/user", n.authMiddleware.RequiredAuthentication())
	user.Get("/total/unread", n.notifyHandler.TotalUnreadNotification)
	user.Get("", n.notifyHandler.GetNotificationsOfUser)
	user.Get("/:id", n.notifyHandler.GetNotificationDetail)
	user.Patch("/markAsRead", n.notifyHandler.MarkAllRead)
	user.Delete("", n.notifyHandler.ClearAllNotification)
	user.Post("/register-device", n.notifyHandler.RegisterNewUserDevice)

	//internal services
	internal := notify.Group("/internal", n.authMiddleware.RequiredAPIKeyHeader())
	internal.Post("/notify-message", n.notifyHandler.SendNotification)
	internal.Post("/notify-campaign", n.notifyHandler.SendCampaignNotification)

}
