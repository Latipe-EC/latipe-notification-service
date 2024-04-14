package dto

type SendCampaignNotificationRequest struct {
	CampaignName    string `json:"campaign_name" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Body            string `json:"body" validate:"required"`
	Image           string `json:"image" validate:"required"`
	ScheduleDisplay string `json:"schedule_display"`
}

type SendCampaignNotificationResponse struct {
	ID string `json:"id"`
}
