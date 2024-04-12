package notication

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner           primitive.ObjectID `bson:"owner,omitempty" json:"owner,omitempty"`
	Title           string             `bson:"title" json:"title"`
	Image           string             `bson:"image,omitempty" json:"image,omitempty"`
	Body            string             `bson:"body" json:"body"`
	Type            int                `bson:"type" json:"type"`
	UnRead          bool               `bson:"unread" json:"unread"`
	ScheduleDisplay time.Time          `bson:"schedule_display" json:"schedule_display"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
