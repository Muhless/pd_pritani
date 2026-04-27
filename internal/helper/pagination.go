package helper

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

func GetPagination(ctx *gin.Context) (page, limit int) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}
	return page,limit
}

func NewPagination(page, limit int, total int64) Pagination  {
	totalPages := int(math.Ceil(float64(total)/float64(limit)))
	return  Pagination{
		Page: page,
		Limit: limit,
		Total: total,
		TotalPages: totalPages,
	}
}
