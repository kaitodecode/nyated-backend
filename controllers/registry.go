package controllers

import "github.com/kaitodecode/nyated-backend/services"

type ControllerRegistry struct {
	services services.IServiceRegistery
}

// UserController implements IControllerRegistry.
func (c *ControllerRegistry) UserController() IUserController {
	return NewUserController(c.services)
}

type IControllerRegistry interface {
	UserController() IUserController
}

func NewControllerRegistry(services services.IServiceRegistery) IControllerRegistry {
	return &ControllerRegistry{services}
}
