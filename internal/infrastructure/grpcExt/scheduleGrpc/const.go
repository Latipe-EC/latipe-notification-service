package scheduleGrpc

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
)

const (
	ONETIME    = "onetime"
	EVERYDAY   = "every-day"
	EVERYWEEK  = "every-week"
	EVERYMONTH = "every-month"
	EVERYYEAR  = "every-year"
)

// ConvertStructToJSONString Convert the struct to a JSON string
func ConvertStructToJSONString(data interface{}) string {

	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error marshaling struct to JSON:", err)
		return ""
	}
	return string(jsonString)
}
