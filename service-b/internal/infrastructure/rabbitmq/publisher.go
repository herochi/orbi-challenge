package rabbitmq

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

const QueueName = "service-b-notifies"

type Publisher struct {
	channel *amqp091.Channel
}

func NewPublisher(conn *amqp091.Connection) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	_, err = ch.QueueDeclare(
		QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	return &Publisher{channel: ch}, nil
}

func (p *Publisher) Publish(body []byte, queueName string) error {
	log.Printf("Publishing message: %s, to %s", body, queueName)

	err := p.channel.Publish(
		"",
		QueueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
