package notifySubs

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"latipe-notification-service/config"
	"latipe-notification-service/internal/domain/dto"
	"latipe-notification-service/internal/service/notifyService"
	"sync"
	"time"
)

type NotifyToUserSubs struct {
	config     *config.AppConfig
	conn       *amqp.Connection
	notifyServ notifyService.NotificationService
}

func NewNotifyToUserSubs(cfg *config.AppConfig,
	conn *amqp.Connection,
	notifyServ notifyService.NotificationService) *NotifyToUserSubs {

	return &NotifyToUserSubs{
		config:     cfg,
		conn:       conn,
		notifyServ: notifyServ,
	}
}

func (n NotifyToUserSubs) ListenNotificationMessage(wg *sync.WaitGroup) error {
	channel, err := n.conn.Channel()
	defer func(channel *amqp.Channel) {
		err := channel.Close()
		if err != nil {

		}
	}(channel)

	if err != nil {
		return err
	}

	// define an exchange type "topic"
	err = channel.ExchangeDeclare(
		n.config.RabbitMQ.NotificationExchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("cannot declare exchange: %v", err)
		return err
	}

	// create queue
	q, err := channel.QueueDeclare(
		"notify_message",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("cannot declare queue: %v", err)
		return err
	}

	err = channel.QueueBind(
		q.Name,
		n.config.RabbitMQ.SendMessageRoutingKey,
		n.config.RabbitMQ.NotificationExchange,
		false,
		nil)
	if err != nil {
		log.Fatalf("cannot bind exchange: %v", err)
		return err
	}

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		q.Name,                        // queue
		n.config.RabbitMQ.ServiceName, // consumer
		true,                          // auto ack
		false,                         // exclusive
		false,                         // no local
		false,                         // no wait
		nil,                           //args
	)
	if err != nil {
		return err
	}

	defer wg.Done()
	// handle consumed messages from queue
	for msg := range msgs {
		log.Infof("received order message from: %s", msg.RoutingKey)
		if err := n.messageHandler(msg); err != nil {
			log.Infof("The order creation failed cause %s", err)
		}

	}

	log.Infof("message queue has started")
	log.Infof("waiting for messages...")

	return nil
}

func (n NotifyToUserSubs) messageHandler(msg amqp.Delivery) error {
	startTime := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message := dto.SendNotificationRequest{}

	if err := json.Unmarshal(msg.Body, &message); err != nil {
		log.Infof("Parse message to order failed cause: %s", err)
		return err
	}

	if _, err := n.notifyServ.SendNotification(ctx, &message); err != nil {
		log.Infof("Handling order message was failed cause: %s", err)
		return err
	}

	endTime := time.Now()
	log.Infof("The message was processed successfully - duration:%v", endTime.Sub(startTime))
	return nil
}
