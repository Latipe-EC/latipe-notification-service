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

	notify.Get("/total-unread", n.authMiddleware.RequiredAuthentication(), n.notifyHandler.TotalUnreadNotification)
	notify.Get("", n.authMiddleware.RequiredAuthentication(), n.notifyHandler.GetNotificationsOfUser)
	notify.Get("/:id", n.authMiddleware.RequiredAuthentication(), n.notifyHandler.GetNotificationDetail)

	//commands
	notify.Put("/:id/read", n.authMiddleware.RequiredAuthentication(), n.notifyHandler.MarkAsRead)
	notify.Delete("", n.authMiddleware.RequiredAuthentication(), n.notifyHandler.ClearAllNotification)
	notify.Post("", n.authMiddleware.RequiredAPIKeyHeader(), n.notifyHandler.SendNotification)
	notify.Post("/campaign", n.authMiddleware.RequiredAPIKeyHeader(), n.notifyHandler.SendCampaignNotification)

}
