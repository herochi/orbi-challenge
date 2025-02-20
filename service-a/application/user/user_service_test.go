package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	UserServiceClient_mock "github/herochi/orbi/service-a/adapter/grpc/mock"
	Notifier_mock "github/herochi/orbi/service-a/application/mock"
	UserRepository_mock "github/herochi/orbi/service-a/application/user/mock"
	"github/herochi/orbi/service-a/application/user/ports/viewmodel"
	"github/herochi/orbi/service-a/domain"
	"testing"
)

func setup(t *testing.T) (context.Context, *userService) {
	ctrl := gomock.NewController(t)
	repository := UserRepository_mock.NewMockUserRepository(ctrl)
	presenter := NewUserPresenter()
	grpcClient := UserServiceClient_mock.NewMockUserServiceClient(ctrl)
	notifier := Notifier_mock.NewMockNotifier(ctrl)

	return context.Background(), &userService{
		repository: repository,
		presenter:  presenter,
		grpcClient: grpcClient,
		notifier:   notifier,
	}
}

// Test name empty
func TestCreateUser01(t *testing.T) {
	ctx, userS := setup(t)
	_, err := userS.Create(ctx, &viewmodel.UserVM{})

	assert.Error(t, err)
	assert.Equal(t, "name is empty", err.Error())
}

// Test Age less than 0
func TestCreateUser02(t *testing.T) {
	ctx, userS := setup(t)
	_, err := userS.Create(ctx, &viewmodel.UserVM{
		Name: "test",
		Age:  -1,
	})

	assert.Error(t, err)
	assert.Equal(t, "age can't be less than zero", err.Error())
}

// Test Email empty
func TestCreateUser03(t *testing.T) {
	ctx, userS := setup(t)
	_, err := userS.Create(ctx, &viewmodel.UserVM{
		Name: "test",
		Age:  23,
	})

	assert.Error(t, err)
	assert.Equal(t, "email is empty", err.Error())
}

// Test Email wrong format
func TestCreateUser04(t *testing.T) {
	ctx, userS := setup(t)
	_, err := userS.Create(ctx, &viewmodel.UserVM{
		Name:  "test",
		Age:   23,
		Email: "notformat",
	})

	assert.Error(t, err)
	assert.Equal(t, "field email doesn't have email format", err.Error())
}

// Test Repository error
func TestCreateUser05(t *testing.T) {
	ctx, userS := setup(t)

	repository := userS.repository.(*UserRepository_mock.MockUserRepository)
	repository.EXPECT().Save(gomock.Any(), gomock.Any()).Return("", assert.AnError)

	_, err := userS.Create(ctx, &viewmodel.UserVM{
		Name:  "test",
		Age:   23,
		Email: "ricardo@orbi.com",
	})

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
}

// Test Ok
func TestCreateUser06(t *testing.T) {
	ctx, userS := setup(t)

	repository := userS.repository.(*UserRepository_mock.MockUserRepository)
	repository.EXPECT().Save(gomock.Any(), gomock.Any()).Return("123", nil)

	grpcClient := userS.grpcClient.(*UserServiceClient_mock.MockUserServiceClient)
	grpcClient.EXPECT().NotifyUser(gomock.Any(), gomock.Any()).Return(nil, nil)

	id, err := userS.Create(ctx, &viewmodel.UserVM{
		Name:  "test",
		Age:   23,
		Email: "ricardo@orbi.com",
	})

	assert.Nil(t, err)
	assert.Equal(t, "123", id)
}

// Test Email wrong format
func TestUpdateUser01(t *testing.T) {
	ctx, userS := setup(t)
	err := userS.UpdateUser(ctx, "123", &viewmodel.UpdateUser{
		Name:  "testing",
		Email: "notformat",
	})
	assert.Error(t, err)
	assert.Equal(t, "field email doesn't have email format", err.Error())
}

// Test Error updating data
func TestUpdateUser02(t *testing.T) {
	ctx, userS := setup(t)

	repository := userS.repository.(*UserRepository_mock.MockUserRepository)
	repository.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

	err := userS.UpdateUser(ctx, "123", &viewmodel.UpdateUser{
		Name:  "testing",
		Email: "updated@orbi.com",
	})

	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
}

// Test OK
func TestUpdateUser03(t *testing.T) {
	ctx, userS := setup(t)

	repository := userS.repository.(*UserRepository_mock.MockUserRepository)
	repository.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.User{
		Name:  "testing",
		Email: "updated@orbi.com",
	}, nil)

	notifier := userS.notifier.(*Notifier_mock.MockNotifier)
	notifier.EXPECT().NotifyServiceB(gomock.Any())

	err := userS.UpdateUser(ctx, "123", &viewmodel.UpdateUser{
		Name:  "testing",
		Email: "updated@orbi.com",
	})

	assert.Nil(t, err)
}
