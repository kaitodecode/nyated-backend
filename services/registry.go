package services

import "github.com/kaitodecode/nyated-backend/repositories"

type ServiceRegistry struct {
	repositories repositories.IRepositoryRegistry
}

// FolderService implements IServiceRegistery.
func (s *ServiceRegistry) FolderService() IFolderService {
	return NewFolderService(s.repositories)
}

// UserService implements IServiceRegistery.
func (s *ServiceRegistry) UserService() IUserService {
	return NewUserService(s.repositories)
}

func (s *ServiceRegistry) NoteService() INoteService {
	return NewNoteService(s.repositories)
}

type IServiceRegistery interface {
	UserService() IUserService
	FolderService() IFolderService
	NoteService() INoteService
}

func NewServiceRegistry(repositories repositories.IRepositoryRegistry) IServiceRegistery {
	return &ServiceRegistry{repositories}
}
