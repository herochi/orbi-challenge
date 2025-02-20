package interfaces

import (
	"context"
	"github/herochi/orbi/service-a/domain"
)

//go:generate mockgen -destination=../mock/mock_UserRepository.go -package=UserRepository_mock . UserRepository
type UserRepository interface {
	Save(ctx context.Context, user *domain.User) (string, error)
	GetById(ctx context.Context, id string) (*domain.User, error)
	UpdateUser(ctx context.Context, userID string, dataToUpdate *domain.UpdateUser) (*domain.User, error)
}
