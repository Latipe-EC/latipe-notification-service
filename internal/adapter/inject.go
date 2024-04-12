package adapter

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/adapter/authserv"
)

var Set = wire.NewSet(authserv.NewAuthService)
