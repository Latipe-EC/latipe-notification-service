package notifyRepos

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/pkgUtils/db/mongodb"
	"latipe-notification-service/pkgUtils/util/pagable"
)

type notificationRepository struct {
	_notiCol *mongo.Collection
}

func NewNotificationRepository(dbClient *mongodb.MongoClient) NotificationRepository {
	col := dbClient.GetDB().Collection("user_notification")
	return &notificationRepository{_notiCol: col}
}

func (n notificationRepository) FindByID(ctx context.Context, entityID string) (*notication.Notification, error) {
	var entity notication.Notification

	err := n._notiCol.FindOne(ctx, bson.M{"_id": entityID}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (n notificationRepository) FindByOwnerID(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, error) {
	var entities []*notication.Notification

	opts := options.Find().SetLimit(int64(query.GetLimit())).SetSkip(int64(query.GetOffset()))
	filter := bson.M{"owner_id": OwnerID}

	cursor, err := n._notiCol.Find(ctx, filter, opts)
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

func (n notificationRepository) FindUnreadMessageOfUser(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, error) {
	var entities []*notication.Notification

	opts := options.Find().SetLimit(int64(query.GetLimit())).SetSkip(int64(query.GetOffset()))
	filter := bson.M{"owner_id": OwnerID, "unread": false}

	cursor, err := n._notiCol.Find(ctx, filter, opts)
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

func (n notificationRepository) TotalUnreadMessageOfUser(ctx context.Context, OwnerID string) (int64, error) {
	filter := bson.M{"owner_id": OwnerID, "unread": false}
	total, err := n._notiCol.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (n notificationRepository) Save(ctx context.Context, entity *notication.Notification) (*notication.Notification, error) {
	_, err := n._notiCol.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (n notificationRepository) Update(ctx context.Context, entity *notication.Notification) error {
	_, err := n._notiCol.UpdateOne(ctx, bson.M{"_id": entity.ID}, entity)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (n notificationRepository) Delete(ctx context.Context, entityId string) error {
	_, err := n._notiCol.DeleteOne(ctx, bson.M{"_id": entityId})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
