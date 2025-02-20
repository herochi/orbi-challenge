package interfaces

import (
	"github/herochi/orbi/service-a/application/user/ports/viewmodel"
	"github/herochi/orbi/service-a/domain"
)

//go:generate mockgen -destination=../mock/mock_UserPresenter.go -package=UserPresenter_mock . UserPresenter
type UserPresenter interface {
	UserResponse(d *viewmodel.UserVM) *domain.User
	UserVMResponse(d *domain.User) *viewmodel.UserVM
	UpdateUserToDom(u *viewmodel.UpdateUser) *domain.UpdateUser
}
