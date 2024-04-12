package service

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/service/notifyService"
)

var Set = wire.NewSet(notifyService.NewNotificationService)
