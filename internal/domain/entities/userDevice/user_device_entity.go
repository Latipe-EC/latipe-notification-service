package userDevice

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserDevice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      string             `bson:"user_id" json:"user_id"`
	DeviceInfo  string             `bson:"device_info" json:"device_info"`
	DeviceToken string             `bson:"device_token" json:"device_token"`
	DeviceType  string             `bson:"device_type" json:"device_type"`
	IsActive    bool               `bson:"is_active" json:"is_deleted"`
	LoggedDate  time.Time          `bson:"logged_date" json:"logged_date"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
