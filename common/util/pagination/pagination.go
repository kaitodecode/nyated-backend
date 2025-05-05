package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int   `json:"page"`
	Limit    int   `json:"limit"`
	Offset   int   `json:"-"`
	Total    int64 `json:"total,omitempty"`
	LastPage int   `json:"last_page,omitempty"`
}

func GetPaginationParams(c *gin.Context) *Pagination {
	// Default values
	const (
		defaultPage  = 1
		defaultLimit = 10
		maxLimit     = 100
	)

	page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(defaultPage)))
	if err != nil || page < 1 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(defaultLimit)))
	if err != nil || limit < 1 {
		limit = defaultLimit
	} else if limit > maxLimit {
		limit = maxLimit
	}

	offset := (page - 1) * limit

	return &Pagination{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}
