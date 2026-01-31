package handler

import (
	"net/http"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSales(ctx *gin.Context) {
	var sales []model.Sales

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	if err := config.DB.
		Order("id ASC").
		Preload("Customer").
		Limit(limit).Offset(offset).
		Find(&sales).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch sales data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Sales fetched successfully",
		"data":    sales,
	})
}
