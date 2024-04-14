package notifyRepos

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/pkgUtils/db/mongodb"
	"latipe-notification-service/pkgUtils/util/pagable"
	"time"
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

	id, _ := primitive.ObjectIDFromHex(entityID)

	err := n._notiCol.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (n notificationRepository) FindByOwnerID(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, int, error) {
	var entities []*notication.Notification

	opts := options.Find().SetLimit(int64(query.GetLimit())).SetSkip(int64(query.GetOffset()))
	filter := bson.M{"owner_id": OwnerID}

	cursor, err := n._notiCol.Find(ctx, filter, opts)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	if err = cursor.All(ctx, &entities); err != nil {
		log.Error(err)
		return nil, 0, err
	}
	total, err := n._notiCol.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return entities, int(total), nil
}

func (n notificationRepository) FindUnreadMessageOfUser(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, int, error) {
	var entities []*notication.Notification

	opts := options.Find().SetLimit(int64(query.GetLimit())).SetSkip(int64(query.GetOffset()))
	filter := bson.M{"owner_id": bson.M{"$in": []string{OwnerID, "all"}}, "unread": true}

	cursor, err := n._notiCol.Find(ctx, filter, opts)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	if err = cursor.All(ctx, &entities); err != nil {
		log.Error(err)
		return nil, 0, err
	}

	total, err := n._notiCol.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return entities, int(total), nil
}

func (n notificationRepository) TotalUnreadMessageOfUser(ctx context.Context, OwnerID string) (int64, error) {
	filter := bson.M{"owner_id": OwnerID, "unread": true}
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
	filter := bson.D{{"_id", entity.ID}}

	update := bson.D{
		{"$set", bson.D{
			{"unread", false},
			{"update_at", time.Now()},
		}},
	}
	_, err := n._notiCol.UpdateOne(ctx, filter, update)
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

func (n notificationRepository) DeleteManyNotificationOfUser(ctx context.Context, userId string) error {
	_, err := n._notiCol.DeleteMany(ctx, bson.M{"owner_id": userId})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (n notificationRepository) UpdateAllReadMessageOfUser(ctx context.Context, OwnerID string) error {
	_, err := n._notiCol.UpdateMany(ctx, bson.M{"owner_id": OwnerID, "unread": true}, bson.M{"$set": bson.M{"unread": false}})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
