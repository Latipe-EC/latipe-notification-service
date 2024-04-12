package dto

type SendNotificationRequest struct {
	UserID          string `json:"user_id" validate:"required"`
	Type            int    `json:"type" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Body            string `json:"body" validate:"required"`
	Image           string `json:"image"`
	ScheduleDisplay string `json:"schedule_display"`
}

type SendNotificationResponse struct {
	ID string `json:"id"`
}
