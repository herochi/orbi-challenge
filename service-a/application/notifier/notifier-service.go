package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github/herochi/orbi/service-a/application/user/ports/viewmodel"
	"log"
)

const serviceB = "service-b"

//go:generate mockgen -destination=../mock/mock_Notifier.go -package=Notifier_mock . Notifier
type Notifier interface {
	NotifyServiceB(data *viewmodel.UserVM)
}
type Service struct {
	rabbitMQChannel *amqp091.Channel
	channelR        *string
}

func NewNotifier(rabbitMQChannel *amqp091.Channel, channelR *string) Notifier {
	return &Service{
		rabbitMQChannel: rabbitMQChannel,
		channelR:        channelR,
	}
}

func (s *Service) NotifyServiceB(data *viewmodel.UserVM) {
	bodyMessage, err := json.Marshal(&viewmodel.Message{
		To:          serviceB,
		UserUpdated: *data,
	})
	if err != nil {
		log.Println(fmt.Sprintf("Error marshalling message: %s", err))
	}

	err = s.rabbitMQChannel.Publish(
		"",
		*s.channelR,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        bodyMessage,
		},
	)
	if err != nil {
		log.Println(fmt.Sprintf("Error when publishing message: %s", err))
	}

}
