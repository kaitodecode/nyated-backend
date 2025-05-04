package repositories

import (
	"context"
	"errors"

	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/domain/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

type IRoleRepository interface {
	FindByID(context.Context, string) (*models.Role, error)
	FindByCode(context.Context, string) (*models.Role, error)
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) FindByID(c context.Context, id string) (role *models.Role, err error){
	role = &models.Role{}
	if err = r.db.WithContext(c).Where("id = ?",id).First(&role).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrRoleNotFound)))
		}
		return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrSqlError)))
	}
	return
}
func (r *RoleRepository) FindByCode(c context.Context, code string) (role *models.Role, err error){
	role = &models.Role{}
	if err = r.db.WithContext(c).Where("code = ?",code).First(&role).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrRoleNotFound)))
		}
		return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrSqlError)))
	}
	return
}
