package user

import (
	"fmt"

	"github/herochi/orbi/service-a/application/user/interfaces"
	"github/herochi/orbi/service-a/application/user/ports/viewmodel"
	"github/herochi/orbi/service-a/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userPresenter struct{}

func NewUserPresenter() interfaces.UserPresenter {
	return &userPresenter{}
}

func (*userPresenter) UserResponse(u *viewmodel.UserVM) *domain.User {
	fmt.Println(u)
	user := &domain.User{}
	ID, _ := primitive.ObjectIDFromHex(u.ID)
	user.ID = ID
	user.Name = u.Name
	user.Email = u.Email
	user.CreatedAt = u.CreatedAt
	user.Age = u.Age

	return user
}

func (*userPresenter) UserVMResponse(d *domain.User) *viewmodel.UserVM {

	userVM := viewmodel.UserVM{}
	userVM.ID = d.ID.Hex()
	userVM.Email = d.Email
	userVM.Name = d.Name
	userVM.CreatedAt = d.CreatedAt
	userVM.Age = d.Age

	return &userVM
}

func (*userPresenter) UpdateUserToDom(u *viewmodel.UpdateUser) *domain.UpdateUser {

	user := &domain.UpdateUser{}
	user.Name = u.Name
	user.Email = u.Email

	return user
}
