package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model"
	"pd_pritani/internal/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateProfile(ctx *gin.Context) {
	var profile model.Profile

	profile.Name = ctx.PostForm("name")
	profile.Phone = ctx.PostForm("phone")

	if err := utils.ValidatePhone(profile.Phone); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid phone number",
			"error":   err.Error(),
		})
		return
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Photo is required",
			"error":   err.Error(),
		})
		return
	}

	// prefix
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	path := filepath.Join("uploads", "profile", fileName)

	if err := os.MkdirAll("uploads/profile", os.ModePerm); err != nil {
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
	profile.Photo = fmt.Sprintf("%s/%s", baseURL, path)

	if err := config.DB.Create(&profile).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create profile data",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    profile,
	})
}

func GetProfile(ctx *gin.Context) {
	var profile []model.Profile

	if err := config.DB.Find(&profile).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Data not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    profile,
	})
}

func GetProfileByID(ctx *gin.Context) {
	var profile model.Profile
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	if err := config.DB.First(&profile, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Profile data not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    profile,
	})
}

func UpdateProfile(ctx *gin.Context) {
	var profile model.Profile
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	if err := config.DB.Find(&profile, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "ID not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    profile,
	})
}
