package RabbitMQ

import (
	"log"

	"go-learning/src/Utils/RabbitMQ"

	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
)

func SendRebbitMQ(c *fiber.Ctx) error {
	conn := RabbitMQ.InitRabbitMQ()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		true,    // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		log.Fatal(err)
	}

	body := "Hello World!"
	err = ch.PublishWithContext(
		c.Context(),
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	return c.JSON(fiber.Map{
		"statusCode": 200,
		"data": fiber.Map{
			"valid":    true,
			"messages": "success-send-hello-world",
		},
	})
}
