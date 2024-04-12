package dto

type TotalUnreadNotificationRequest struct {
	UserID string `json:"user_id"`
}

type TotalUnreadNotificationResponse struct {
	Total int `json:"total"`
}
