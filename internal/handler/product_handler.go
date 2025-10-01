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
	"github.com/shopspring/decimal"
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
	// validasi stock
	stock, err := strconv.Atoi(ctx.PostForm("stock"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid stock value",
		})
		return
	}
	product.Stock = stock

	priceStr := ctx.PostForm("price")
	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid price value",
		})
		return
	}

	product.Price = price

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

	// buat folder uploads bila belum ada
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create folder",
			"error":   err.Error(),
		})
		return
	}

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to upload photo",
			"error":   err.Error(),
		})
		return
	}
	baseURL := os.Getenv("BASE_URL")
	product.Photo = fmt.Sprintf("%s/%s", baseURL, path)

	// simpan ke database
	if err := config.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create product data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var product model.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "ID not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	var product model.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "ID not found",
			"error":   err.Error(),
		})
		return
	}

	file, _ := ctx.FormFile("photo")
	if file != nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		path := fmt.Sprintf("upload/%s", filename)

		if err := os.Mkdir("uploads", os.ModePerm); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to create folder", "error": err.Error()})
			return
		}

		if err := ctx.SaveUploadedFile(file, path); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Failed to upload photo", "error": err.Error()})
			return
		}

		product.Photo = fmt.Sprintf("http://localhost:8080/%s", path)
	}

	var input model.UpdateProductInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to bind data",
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Model(&product).Updates(input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update data product",
			"error":   err.Error(),
		})
		return
	}

	// refresh data
	config.DB.First(&product, id)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully updated product data",
		"data":    product,
	})
}

func DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}
	var product model.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Product not found",
			"error":   err.Error(),
		})
		return
	}

	// tidak perlu id lagi sudah di first
	if err := config.DB.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete product data",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product data successfully deleted",
		"data":    product,
	})
}
