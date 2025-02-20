package rest

import (
	"net/http"

	ucc "github/herochi/orbi/service-a/application/user"
	errovm "github/herochi/orbi/service-a/domain/view-model"

	"github/herochi/orbi/service-a/application/user/ports/viewmodel"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type UserControllerImpl struct {
	service ucc.UserService
}

func NewUserHandler(s ucc.UserService) UserController {
	return &UserControllerImpl{s}
}

func (c *UserControllerImpl) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := c.service.GetById(ctx, id)
	if user == nil || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserControllerImpl) Create(ctx *gin.Context) {
	errVMResp := errovm.ErrorMessage{}
	var user viewmodel.UserVM

	if err := ctx.ShouldBindJSON(&user); err != nil {
		errVMResp = createError(err)
		ctx.JSON(http.StatusBadRequest, &errVMResp)
		return
	}

	id, err := c.service.Create(ctx, &user)
	if err != nil {
		errVMResp = createError(err)
		ctx.JSON(http.StatusBadRequest, &errVMResp)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (c *UserControllerImpl) UpdateUser(ctx *gin.Context) {
	errVMResp := errovm.ErrorMessage{}

	var user viewmodel.UpdateUser

	if err := ctx.ShouldBindJSON(&user); err != nil {
		errVMResp = createError(err)
		ctx.JSON(http.StatusBadRequest, &errVMResp)
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := c.service.UpdateUser(ctx, id, &user)
	if err != nil {
		errVMResp = createError(err)
		ctx.JSON(http.StatusBadRequest, &errVMResp)
		return
	}

	ctx.Status(http.StatusOK)
}

func createError(err error) errovm.ErrorMessage {
	errVMResp := errovm.ErrorMessage{}
	errVMResp.Code = 402
	errVMResp.Errors = append(errVMResp.Errors, err.Error())
	return errVMResp
}
