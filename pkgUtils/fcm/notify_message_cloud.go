package fcm

import (
	"context"
	"firebase.google.com/go/messaging"
)

type NotificationCloudMessage interface {
	SendToSingleDevice(ctx context.Context, message *messaging.Message) error
	SendToMultipleDevices(ctx context.Context, message *messaging.MulticastMessage) error
	SendToTopic(ctx context.Context, message *messaging.Message) error
	SendToCondition(ctx context.Context, message *messaging.Message) error
	SendToDeviceGroup(ctx context.Context, message *messaging.Message) error
	SubscribeToTopic(ctx context.Context, registrationTokens []string, topic string) error
	UnsubscribeFromTopic(ctx context.Context, registrationTokens []string, topic string) error
}
