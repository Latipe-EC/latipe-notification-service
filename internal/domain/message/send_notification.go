package message

type SendNotificationMessage struct {
	UserID          string `json:"user_id" validate:"required"`
	Type            int    `json:"type" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Body            string `json:"body" validate:"required"`
	PushToDevice    bool   `json:"push_to_device"`
	Image           string `json:"image"`
	ScheduleDisplay string `json:"schedule_display"`
}
