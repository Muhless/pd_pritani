package main

import (
	"fmt"
	"log"
	"pd_pritani/internal/config"
	"pd_pritani/internal/model"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	godotenv.Load()
	db := config.ConnectDB()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{
		Username: "admin",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("User admin successfully created:", user.ID)
}
