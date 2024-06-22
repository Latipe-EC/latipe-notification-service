package dto

import (
	"latipe-notification-service/pkgUtils/util/pagable"
)

type AdminGetAllCampaignRequest struct {
	Query *pagable.Query
}

type AdminGetAllCampaignResponse struct {
	pagable.ListResponse
}

type AdminGetCampaignDetailResponse struct {
	ID              string `json:"id"`
	CampaignTopic   string `json:"campaign_topic"`
	Title           string `json:"title"`
	Image           string `json:"image"`
	Body            string `json:"body"`
	IsActive        bool   `json:"is_active"`
	RecallReason    string `json:"recall_reason"`
	ScheduleDisplay string `json:"schedule_display"`
	Type            int    `json:"type"`
	CreatedBy       string `json:"created_by"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
