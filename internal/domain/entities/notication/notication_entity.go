package notication

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	OwnerID         string             `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	Title           string             `bson:"title" json:"title"`
	CampaignTopic   string             `bson:"campaign_topic,omitempty" json:"campaign_topic,omitempty"`
	Image           string             `bson:"image,omitempty" json:"image,omitempty"`
	Body            string             `bson:"body" json:"body"`
	Type            int                `bson:"type" json:"type"`
	UnRead          bool               `bson:"unread" json:"unread"`
	ScheduleDisplay time.Time          `bson:"schedule_display" json:"schedule_display"`
	IsActive        bool               `bson:"is_active" json:"is_active"`
	RecallReason    string             `bson:"recall_reason,omitempty" json:"recall_reason,omitempty"`
	CreatedBy       string             `bson:"created_by,omitempty" json:"created_by,omitempty"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}

func NewNotification() Notification {
	return Notification{
		Image:     "https://res.cloudinary.com/dddb8btv0/image/upload/f_auto,q_auto/v1/latipe/dtzjmllk215rggpznjar",
		UnRead:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
