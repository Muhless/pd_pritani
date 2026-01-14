package handler

import (
	"net/http"
	"pd_pritani/internal/config"
	"pd_pritani/internal/dto"
	"pd_pritani/internal/model"
	"pd_pritani/internal/model/employee"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(ctx *gin.Context) {
	var input dto.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{
				"message": "Input not valid",
				"error":   err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Hashing password failed"})
		return
	}

	user := model.User{
		Username:  input.Username,
		Password:  string(hashedPassword),
		Role:      input.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	switch input.Role {
	case "admin":
		admin := model.Admin{
			UserID: user.ID,
			Name:   input.Name,
			Email:  input.Email,
			Phone:  input.Phone,
			Photo:  input.Photo,
			Status: "active",
		}
		if err := config.DB.Create(&admin).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

	case "employee":
		employee := employee.Employee{
			UserID:  user.ID,
			Name:    input.Name,
			Phone:   input.Phone,
			Address: input.Address,
			Photo:   input.Photo,
			Status:  "active",
		}

		if err := config.DB.Create(&employee).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	user.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil", "user": user})
}
