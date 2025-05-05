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

type FolderController struct {
	services services.IServiceRegistery
}

// Destroy implements IFolderController.
func (f *FolderController) Destroy(c *gin.Context) {
	id := c.Param("id")
	if err := f.services.FolderService().HandleDestroy(c, id); err != nil {
		message := http.StatusText(http.StatusBadRequest)
		response.HttpResponse(response.ParamHttpRes{
			Code: http.StatusBadRequest,
			Message: &message,
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

// Index implements IFolderController.
func (f *FolderController) Index(c *gin.Context) {
	pg := pagination.GetPaginationParams(c)
	query := &dto.GetFolderQuery{
		Pagination: pg,
		Filter: &dto.GetFolderFilter{
			Name: c.Query("name"),
		},
	}
	result, err := f.services.FolderService().HandleIndex(c.Request.Context(), query)
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

// Show implements IFolderController.
func (f *FolderController) Show(c *gin.Context) {
	id := c.Param("id")
	folder, err := f.services.FolderService().HandleShow(c, id)

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
		Data: &folder,
		Gin: c,
	})
}

// Store implements IFolderController.
func (f *FolderController) Store(c *gin.Context) {
	request := &dto.StoreFolderRequest{}
	if isErr, res := validation.ValidateBodyJson(c, request); isErr {
		response.HttpResponse(res)
		return
	}

	if err := f.services.FolderService().HandleStore(c.Request.Context(), request); err != nil {
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

// Update implements IFolderController.
func (f *FolderController) Update(c *gin.Context) {
	id := c.Param("id")
	request := &dto.UpdateFolderRequest{}
	if isErr, res := validation.ValidateBodyJson(c, request); isErr {
		response.HttpResponse(res)
		return
	}

	if err := f.services.FolderService().HandleUpdate(c.Request.Context(), request,id); err != nil {
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

type IFolderController interface {
	Index(*gin.Context)
	Show(*gin.Context)
	Store(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
}

func NewFolderController(services services.IServiceRegistery) IFolderController {
	return &FolderController{services}
}
