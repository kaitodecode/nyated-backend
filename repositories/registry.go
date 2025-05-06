package repositories

import (
	"gorm.io/gorm"
)

type RepositoryRegistry struct {
	db *gorm.DB
}



type IRepositoryRegistry interface {
	UserRepository() IUserRepository
	RoleRepository() IRoleRepository
	FolderRepository() IFolderRepository
	NoteRepository() INoteRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &RepositoryRegistry{
		db: db,
	}
}

func (r *RepositoryRegistry) NoteRepository() INoteRepository {
	return NewNoteRepository(r.db)
}

func (r *RepositoryRegistry) UserRepository() IUserRepository {
	return NewUserRepository(r.db)
}

func (r *RepositoryRegistry) RoleRepository() IRoleRepository {
	return NewRoleRepository(r.db)
}

func (r *RepositoryRegistry) FolderRepository() IFolderRepository {
	return NewFolderRepository(r.db)
}