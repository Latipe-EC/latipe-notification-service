package dto

type AdminCreateCampaignRequest struct {
	UserID          string
	CampaignTopic   string `json:"campaign_topic" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Body            string `json:"body" validate:"required"`
	Image           string `json:"image" validate:"required"`
	ScheduleDisplay string `json:"schedule_display"`
}

type AdminCreateCampaignResponse struct {
	ID string `json:"id"`
}
