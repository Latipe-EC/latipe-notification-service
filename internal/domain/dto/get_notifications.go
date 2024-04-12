package dto

import "latipe-notification-service/pkgUtils/util/pagable"

type GetNotificationsRequest struct {
	UserID string `json:"user_id"`
	Query  *pagable.Query
}

type GetNotificationsResponse struct {
	pagable.ListResponse
}
