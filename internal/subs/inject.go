package subs

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/subs/notifySubs"
)

var Set = wire.NewSet(notifySubs.NewNotifyToUserSubs)
