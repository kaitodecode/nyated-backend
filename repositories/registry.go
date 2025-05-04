package repositories

import (
	"gorm.io/gorm"
)

type RepositoryRegistry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface{
	UserRepository() IUserRepository
	RoleRepository() IRoleRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry{
	return &RepositoryRegistry{
		db: db,
	}
}

func (r *RepositoryRegistry) UserRepository() IUserRepository {
	return NewUserRepository(r.db)
}

func (r *RepositoryRegistry) RoleRepository() IRoleRepository {
	return NewRoleRepository(r.db)
}