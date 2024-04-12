package router

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/router/notifyRouter"
)

var Set = wire.NewSet(notifyRouter.NewNotificationRouter)
