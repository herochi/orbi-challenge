package application

import (
	"encoding/json"
	"fmt"
	"github.com/herochi/orbi/service-b/internal/domain"
	"github.com/herochi/orbi/service-b/internal/infrastructure/http"
	"github.com/herochi/orbi/service-b/internal/infrastructure/rabbitmq"
	"log"
)

type NotifyService struct {
	messageBroker rabbitmq.Publisher
	userClient    http.UserClient
}

func NewNotifyService(mb rabbitmq.Publisher, uc http.UserClient) *NotifyService {
	return &NotifyService{messageBroker: mb, userClient: uc}
}

func (s *NotifyService) NotifyUser(userID string) (string, error) {
	fmt.Printf("Received user ID: %s\n", userID)

	if userID == "" {
		return "", fmt.Errorf("id is empty")
	}

	user, err := s.userClient.GetUser(userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user data: %w", err)
	}

	log.Printf("user: %v", *user)

	bodyMessage, err := json.Marshal(domain.Message{
		To:      user.Email,
		Message: "hi this is our service b",
	})
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	err = s.messageBroker.Publish(bodyMessage, rabbitmq.QueueName)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %w", err)
	}

	return fmt.Sprintf("User %s has been notified, email sent to %s", user.Name, user.Email), nil
}
