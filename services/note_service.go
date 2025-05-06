package services

import (
	"context"
	"github.com/kaitodecode/nyated-backend/common/util/mapper"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/repositories"
)

type NoteService struct {
	repositories repositories.IRepositoryRegistry
}

// HandleDestroy implements INoteService.
func (f *NoteService) HandleDestroy(c context.Context,id string) (error) {
	return f.repositories.NoteRepository().Destroy(c, id)
}

// HandleIndex implements INoteService.
func (f *NoteService) HandleIndex(c context.Context, query *dto.GetNoteQuery) (res *dto.GetNotePaginationResponse, err error) {
	res = &dto.GetNotePaginationResponse{}

	res, err = f.repositories.NoteRepository().FindAll(c, query)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// HandleShow implements INoteService.
func (f *NoteService) HandleShow(c context.Context,id string) (res *dto.GetNoteResponse,err error) { 
	Note, err := f.repositories.NoteRepository().FindByID(c,id)
	res = &dto.GetNoteResponse{}
	if err != nil {
		return nil, err
	}

	_ = mapper.MapModelToDTO(&Note, &res)

	return
}

// HandleStore implements INoteService.
func (f *NoteService) HandleStore(c context.Context, req *dto.StoreNoteRequest) error {
	return f.repositories.NoteRepository().Store(c, req)
}

// HandleUpdate implements INoteService.
func (f *NoteService) HandleUpdate(c context.Context, req *dto.UpdateNoteRequest, id string) error {
	return f.repositories.NoteRepository().Update(c,req,id)
}

type INoteService interface {
	HandleIndex(context.Context, *dto.GetNoteQuery) (*dto.GetNotePaginationResponse, error)
	HandleShow(context.Context, string) (*dto.GetNoteResponse, error)
	HandleStore(context.Context, *dto.StoreNoteRequest) error
	HandleUpdate(context.Context, *dto.UpdateNoteRequest, string) error
	HandleDestroy(context.Context, string) (error)
}

func NewNoteService(repositories repositories.IRepositoryRegistry) INoteService {
	return &NoteService{repositories}
}
