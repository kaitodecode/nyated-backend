package seeder

import "gorm.io/gorm"

type SeederRegistry struct {
	db *gorm.DB
}

// Run implements ISeederRegistry.
func (s *SeederRegistry) Run() {
	RunRoleSeeder(s.db)
	RunUserSeeder(s.db)
}

type ISeederRegistry interface {
	Run()
}

func NewSeederRegistry(db *gorm.DB) ISeederRegistry {
	return &SeederRegistry{db}
}
