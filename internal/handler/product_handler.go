package handler

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func GetProduct(ctx *gin.Context) {
	var products []model.Product
	if err := config.DB.Order("id ASC").Find(&products).Error; err != nil {
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
	contentType := ctx.ContentType()

	var product model.Product

	// Parse product data berdasarkan content type
	if strings.Contains(contentType, "application/json") {
		// Mode JSON (tanpa file upload)
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid JSON body",
				"error":   err.Error(),
			})
			return
		}

		// Untuk JSON mode, photo bisa berupa URL atau kosong
		// Validasi basic
		if err := validateProduct(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		// Simpan ke database
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
		return
	}

	// Multipart form mode (dengan file upload)
	product.Name = strings.TrimSpace(ctx.PostForm("name"))
	product.Type = strings.TrimSpace(ctx.PostForm("type"))

	// Parse stock
	stockStr := ctx.PostForm("stock")
	if stockStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Stock is required",
		})
		return
	}
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid stock value, must be a number",
		})
		return
	}
	product.Stock = stock

	// Parse price
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

	// Validasi basic fields
	if err := validateProduct(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Ambil dan validasi file photo
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Photo is required",
			"error":   err.Error(),
		})
		return
	}

	// Validasi file size (max 5MB)
	maxFileSize := int64(5 * 1024 * 1024) // 5MB
	if file.Size > maxFileSize {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "File size must be less than 5MB",
		})
		return
	}

	// Validasi file type
	if err := validateImageType(file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Generate unique filename menggunakan UUID
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	path := filepath.Join("uploads", filename)

	// Buat folder uploads bila belum ada
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create upload folder",
			"error":   err.Error(),
		})
		return
	}

	// Simpan file
	if err := ctx.SaveUploadedFile(file, path); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to upload photo",
			"error":   err.Error(),
		})
		return
	}

	// Generate photo URL
	baseURL := strings.TrimSuffix(os.Getenv("BASE_URL"), "/")
	if baseURL == "" {
		baseURL = "http://localhost:8080" // default untuk development
	}
	product.Photo = fmt.Sprintf("%s/%s", baseURL, path)

	// Simpan ke database
	if err := config.DB.Create(&product).Error; err != nil {
		// Cleanup: hapus file yang sudah diupload jika database gagal
		os.Remove(path)

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

func validateProduct(product *model.Product) error {
	if product.Name == "" {
		return fmt.Errorf("product name is required")
	}

	if product.Type == "" {
		return fmt.Errorf("product type is required")
	}

	if product.Stock < 0 {
		return fmt.Errorf("stock cannot be negative")
	}

	if product.Price.IsNegative() {
		return fmt.Errorf("price cannot be negative")
	}

	if product.Price.IsZero() {
		return fmt.Errorf("price must be greater than zero")
	}

	return nil
}

func validateImageType(file *multipart.FileHeader) error {
	// Buka file untuk cek content type
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file")
	}
	defer src.Close()

	// Baca 512 bytes pertama untuk detect content type
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return fmt.Errorf("failed to read file")
	}

	// Detect content type
	contentType := http.DetectContentType(buffer)

	// Allowed image types
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
		"image/gif":  true,
		"image/webp": true,
	}

	if !allowedTypes[contentType] {
		return fmt.Errorf("file must be an image (JPEG, PNG, GIF, or WebP)")
	}

	return nil
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
			"message": "Invalid product ID format",
		})
		return
	}

	var product model.Product

	// Cek apakah product exists
	if err := config.DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Product not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch product",
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete product",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product successfully deleted",
		"data": gin.H{
			"id": id,
		},
	})
}
