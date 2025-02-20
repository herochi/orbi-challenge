package user

import (
	"context"
	"fmt"
	gp "github/herochi/orbi/service-a/adapter/grpc/user"
	"github/herochi/orbi/service-a/application/notifier"
	"github/herochi/orbi/service-a/application/user/interfaces"
	"github/herochi/orbi/service-a/application/user/ports/viewmodel"
	"log"
)

type UserService interface {
	Create(ctx context.Context, user *viewmodel.UserVM) (string, error)
	GetById(ctx context.Context, id string) (*viewmodel.UserVM, error)
	UpdateUser(ctx context.Context, userID string, userVMUpdate *viewmodel.UpdateUser) error
}

type userService struct {
	repository interfaces.UserRepository
	presenter  interfaces.UserPresenter
	grpcClient gp.UserServiceClient
	notifier   notifier.Notifier
}

func NewUserService(r interfaces.UserRepository, p interfaces.UserPresenter, grpcClient gp.UserServiceClient,
	notifier notifier.Notifier) UserService {
	return &userService{repository: r, presenter: p, grpcClient: grpcClient, notifier: notifier}
}

func (u *userService) Create(ctx context.Context, userVM *viewmodel.UserVM) (string, error) {
	user := u.presenter.UserResponse(userVM)

	if err := user.Validate(); err != nil {
		return "", err
	}

	id, err := u.repository.Save(ctx, user)
	if err != nil {
		return "", err
	}

	u.NotifyInServiceB(ctx, id)

	return id, nil
}

func (u *userService) GetById(ctx context.Context, id string) (*viewmodel.UserVM, error) {
	user, err := u.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return u.presenter.UserVMResponse(user), nil
}

func (u *userService) UpdateUser(ctx context.Context, userID string, userVMUpdate *viewmodel.UpdateUser) error {
	userToUpdate := u.presenter.UpdateUserToDom(userVMUpdate)

	if err := userToUpdate.Validate(); err != nil {
		return err
	}

	dataUser, err := u.repository.UpdateUser(ctx, userID, u.presenter.UpdateUserToDom(userVMUpdate))
	if err != nil {
		log.Println(err)
		return err
	}

	u.notifier.NotifyServiceB(u.presenter.UserVMResponse(dataUser))

	return nil
}

func (u *userService) NotifyInServiceB(ctx context.Context, id string) {
	resp, err := u.grpcClient.NotifyUser(ctx, &gp.Request{UserId: id})
	if err != nil {
		log.Println(fmt.Errorf("error from service-b: %w", err))
		return
	}

	if resp != nil {
		log.Println(fmt.Sprintf("message response from service-b: %s \n", resp.Message))
	}

	return
}
