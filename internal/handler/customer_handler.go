package handler

import (
	"log"
	"net/http"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model/customer"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	var customers []customer.Customer
	var total int64

	// Hitung total data
	config.DB.Model(&customer.Customer{}).Count(&total)

	// Ambil data dengan pagination
	config.DB.
		Limit(limit).
		Offset(offset).
		Order("id DESC").
		Find(&customers)

	// Response
	c.JSON(200, gin.H{
		"data":  customers,
		"total": total,
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
