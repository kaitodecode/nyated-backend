package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/common/lib/validation"
	"github.com/kaitodecode/nyated-backend/common/response"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/services"
)

type UserController struct {
	services services.IServiceRegistery
}

// Login implements IUserController.
func (u *UserController) Login(c *gin.Context) {
	request := &dto.UserLoginRequest{}
	if res, err := validation.ValidateBodyJson(c, request); err != nil {
		response.HttpResponse(res)
		return
	}

	 res, err := u.services.UserService().Login(c,request)
	 if err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}

	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Data: res,
		Gin: c,
	})
}

// Me implements IUserController.
func (u *UserController) Me(c *gin.Context) {
	res, err := u.services.UserService().GetUserLogin(c.Request.Context())
	if err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}
	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Data: res,
		Gin: c,
	})
}

// Register implements IUserController.
func (u *UserController) Register(c *gin.Context) {
	request := &dto.UserRegisterRequest{}
	if res, err := validation.ValidateBodyJson(c, request); err != nil {
		response.HttpResponse(res)
		return
	}

	err := u.services.UserService().Register(c, request)
	 if err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}

	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Gin: c,
	})
}

type IUserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Me(*gin.Context)
}

func NewUserController(services services.IServiceRegistery) IUserController {
	return &UserController{services: services}
}
