package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/controllers"
	middlewares "github.com/kaitodecode/nyated-backend/middleware"
)

type UserRoutes struct {
	controllers controllers.IControllerRegistry
	group *gin.RouterGroup
}

type IUserRoutes interface {
	Run()
}

func NewUserRoutes(controller controllers.IControllerRegistry, group *gin.RouterGroup) IUserRoutes {
	return &UserRoutes{controllers: controller, group: group}
}

func (r *UserRoutes) Run() {
	group := r.group.Group("/auth/")
	group.GET("/me/", middlewares.Authenticate(), r.controllers.UserController().Me)
	group.POST("/login/", r.controllers.UserController().Login)
	group.POST("/register/", r.controllers.UserController().Register)
}