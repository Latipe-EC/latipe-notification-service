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
	ScheduleDisplay string             `bson:"schedule_display" json:"schedule_display"`
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

func (n Notification) ParseScheduleToTime() (time.Time, error) {
	// Define the layout according to the given date format
	layout := "2006-01-02 15:04:05"

	// Get the local location
	location := time.Now().Location()

	// Parse the date string using the specified layout and location
	parsedDate, err := time.ParseInLocation(layout, n.ScheduleDisplay, location)
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
