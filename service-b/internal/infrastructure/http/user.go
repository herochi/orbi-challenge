package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/herochi/orbi/service-b/internal/domain"
)

type UserClient struct {
	client *resty.Client
}

func NewUserClient() *UserClient {
	return &UserClient{client: resty.New()}
}

func (uc *UserClient) GetUser(userID string) (*domain.User, error) {
	resp, err := uc.client.R().
		SetHeader("Accept", "application/json").
		Get(fmt.Sprintf("http://service-a:8080/orbi-api/v1/a/users/%s", userID))
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}

	var user domain.User
	err = json.Unmarshal(resp.Body(), &user)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return &user, nil
}
