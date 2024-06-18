package userDeviceRepos

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var entity userDevice.UserDevice
	id, _ := primitive.ObjectIDFromHex(entityID)

	err := u._deviceCol.FindOne(ctx, map[string]interface{}{"_id": id}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (u userDeviceRepository) FindByDeviceID(ctx context.Context, deviceID string) (*userDevice.UserDevice, error) {
	var entity userDevice.UserDevice

	err := u._deviceCol.FindOne(ctx, bson.M{"device_id": deviceID}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (u userDeviceRepository) FindActiveDeviceByUserID(ctx context.Context, userID string) ([]*userDevice.UserDevice, error) {
	var entities []*userDevice.UserDevice

	cursor, err := u._deviceCol.Find(ctx, bson.M{"user_id": userID, "is_active": true})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err = cursor.All(ctx, &entities); err != nil {
		log.Error(err)
		return nil, err
	}

	return entities, nil
}

func (u userDeviceRepository) Save(ctx context.Context, entity *userDevice.UserDevice) (*userDevice.UserDevice, error) {
	_, err := u._deviceCol.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u userDeviceRepository) Update(ctx context.Context, entity *userDevice.UserDevice) error {
	_, err := u._deviceCol.UpdateOne(ctx, bson.M{"_id": entity.ID}, entity)
	if err != nil {
		return err
	}

	return nil
}

func (u userDeviceRepository) GetAllActiveDeviceToken(ctx context.Context) ([]string, error) {
	tokens, err := u._deviceCol.Distinct(ctx, "device_token", bson.M{"is_active": true})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var deviceTokens []string
	for _, token := range tokens {
		deviceTokens = append(deviceTokens, token.(string))
	}

	return deviceTokens, nil
}
