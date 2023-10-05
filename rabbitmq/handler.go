package rabbitmq

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Handler(msg amqp.Delivery) {
	var message AMQPMessage

	error := json.Unmarshal([]byte(msg.Body), &message)

	if error != nil {
		log.Panicf("Failed to parse message: %s", msg.Body)
	}

	log.Printf("id: %d, name: %s", message.Id, message.Name)
}
