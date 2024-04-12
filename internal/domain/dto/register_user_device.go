package dto

type RegisterNewDevice struct {
	UserID      string `json:"user_id" validate:"required"`
	DeviceID    string `json:"device_id" validate:"required"`
	DeviceToken string `json:"device_token" validate:"required"`
	DeviceType  string `json:"device_type" validate:"required"`
}

type RegisterNewDeviceResponse struct {
	ID string `json:"id"`
}
