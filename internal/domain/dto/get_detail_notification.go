package dto

type GetNotificationDetailRequest struct {
	ID string `params:"id" json:"id"`
}

type NotificationDetailResponse struct {
	ID        string `json:"id"`
	OwnerID   string `json:"owner_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Message   string `json:"message,omitempty"`
	IsRead    bool   `json:"is_read,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type GetNotificationDetailResponse struct {
	NotificationDetailResponse
}
