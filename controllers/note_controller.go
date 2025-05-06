package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaitodecode/nyated-backend/common/lib/validation"
	"github.com/kaitodecode/nyated-backend/common/response"
	"github.com/kaitodecode/nyated-backend/common/util/pagination"
	"github.com/kaitodecode/nyated-backend/domain/dto"
	"github.com/kaitodecode/nyated-backend/services"
)

type NoteController struct {
	services services.IServiceRegistery
}

// Destroy implements INoteController.
func (f *NoteController) Destroy(c *gin.Context) {
	id := c.Param("id")
	if err := f.services.NoteService().HandleDestroy(c, id); err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}

	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Gin: c,
	})
}

// Index implements INoteController.
func (f *NoteController) Index(c *gin.Context) {
	pg := pagination.GetPaginationParams(c)
	query := &dto.GetNoteQuery{
		Pagination: pg,
		Filter: &dto.GetNoteFilter{
			Title: c.Query("title"),
			FolderID: c.Query("folder_id"),
		},
	}

	result, err := f.services.NoteService().HandleIndex(c.Request.Context(), query)
	if err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}
	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Data: result,
		Gin: c,
	})
}

// Show implements INoteController.
func (f *NoteController) Show(c *gin.Context) {
	id := c.Param("id")
	Note, err := f.services.NoteService().HandleShow(c, id)

	if err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}
	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Data: &Note,
		Gin: c,
	})
}

// Store implements INoteController.
func (f *NoteController) Store(c *gin.Context) {
	request := &dto.StoreNoteRequest{}
	if isErr, res := validation.ValidateBodyJson(c, request); isErr {
		response.HttpResponse(res)
		return
	}

	if err := f.services.NoteService().HandleStore(c.Request.Context(), request); err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}

	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Gin: c,
	})
}

// Update implements INoteController.
func (f *NoteController) Update(c *gin.Context) {
	id := c.Param("id")
	request := &dto.UpdateNoteRequest{}
	if isErr, res := validation.ValidateBodyJson(c, request); isErr {
		response.HttpResponse(res)
		return
	}

	if err := f.services.NoteService().HandleUpdate(c.Request.Context(), request,id); err != nil {
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Err: err,
			Gin: c,
		})
		return
	}

	response.HttpResponse(response.ParamHttpRes{
		Code: http.StatusOK,
		Gin: c,
	})
}

type INoteController interface {
	Index(*gin.Context)
	Show(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
}

func NewNoteController(services services.IServiceRegistery) INoteController {
	return &NoteController{services}
}
