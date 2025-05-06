package controllers

import "github.com/kaitodecode/nyated-backend/services"

type ControllerRegistry struct {
	services services.IServiceRegistery
}

// FolderController implements IControllerRegistry.
func (c *ControllerRegistry) FolderController() IFolderController {
	return NewFolderController(c.services)
}

// UserController implements IControllerRegistry.
func (c *ControllerRegistry) UserController() IUserController {
	return NewUserController(c.services)
}

func (c *ControllerRegistry) NoteController() INoteController {
	return NewNoteController(c.services)
}

type IControllerRegistry interface {
	UserController() IUserController
	FolderController() IFolderController
	NoteController() INoteController
}

func NewControllerRegistry(services services.IServiceRegistery) IControllerRegistry {
	return &ControllerRegistry{services}
}
