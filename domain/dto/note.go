package dto

import (
	"time"

	"github.com/kaitodecode/nyated-backend/common/util/pagination"
)

type StoreNoteRequest struct {
	Title   string `json:"title" validate:"required,min=3"`
	Content string `json:"content" validate:"required,min=3"`
	FolderID  string	`json:"folder_id" validate:"required,uuid"`
}

type UpdateNoteRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title" validate:"required,min=3"`
	Content string `json:"content" validate:"required,min=3"`
}

type GetNoteFilter struct {
	Title string `json:"title"`
	FolderID string `json:"folder_id"`
}
type GetNoteQuery struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Filter     *GetNoteFilter         `json:"filter"`
}

type GetNoteResponse struct {
	ID        string        `json:"id"`
	Title     string        `json:"title" validate:"required,min=3"`
	Content   string        `json:"content"`
	FolderID    string        `json:"folder_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Folder      *GetNoteResponse `json:"folder"`
}

type GetNotePaginationResponse struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Filter     *GetNoteFilter         `json:"filter"`
	Result     []GetNoteResponse      `json:"result"`
}
