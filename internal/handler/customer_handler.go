package handler

import (
	"net/http"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model/customer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomer(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var customers []customer.Customer
	var total int64

	config.DB.Model(&customer.Customer{}).Count(&total)

	if err := config.DB.Limit(limit).Offset(offset).Find(&customers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Customer data not found",
			"error":   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": total,
		"data":    customers,
	})
}

func CreateCustomer(ctx *gin.Context) {

}
