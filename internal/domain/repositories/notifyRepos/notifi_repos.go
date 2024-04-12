package notifyRepos

import (
	"context"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/pkgUtils/util/pagable"
)

type NotificationRepository interface {
	// queries
	FindByID(ctx context.Context, entityID string) (*notication.Notification, error)
	FindByOwnerID(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, error)
	FindUnreadMessageOfUser(ctx context.Context, OwnerID string, query *pagable.Query) ([]*notication.Notification, error)
	TotalUnreadMessageOfUser(ctx context.Context, OwnerID string) (int64, error)
	// commands
	Save(ctx context.Context, entity *notication.Notification) (*notication.Notification, error)
	Update(ctx context.Context, entity *notication.Notification) error
	Delete(ctx context.Context, entityId string) error
}
