package services

import "github.com/kaitodecode/nyated-backend/repositories"

type ServiceRegistry struct {
	repositories repositories.IRepositoryRegistry
}

// UserService implements IServiceRegistery.
func (s *ServiceRegistry) UserService() IUserService {
	return NewUserService(s.repositories)
}

type IServiceRegistery interface {
	UserService() IUserService
}

func NewServiceRegistry(repositories repositories.IRepositoryRegistry) IServiceRegistery {
	return &ServiceRegistry{repositories}
}
