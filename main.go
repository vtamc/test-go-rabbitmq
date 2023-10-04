package main

import (
	"app/rabbitmq"
)

func main() {
	consumer := new(rabbitmq.Consumer)

	consumer.Start()
}
