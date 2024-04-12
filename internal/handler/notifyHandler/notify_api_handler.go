package notifyHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewNotifyHandler)

type NotifyHandler interface {
	// queries
	GetNotificationsOfUser(ctx *fiber.Ctx) error
	GetNotificationDetail(ctx *fiber.Ctx) error
	TotalUnreadNotification(ctx *fiber.Ctx) error
	SendCampaignNotification(ctx *fiber.Ctx) error
	// commands
	SendNotification(ctx *fiber.Ctx) error
	MarkAsRead(ctx *fiber.Ctx) error
	ClearAllNotification(ctx *fiber.Ctx) error
}
