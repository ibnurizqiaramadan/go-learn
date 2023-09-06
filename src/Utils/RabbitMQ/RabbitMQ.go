package RabbitMQ

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Connection, _ = amqp.Dial(os.Getenv("RABBITMQ_CONNECTION"))
