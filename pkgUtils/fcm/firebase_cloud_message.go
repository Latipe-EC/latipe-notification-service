package fcm

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/api/option"
	"latipe-notification-service/config"

	"path/filepath"
	"time"
)

type FirebaseCloudMessage struct {
	_app    *firebase.App
	_client *messaging.Client
}

func NewFirebaseSDK(config *config.AppConfig) *FirebaseCloudMessage {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serviceAccountKeyFilePath, err := filepath.Abs(config.Firebase.ServiceAccountKeyFilePath)
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		panic("Firebase message service load error")
	}

	return &FirebaseCloudMessage{_app: app, _client: client}

}
func (fcm *FirebaseCloudMessage) SendToSingleDevice(ctx context.Context, message *messaging.Message) error {
	response, err := fcm._client.Send(ctx, message)
	if err != nil {
		log.Errorf("Error sending message: %s", err)
	}
	// Response is a message ID string.
	log.Infof("Successfully sent message: %v", response)
	return nil
}

func (fcm *FirebaseCloudMessage) SendToMultipleDevices(ctx context.Context, message *messaging.MulticastMessage) error {
	response, err := fcm._client.SendMulticast(ctx, message)
	if err != nil {
		log.Errorf("Error sending message: %s", err)
	}
	// Response is a message ID string.
	log.Infof("Successfully sent message: %s", response)
	return nil
}

func (fcm *FirebaseCloudMessage) SendToTopic(ctx context.Context, message *messaging.Message) error {
	response, err := fcm._client.Send(ctx, message)
	if err != nil {
		log.Errorf("Error sending message: %s", err)
	}
	// Response is a message ID string.
	log.Infof("Successfully sent message: %v", response)
	return nil
}

func (fcm *FirebaseCloudMessage) SendToCondition(ctx context.Context, message *messaging.Message) error {
	response, err := fcm._client.Send(ctx, message)
	if err != nil {
		log.Errorf("Error sending message: %s", err)
	}
	// Response is a message ID string.
	log.Infof("Successfully sent message: %v", response)
	return nil
}

func (fcm *FirebaseCloudMessage) SendToDeviceGroup(ctx context.Context, message *messaging.Message) error {
	response, err := fcm._client.Send(ctx, message)
	if err != nil {
		log.Errorf("Error sending message: %s", err)
	}
	// Response is a message ID string.
	log.Infof("Successfully sent message: %v", response)
	return nil
}

func (fcm *FirebaseCloudMessage) SubscribeToTopic(ctx context.Context, registrationTokens []string, topic string) error {
	if _, err := fcm._client.SubscribeToTopic(ctx, registrationTokens, topic); err != nil {
		log.Errorf("Error subscribing to topic: %s", err)
	}

	log.Infof("Subscribed to topic: %s", topic)
	return nil
}

func (fcm *FirebaseCloudMessage) UnsubscribeFromTopic(ctx context.Context, registrationTokens []string, topic string) error {
	if _, err := fcm._client.UnsubscribeFromTopic(ctx, registrationTokens, topic); err != nil {
		log.Errorf("Error unsubscribing from topic: %s", err)
	}

	log.Infof("Unsubscribed from topic: %s", topic)
	return nil
}
