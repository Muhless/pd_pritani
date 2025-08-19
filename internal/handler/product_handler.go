package handler

import (
	"fmt"
	"net/http"
	"os"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProduct(ctx *gin.Context) {
	var products []model.Product
	if err := config.DB.Find(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch product data",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "data": products})
}

func CreateProduct(ctx *gin.Context) {
	var product model.Product

	// ambil field selain file
	product.Name = ctx.PostForm("name")
	product.Type = ctx.PostForm("type")
	product.Stock, _ = strconv.Atoi(ctx.PostForm("stock"))
	product.Price, _ = strconv.Atoi(ctx.PostForm("price"))

	// ambil file photo
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Photo is required",
			"error":   err.Error(),
		})
		return
	}

	// prefix agar tidak ada file dengan nama sama
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	path := fmt.Sprintf("uploads/%s", filename)
	os.MkdirAll("uploads", os.ModePerm)

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to upload photo",
			"error":   err.Error(),
		})
		return
	}

	product.Photo = path

	// TODO:GUNAKAN INI KALO GAADA FILE
	// bind JSON ke struct product
	// if err := ctx.ShouldBindJSON(&product); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"success": false,
	// 		"message": "Failed to create data product",
	// 		"error":   err.Error(),
	// 	})
	// 	return
	// }

	// simpan ke database
	if err := config.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create product data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true,
		"data": product,
	})
}
