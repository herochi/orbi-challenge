package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func StartConsumer(conn *amqp091.Connection, queueName string) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		true, false, false, false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)
		}
	}()
}
