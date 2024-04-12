package dto

type MarkAsReadRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type MarkAsReadResponse struct {
}
