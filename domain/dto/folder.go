package dto

import (
	"time"

	"github.com/kaitodecode/nyated-backend/common/util/pagination"
)

type StoreFolderRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	Description *string `json:"description,omitempty"`
	UserID      string
}

type UpdateFolderRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" validate:"required,min=3"`
	Description *string `json:"description,omitempty"`
	UserID      string
}

type GetFolderFilter struct {
	Name	string `json:"name"`
}
type GetFolderQuery struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Filter	   *GetFolderFilter		  `json:"filter"`
}

type GetFolderResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name" validate:"required,min=3"`
	Description *string `json:"description,omitempty"`
	UserID 	string		`json:"user_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	User	*UserResponse	`json:"user"`
}

type GetFolderPaginationResponse struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Filter	   *GetFolderFilter		  `json:"filter"`
	Result		[]GetFolderResponse		`json:"result"`
}

