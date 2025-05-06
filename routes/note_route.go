package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/controllers"
	middlewares "github.com/kaitodecode/nyated-backend/middleware"
)

type NoteRoutes struct {
	controllers controllers.IControllerRegistry
	group       *gin.RouterGroup
}

// Run implements INoteRoutes.
func (f *NoteRoutes) Run() {
	group := f.group.Group("/note/").Use(middlewares.Authenticate())
	group.GET("/", f.controllers.NoteController().Index)
	group.GET("/:id/", f.controllers.NoteController().Show)
	group.POST("/", f.controllers.NoteController().Store)
	group.PUT("/:id", f.controllers.NoteController().Update)
	group.DELETE("/:id", f.controllers.NoteController().Destroy)
}

type INoteRoutes interface {
	Run()
}

func NewNoteRoutes(controllers controllers.IControllerRegistry, group *gin.RouterGroup) INoteRoutes {
	return &NoteRoutes{controllers: controllers, group: group}
}
