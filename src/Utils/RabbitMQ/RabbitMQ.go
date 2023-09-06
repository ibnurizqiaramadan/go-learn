package RabbitMQ

import (
	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

// var Conn, err = amqp091.Dial("amqp://username:password@ip:port/")

var Connection *amqp.Connection

func InitRabbitMQ() *amqp.Connection {
	Connection, err := amqp.Dial("amqp://admin:JA5txh4j2u2iwHX2@194.233.95.186:15674/")
	if err != nil {
		log.Error("RabbitMQ error: ", err)
	}
	return Connection
}
