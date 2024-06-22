package notifyHookService

import (
	"context"
	dto "latipe-notification-service/internal/domain/dto/schedule_callback_dto"
)

type NotifyHookService interface {
	// commands
	CallBackFromScheduleService(ctx context.Context, req *dto.ScheduleCallbackRequest) error
}
