package RabbitMQ

import (
	"github.com/rabbitmq/amqp091-go"
)

var Conn, err = amqp091.Dial("amqp://username:password@ip:port/")
