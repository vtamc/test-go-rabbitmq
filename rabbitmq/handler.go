package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Handler(msg amqp.Delivery) {
	log.Printf("Received a message: %s", msg.Body)
}
