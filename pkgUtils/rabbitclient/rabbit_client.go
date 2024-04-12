package rabbitclient

import (
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"latipe-notification-service/config"
)

func NewRabbitClientConnection(globalCfg *config.AppConfig) *amqp.Connection {
	cfg := amqp.Config{
		Properties: amqp.Table{
			"connection_name": globalCfg.RabbitMQ.ServiceName,
		},
	}

	conn, err := amqp.DialConfig(globalCfg.RabbitMQ.Connection, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ cause:%v", err)
	}

	log.Info("Comsumer has been connected")
	return conn
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
