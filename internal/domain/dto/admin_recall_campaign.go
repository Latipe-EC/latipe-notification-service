package dto

type RecallNotificationRequest struct {
	NotificationID string `json:"notification_id" validate:"required"`
	Reason         string `json:"reason" validate:"required"`
}

type RecallNotificationResponse struct {
}
