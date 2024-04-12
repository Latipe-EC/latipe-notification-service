package repositories

import (
	"github.com/google/wire"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
)

var Set = wire.NewSet(
	notifyRepos.NewNotificationRepository,
	userDeviceRepos.NewUserDeviceRepository,
)
