package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model/employee"
	"pd_pritani/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(ctx *gin.Context) {
	var employee employee.Employee

	employee.Name = ctx.PostForm("name")
	employee.Phone = ctx.PostForm("phone")
	employee.Address = ctx.PostForm("address")
	employee.Status = ctx.PostForm("status")

	if err := utils.ValidatePhone(employee.Phone); err != nil {
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
	path := filepath.Join("uploads", "employee", fileName)

	if err := os.MkdirAll("uploads/employee", os.ModePerm); err != nil {
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

	baseURL := strings.TrimSuffix(os.Getenv("BASE_URL"), "/")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	employee.Photo = fmt.Sprintf("%s/%s", baseURL, path)

	if err := config.DB.Create(&employee).Error; err != nil {
		os.Remove(path)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create employee data",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

func GetEmployee(ctx *gin.Context) {
	var employee []employee.Employee

	if err := config.DB.Order("id ASC").Find(&employee).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Data not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

func GetEmployeeByID(ctx *gin.Context) {
	var employee employee.Employee
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	if err := config.DB.First(&employee, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "employee data not found",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

func UpdateEmployee(ctx *gin.Context) {
	var employee employee.Employee

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid ID"})
		return
	}

	if err := config.DB.First(&employee, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Employee not found"})
		return
	}

	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	address := ctx.PostForm("address")
	status := ctx.PostForm("status")

	if name != "" {
		employee.Name = name
	}
	if phone != "" {
		employee.Phone = phone
	}
	if address != "" {
		employee.Address = address
	}
	if status != "" {
		employee.Status = status
	}

	file, _ := ctx.FormFile("photo")
	if file != nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		savePath := filepath.Join("uploads", "employee", filename)

		os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Failed to upload photo"})
			return
		}

		baseURL := os.Getenv("BASE_URL")
		publicURL := fmt.Sprintf("%s/%s", baseURL, savePath)
		employee.Photo = publicURL
	}

	// SIMPAN KE DB
	if err := config.DB.Save(&employee).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Failed to update employee", "error": err.Error()})
		return
	}

	// RESPONSE
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}

func DeleteEmployee(ctx *gin.Context) {
	var employee employee.Employee
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	if err := config.DB.First(&employee, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data employee not found",
			"error":   err.Error(),
		})
		return
	}

	if err := config.DB.Delete(&employee).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete employee data",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    employee,
	})
}
