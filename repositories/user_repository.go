package repositories

import (
	"context"
	"errors"

	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Register(context.Context, *dto.UserRegisterRequest) error
	Update(context.Context, *dto.UpdateUserRequest, string) error
	FindByEmail(context.Context, string) (*models.User, error)
	FindByID(context.Context, string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}




func (r *UserRepository) Update(c context.Context, req *dto.UpdateUserRequest, id string) error {

	user := &models.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		RoleID: req.RoleID,
	}

	if err := r.db.WithContext(c).Where("id = ?",id).Updates(&user).Error; err != nil {
		return AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrSqlError)))
	}
	return nil
}
func (r *UserRepository) Register(c context.Context, req *dto.UserRegisterRequest) error {

	user := &models.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		RoleID: req.RoleID,
	}

	if err := r.db.WithContext(c).Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey){
			return AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrUserAlreadyExist)))
		}
		return AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrSqlError)))
	}
	return nil
}

func (r *UserRepository) FindByEmail(c context.Context, email string) (user *models.User, err error){
	if err = r.db.WithContext(c).Preload(constants.PRELOAD_ROLE).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrUserNotFound)))
		}
		return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrSqlError)))
	}
	return user, nil
}

func (r *UserRepository) FindByID(c context.Context, id string) (user *models.User, err error){
	if err = r.db.WithContext(c).Preload(constants.PRELOAD_ROLE).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrUserNotFound)))
		}
		return nil, AppError.WrapError(errors.New(AppError.GetMessage(c, AppError.ErrSqlError)))
	}
	return user, nil
}