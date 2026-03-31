package service

import (
	"errors"
	"log"
	"os"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password, role string) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Login(username, password string) (string, error) {
	// cari user by username
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		log.Println("error find user:", err)
		return "", errors.New("Invalid Username or Password")
	}

	// make bycript
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("error password:", err)
		return "", errors.New("Invalid username or password")
	}

	// generate token
	token, err := generateJWT(user)
	if err != nil {
		log.Println("error jwt:", err)
		return "", errors.New("Failed generating token")
	}
	return token, nil
}

func generateJWT(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	log.Println("secret generated in JWT:", secret)

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *authService) Register(username, password, role string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Failed hashing password")
	}

	// make user
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	// save to db
	err = s.userRepo.Create(user)
	if err != nil {
		return errors.New("Username already used")
	}
	return nil

}
