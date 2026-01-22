package handler

import (
	"net/http"
	"pd_pritani/internal/config"
	"pd_pritani/internal/dto"
	"pd_pritani/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(ctx *gin.Context) {
	var input dto.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// validasi role
	if input.Role != "admin" && input.Role != "employee" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Role tidak valid",
		})
		return
	}

	// cek username
	var count int64
	config.DB.Model(&model.User{}).
		Where("username = ?", input.Username).
		Count(&count)

	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Username sudah digunakan",
		})
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal hash password",
		})
		return
	}

	// TRANSACTION
	tx := config.DB.Begin()

	user := model.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
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
		}

		if err := tx.Create(&admin).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

	case "employee":
		employee := model.Employee{
			UserID:  user.ID,
			Name:    input.Name,
			Phone:   input.Phone,
			Address: input.Address,
			Photo:   input.Photo,
			Status:  "active",
		}

		if err := tx.Create(&employee).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	tx.Commit()

	user.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi berhasil",
		"user":    user,
	})
}
