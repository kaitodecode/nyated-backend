package services

import (
	"context"

	"github.com/kaitodecode/nyated-backend/common/util"
	"github.com/kaitodecode/nyated-backend/common/util/mapper"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/repositories"
)

type FolderService struct {
	repositories repositories.IRepositoryRegistry
}

// HandleDestroy implements IFolderService.
func (f *FolderService) HandleDestroy(c context.Context,id string) (error) {
	return f.repositories.FolderRepository().Destroy(c, id)
}

// HandleIndex implements IFolderService.
func (f *FolderService) HandleIndex(c context.Context, query *dto.GetFolderQuery) (res *dto.GetFolderPaginationResponse, err error) {
	res = &dto.GetFolderPaginationResponse{}

	res, err = f.repositories.FolderRepository().FindAll(c, query)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleShow implements IFolderService.
func (f *FolderService) HandleShow(c context.Context,id string) (res *dto.GetFolderResponse,err error) { 
	folder, err := f.repositories.FolderRepository().FindByID(c,id)
	res = &dto.GetFolderResponse{}
	if err != nil {
		return nil, err
	}

	_ = mapper.MapModelToDTO(&folder, &res)

	return
}

// HandleStore implements IFolderService.
func (f *FolderService) HandleStore(c context.Context, req *dto.StoreFolderRequest) error {
	user, err := util.GetUser(c)
	if err != nil {
		return err
	}
	req.UserID = user.ID
	return f.repositories.FolderRepository().Store(c, req)
}

// HandleUpdate implements IFolderService.
func (f *FolderService) HandleUpdate(c context.Context, req *dto.UpdateFolderRequest, id string) error {
	user, err := util.GetUser(c)
	if err != nil {
		return err
	}
	req.UserID = user.ID
	return f.repositories.FolderRepository().Update(c,req,id)
}

type IFolderService interface {
	HandleIndex(context.Context, *dto.GetFolderQuery) (*dto.GetFolderPaginationResponse, error)
	HandleShow(context.Context, string) (*dto.GetFolderResponse, error)
	HandleStore(context.Context, *dto.StoreFolderRequest) error
	HandleUpdate(context.Context, *dto.UpdateFolderRequest, string) error
	HandleDestroy(context.Context, string) (error)
}

func NewFolderService(repositories repositories.IRepositoryRegistry) IFolderService {
	return &FolderService{repositories}
}
