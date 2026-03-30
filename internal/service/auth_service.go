package service

import (
	"errors"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
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
		return "", errors.New("Invalid Username or Password")
	}

	// make bycript
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid username or password")
	}

	// generate token
	token, err := generateJWT(user)
	if err != nil {
		return "", errors.New("Failed generating token")
	}
	return token, nil
}

func generateJWT(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte("secret_key"))
}
