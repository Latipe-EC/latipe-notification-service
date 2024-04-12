package userDeviceRepos

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"latipe-notification-service/internal/domain/entities/userDevice"
	"latipe-notification-service/pkgUtils/db/mongodb"
)

type userDeviceRepository struct {
	_deviceCol *mongo.Collection
}

func NewUserDeviceRepository(dbClient *mongodb.MongoClient) UserDeviceRepository {
	col := dbClient.GetDB().Collection("user_device")
	return &userDeviceRepository{_deviceCol: col}
}

func (u userDeviceRepository) FindByID(ctx context.Context, entityID string) (*userDevice.UserDevice, error) {
	//TODO implement me
	panic("implement me")
}

func (u userDeviceRepository) FindByDeviceID(ctx context.Context, deviceID string) (*userDevice.UserDevice, error) {
	//TODO implement me
	panic("implement me")
}

func (u userDeviceRepository) FindActiveDeviceByUserID(ctx context.Context, userID string) ([]*userDevice.UserDevice, error) {
	//TODO implement me
	panic("implement me")
}

func (u userDeviceRepository) Save(ctx context.Context, entity *userDevice.UserDevice) (*userDevice.UserDevice, error) {
	//TODO implement me
	panic("implement me")
}

func (u userDeviceRepository) Update(ctx context.Context, entity *userDevice.UserDevice) error {
	//TODO implement me
	panic("implement me")
}
