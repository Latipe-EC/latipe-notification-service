package schedule_callback_dto

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2/log"
)

type ScheduleCallbackRequest struct {
	Data     string `json:"data" validate:"required"`
	MetaData string `json:"meta_data"`
}

func (s ScheduleCallbackRequest) ParseDataToStruct() *ScheduleJsonData {
	var data ScheduleJsonData
	err := sonic.Unmarshal([]byte(s.Data), &data)
	if err != nil {
		log.Errorf("Error unmarshaling JSON to struct: %s", err)
		return nil
	}
	return &data

}

type ScheduleCallbackResponse struct {
}

type ScheduleJsonData struct {
	NotificationID string `json:"notification_id"`
}
