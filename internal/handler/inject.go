package handler

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/handler/notificationHookHandler"
	"latipe-notification-service/internal/handler/notifyHandler"
)

var Set = wire.NewSet(
	notifyHandler.NewNotifyHandler,
	notificationHookHandler.NewNotificationHookHandler,
)
