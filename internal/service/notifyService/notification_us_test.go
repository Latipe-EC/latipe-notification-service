package notifyService_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/domain/entities/notication"
	"latipe-notification-service/internal/domain/repositories/notifyRepos"
	"latipe-notification-service/internal/domain/repositories/userDeviceRepos"
	"latipe-notification-service/internal/service/notifyService"
	"latipe-notification-service/pkgUtils/fcm"
	"latipe-notification-service/pkgUtils/util/pagable"
	"testing"
	"time"
)

func TestNotificationService_GetNotificationsOfUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	type fields struct {
		notificationRepo notifyRepos.NotificationRepository
		userDeviceRepo   userDeviceRepos.UserDeviceRepository
		fbCloudMessage   fcm.NotificationCloudMessage
	}

	type args struct {
		ctx context.Context
		req *dto.GetNotificationsRequest
	}

	dataResp := []*notication.Notification{
		{
			ID:     primitive.NewObjectIDFromTimestamp(time.Now()),
			Owner:  "123",
			Title:  "Xin chào mừng bạn đã đến với Latipe",
			Image:  "https://latipe.com.vn/assets/images/logo.png",
			Body:   "Xin chào mừng bạn đã đến với Latipe",
			Type:   1,
			UnRead: true,
		},
	}

	notifiRepo := notifyRepos.NewMockNotificationRepository(ctrl)
	notifiRepo.EXPECT().FindUnreadMessageOfUser(ctx, "2", &pagable.Query{
		Page: 0,
		Size: 1,
	}).Return(dataResp, 10, nil)

	wantResp := dto.GetNotificationsResponse{}
	wantResp.Total = 10
	wantResp.Items = dataResp
	wantResp.Size = 1
	wantResp.HasMore = true

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dto.GetNotificationsResponse
		wantErr bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				notificationRepo: notifiRepo,
				userDeviceRepo:   userDeviceRepos.NewMockUserDeviceRepository(ctrl),
				fbCloudMessage:   fcm.NewMockNotificationCloudMessage(ctrl),
			},
			args: args{
				ctx: ctx,
				req: &dto.GetNotificationsRequest{
					UserID: "2",
					Query: &pagable.Query{
						Page: 0,
						Size: 1,
					},
				},
			},
			want:    wantResp,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := notifyService.NewNotificationService(tt.fields.notificationRepo, tt.fields.userDeviceRepo, tt.fields.fbCloudMessage)
			got, err := n.GetNotificationsOfUser(tt.args.ctx, tt.args.req)

			assert.Nil(t, err)
			assert.Equal(t, got.ListResponse, tt.want.ListResponse)
		})
	}
}
