package userDeviceRepos

import (
	"context"
	"latipe-notification-service/internal/domain/entities/userDevice"
)

type UserDeviceRepository interface {
	// queries
	FindByID(ctx context.Context, entityID string) (*userDevice.UserDevice, error)
	FindByDeviceToken(ctx context.Context, deviceID string) (*userDevice.UserDevice, error)
	FindActiveDeviceByUserID(ctx context.Context, userID string) ([]*userDevice.UserDevice, error)
	GetAllActiveDeviceToken(ctx context.Context) ([]string, error)
	// commands
	Save(ctx context.Context, entity *userDevice.UserDevice) (*userDevice.UserDevice, error)
	Update(ctx context.Context, entity *userDevice.UserDevice) error
}
