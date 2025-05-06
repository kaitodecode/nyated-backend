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
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

// Destroy implements INoteRepository.
func (f *NoteRepository) Destroy(c context.Context, id string) error {
	logrus.Info("NoteRepository.Destroy()")
	logrus.Infof("id => %s", id)

	tx := f.db.WithContext(c).Where("id = ?", id).Delete(&models.Note{})
	if tx.Error != nil {
		logrus.Infof("err => %v", tx.Error)
		return errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}

	if tx.RowsAffected == 0 {
		return errors.New(AppError.GetMessage(c, AppError.ErrNoteNotFound))
	}

	return nil
}


// FindByID implements INoteRepository.
func (f *NoteRepository) FindByID(c context.Context,id string) (*models.Note, error) {
	var note models.Note
	if err := f.db.WithContext(c).Preload(constants.PRELOAD_FOLDER).Where("id = ?",id).First(&note).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(AppError.GetMessage(c, AppError.ErrNoteNotFound))
		}
		return nil, errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return &note, nil
}

// Store implements INoteRepository.
func (f *NoteRepository) Store(c context.Context, req *dto.StoreNoteRequest) error {
	if err := f.db.WithContext(c).Model(&models.Note{}).Create(&models.Note{
		Title: req.Title,
		FolderID: req.FolderID,
		Content: req.Content,
	}).Error; err != nil {
		return errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return nil
}

// Update implements INoteRepository.
func (f *NoteRepository) Update(c context.Context, req *dto.UpdateNoteRequest, id string) error {
	if err := f.db.WithContext(c).Model(&models.Note{}).Where("id = ?", id).Updates(&models.Note{
		Title: req.Title,
		Content: req.Content,
		}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(AppError.GetMessage(c, AppError.ErrFolderNotFound))
		}
		return errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}
	return nil
}



// FindAll implements INoteRepository.
func (f *NoteRepository) FindAll(c context.Context, req *dto.GetNoteQuery) (*dto.GetNotePaginationResponse, error) {
	var notes []models.Note
	var result []dto.GetNoteResponse
	var total int64
	user, err := util.GetUser(c)
	if err != nil {
		return nil, err
	}
	db := f.db.WithContext(c).Joins("JOIN folders ON folders.id = notes.folder_id").Where("folders.user_id", user.ID).Preload(constants.PRELOAD_FOLDER).Model(&models.Note{})

	if req.Filter.Title != "" {
		db = db.Where("title ILIKE ?", "%"+req.Filter.Title+"%")
	}

	if req.Filter.FolderID != "" {
		db = db.Where("folder_id = ?", req.Filter.FolderID)
	}

	db.Count(&total)

	if req.Pagination.Page <= 0 {
		req.Pagination.Page = 1
	}
	if req.Pagination.Limit <= 0 {
		req.Pagination.Limit = 10
	}

	lastPage := int((total + int64(req.Pagination.Limit) - 1) / int64(req.Pagination.Limit))

	offset := (req.Pagination.Page - 1) * req.Pagination.Limit

	err = db.Limit(req.Pagination.Limit).Offset(offset).Order("created_at desc").Find(&notes).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(AppError.GetMessage(c, AppError.ErrFolderNotFound))
		}
		return nil, errors.New(AppError.GetMessage(c, AppError.ErrSqlError))
	}

	_ = mapper.MapModelToDTO(&notes, &result)

	res := &dto.GetNotePaginationResponse{
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

type INoteRepository interface {
	FindAll(context.Context, *dto.GetNoteQuery) (*dto.GetNotePaginationResponse, error)
	FindByID(context.Context, string) (*models.Note, error)
	Store(context.Context, *dto.StoreNoteRequest) error
	Update(context.Context, *dto.UpdateNoteRequest, string) error
	Destroy(context.Context, string) error
}

func NewNoteRepository(db *gorm.DB) INoteRepository {
	return &NoteRepository{db}
}
