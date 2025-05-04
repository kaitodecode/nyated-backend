package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/controllers"
)

type RoutesRegistry struct {
	controllers controllers.IControllerRegistry
	group       *gin.RouterGroup
}

// userRoutes implements IRoutesRegistry.
func (r *RoutesRegistry) userRoutes() IUserRoutes{
	return NewUserRoutes(r.controllers, r.group)
}

type IRoutesRegistry interface {
	Serve()
	userRoutes() IUserRoutes
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRoutesRegistry {
	return &RoutesRegistry{
		controllers: controller,
		group:       group,
	}
}

// Serve implements IRoutesRegistry.
func (r *RoutesRegistry) Serve() {
	r.userRoutes().Run()
}
