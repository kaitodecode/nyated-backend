package repositories

import (
	"context"
	"errors"

	AppError "github.com/kaitodecode/nyated-backend/common/error"
	"github.com/kaitodecode/nyated-backend/common/util"
	"github.com/kaitodecode/nyated-backend/common/util/mapper"
	"github.com/kaitodecode/nyated-backend/common/util/pagination"
	"github.com/kaitodecode/nyated-backend/constants"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/domain/models"
	"gorm.io/gorm"
)

type FolderRepository struct {
	db *gorm.DB
}

// Destroy implements IFolderRepository.
func (f *FolderRepository) Destroy(c context.Context,id string) error {
	if err := f.db.WithContext(c).Where("id = ?", id).Delete(&models.Folder{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(AppError.GetMessage(c, AppError.ErrFolderNotFound))
		}
		return errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return nil
}

// FindByID implements IFolderRepository.
func (f *FolderRepository) FindByID(c context.Context,id string) (*models.Folder, error) {
	var folder models.Folder
	if err := f.db.WithContext(c).Preload(constants.PRELOAD_USER).Where("id = ?",id).First(&folder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(AppError.GetMessage(c, AppError.ErrFolderNotFound))
		}
		return nil, errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return &folder, nil
}

// Store implements IFolderRepository.
func (f *FolderRepository) Store(c context.Context,req *dto.StoreFolderRequest) error {
	if err := f.db.WithContext(c).Model(&models.Folder{}).Create(&models.Folder{
		Name: req.Name,
		Description: req.Description,
		UserID: req.UserID,
	}).Error; err != nil {
		return errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return nil
}

// Update implements IFolderRepository.
func (f *FolderRepository) Update(c context.Context,req *dto.UpdateFolderRequest,id string) error {
	if err := f.db.WithContext(c).Model(&models.Folder{}).Where("user_id = ?", req.UserID).Where("id = ?",id).Updates(&models.Folder{
		Name: req.Name,
		Description: req.Description,
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(AppError.GetMessage(c, AppError.ErrFolderNotFound))
		}
		return errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return nil
}



// FindAll implements IFolderRepository.
func (f *FolderRepository) FindAll(c context.Context, req *dto.GetFolderQuery) (*dto.GetFolderPaginationResponse, error) {
	var folders []models.Folder
	var result []dto.GetFolderResponse
	var total int64

	db := f.db.WithContext(c).Preload(constants.PRELOAD_USER).Model(&models.Folder{})

	if req.Filter.Name != "" {
		db = db.Where("name ILIKE ?", "%"+req.Filter.Name+"%")
	}

	user, err := util.GetUser(c)

	if err != nil {
		return nil, err
	}

	db = db.Where("user_id = ?", user.ID)

	db.Count(&total)

	if req.Pagination.Page <= 0 {
		req.Pagination.Page = 1
	}
	if req.Pagination.Limit <= 0 {
		req.Pagination.Limit = 10
	}
	lastPage := int((total + int64(req.Pagination.Limit) - 1) / int64(req.Pagination.Limit))

	offset := (req.Pagination.Page - 1) * req.Pagination.Limit

	err = db.Limit(req.Pagination.Limit).Offset(offset).Order("created_at desc").Find(&folders).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(AppError.GetMessage(c, AppError.ErrFolderNotFound))
		}
		return nil, errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}

	_ = mapper.MapModelToDTO(&folders, &result)

	res := &dto.GetFolderPaginationResponse{
		Pagination: &pagination.Pagination{
			Page:     req.Pagination.Page,
			Limit:    req.Pagination.Limit,
			Offset:   offset,
			Total:    total,
			LastPage: lastPage,
		},
		Filter: req.Filter,
		Result: result,
	}
	return res, nil
}

type IFolderRepository interface {
	FindAll(context.Context, *dto.GetFolderQuery) (*dto.GetFolderPaginationResponse, error)
	FindByID(context.Context, string) (*models.Folder, error)
	Store(context.Context, *dto.StoreFolderRequest) error
	Update(context.Context, *dto.UpdateFolderRequest, string) error
	Destroy(context.Context, string) error
}

func NewFolderRepository(db *gorm.DB) IFolderRepository {
	return &FolderRepository{db}
}
