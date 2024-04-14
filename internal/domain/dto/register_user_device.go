package dto

import "latipe-notification-service/internal/domain/entities/userDevice"

const (
	WEB_TYPE_REQ     = 1
	ANDROID_TYPE_REQ = 2
	IOS_TYPE_REQ     = 3
)

func (r *RegisterNewDevice) GetDeviceType() string {
	switch r.DeviceType {
	case WEB_TYPE_REQ:
		return userDevice.WEB_BROWSER
	case ANDROID_TYPE_REQ:
		return userDevice.ANDROID
	case IOS_TYPE_REQ:
		return userDevice.IOS
	default:
		return "anonymous"
	}
}

type RegisterNewDevice struct {
	UserID      string `validate:"required"`
	DeviceInfo  string `json:"device_info" validate:"required"`
	DeviceToken string `json:"device_token" validate:"required"`
	DeviceType  int    `json:"device_type" validate:"required"`
}

type RegisterNewDeviceResponse struct {
	ID string `json:"id"`
}
