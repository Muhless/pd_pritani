package handler

import (
	"log"
	"net/http"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCustomers(ctx *gin.Context) {
	var customers []model.Customer

	if err := config.DB.Order("id ASC").Find(&customers).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Data not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    customers,
	})
}

func CreateCustomer(ctx *gin.Context) {
	var customer model.Customer

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
		if strings.Contains(err.Error(), "duplicate key") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Nomor telepon sudah terdaftar",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create customer data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully created customer data",
		"data":    customer,
	})

}
