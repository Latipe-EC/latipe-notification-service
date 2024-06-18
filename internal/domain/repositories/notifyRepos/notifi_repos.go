package notifyRepos

import (
	"context"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/pkgUtils/util/pagable"
)

type NotificationRepository interface {
	// queries
	FindByID(ctx context.Context, entityID string) (*notication.Notification, error)
	FindByOwnerID(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, int, error)
	FindUnreadMessageOfUser(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, int, error)
	TotalUnreadMessageOfUser(ctx context.Context, OwnerID string) (int64, error)
	FindAllCampaigns(ctx context.Context, query *pagable.Query) ([]*notication.Notification, int, error)

	// commands
	Save(ctx context.Context, entity *notication.Notification) (*notication.Notification, error)
	UpdateReadStatusNotification(ctx context.Context, entity *notication.Notification) error
	UpdateAllReadMessageOfUser(ctx context.Context, OwnerID string) error
	Delete(ctx context.Context, entityId string) error
	DeleteManyNotificationOfUser(ctx context.Context, userId string) error
	RecallCampaign(ctx context.Context, entity *notication.Notification) error
}
