package notifyRepos

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	//TODO implement me
	panic("implement me")
}

func (n notificationRepository) FindByOwnerID(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationRepository) FindUnreadMessageOfUser(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationRepository) TotalUnreadMessageOfUser(ctx context.Context, OwnerID string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationRepository) Save(ctx context.Context, entity *notication.Notification) (*notication.Notification, error) {
	//TODO implement me
	panic("implement me")
}

func (n notificationRepository) Update(ctx context.Context, entity *notication.Notification) error {
	//TODO implement me
	panic("implement me")
}

func (n notificationRepository) Delete(ctx context.Context, entityId string) error {
	//TODO implement me
	panic("implement me")
}
