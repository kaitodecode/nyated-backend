package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/controllers"
	middlewares "github.com/kaitodecode/nyated-backend/middleware"
)

type FolderRoutes struct {
	controllers controllers.IControllerRegistry
	group       *gin.RouterGroup
}

// Run implements IFolderRoutes.
func (f *FolderRoutes) Run() {
	group := f.group.Group("/folder/").Use(middlewares.Authenticate())
	group.GET("/", f.controllers.FolderController().Index)
	group.GET("/:id/", f.controllers.FolderController().Show)
	group.POST("/", f.controllers.FolderController().Store)
	group.PUT("/:id", f.controllers.FolderController().Update)
	group.DELETE("/:id", f.controllers.FolderController().Destroy)
}

type IFolderRoutes interface {
	Run()
}

func NewFolderRoutes(controllers controllers.IControllerRegistry, group *gin.RouterGroup) IFolderRoutes {
	return &FolderRoutes{controllers: controllers, group: group}
}
