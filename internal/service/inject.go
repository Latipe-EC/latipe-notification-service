package service

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/service/notifyHookService"
	"latipe-notification-service/internal/service/notifyService"
)

var Set = wire.NewSet(
	notifyService.NewNotificationService,
	notifyHookService.NewNotifyHookService,
)
