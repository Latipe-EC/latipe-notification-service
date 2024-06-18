package dto

type GetNotificationDetailRequest struct {
	ID string `params:"id" json:"id"`
}

type NotificationDetailResponse struct {
	ID        string `json:"id"`
	OwnerID   string `json:"owner_id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	Body      string `json:"body"`
	Type      int    `json:"type"`
	UnRead    bool   `json:"unread"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetNotificationDetailResponse struct {
	NotificationDetailResponse
}
