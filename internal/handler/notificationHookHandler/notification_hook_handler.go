package notificationHookHandler

import "github.com/gofiber/fiber/v2"

type NotificationHookHandler interface {
	HandleScheduleCallback(ctx *fiber.Ctx) error
}
