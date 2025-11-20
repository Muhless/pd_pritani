package handler

import (
	"log"
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
	var customer customer.Customer

	if err := ctx.ShouldBindJSON(&customer); err != nil {
		log.Println("JSON bind error:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Create(&customer).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create customer data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Successfully created customer data",
		"customer": customer,
	})

}
