package notifyRouter

import (
	"github.com/gofiber/fiber/v2"
	"latipe-notification-service/internal/handler/notificationHookHandler"
	"latipe-notification-service/internal/handler/notifyHandler"
	"latipe-notification-service/internal/middleware"
)

type NotificationRouter interface {
	Init(root *fiber.Router)
}

type notificationRouter struct {
	notifyHandler   notifyHandler.NotifyHandler
	notiHookHandler notificationHookHandler.NotificationHookHandler
	authMiddleware  *middleware.AuthMiddleware
}

func NewNotificationRouter(notifyHandler notifyHandler.NotifyHandler,
	notiHookHandler notificationHookHandler.NotificationHookHandler,
	authMiddleware *middleware.AuthMiddleware) NotificationRouter {
	return &notificationRouter{
		notifyHandler:   notifyHandler,
		notiHookHandler: notiHookHandler,
		authMiddleware:  authMiddleware,
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

	//admin services
	admin := notify.Group("/admin", n.authMiddleware.RequiredRoles([]string{middleware.ADMIN_ROLE}))
	admin.Post("/notify-campaign", n.notifyHandler.AdminCreateCampaign)
	admin.Delete("/notify-campaign", n.notifyHandler.AdminRecallCampaign)
	admin.Get("/notify-campaign", n.notifyHandler.AdminGetAllCampaigns)

	//notification hook
	notifyHook := notify.Group("/schedule", n.authMiddleware.RequiredAPIKeyHeader())
	notifyHook.Post("/callback", n.notiHookHandler.HandleScheduleCallback)
}
