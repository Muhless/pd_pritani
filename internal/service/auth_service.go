package service

import (
	"errors"
	"fmt"
	"os"
	"pd_pritani/internal/model"
	"pd_pritani/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password, role string) error
	GetProfile(userID uint) (*model.User, error)
}

type authService struct {
	db           *gorm.DB
	userRepo     repository.UserRepository
	adminRepo    repository.AdminRepository
	employeeRepo repository.EmployeeRepository
}

func NewAuthService(
	db *gorm.DB,
	userRepo repository.UserRepository,
	adminRepo repository.AdminRepository,
	employeeRepo repository.EmployeeRepository,
) AuthService {
	return &authService{db, userRepo, adminRepo, employeeRepo}
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
	secret := os.Getenv("JWT_SECRET")

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

	return s.db.Transaction(func(tx *gorm.DB) error {
		// make user
		user := &model.User{
			Username: username,
			Password: string(hashedPassword),
			Role:     role,
		}
		err = s.userRepo.Create(user)
		if err != nil {
			return errors.New("Username already used")
		}

		switch role {
		case "admin":
			admin := &model.Admin{
				UserID: user.ID,
				Name:   fmt.Sprintf("user%d", user.ID),
				Status: model.AdminStatusActive,
			}
			if err := tx.Create(admin).Error; err != nil {
				return errors.New("Failed creating admin record")
			}
		case "employee":
			employee := &model.Employee{
				UserID: user.ID,
				Name:   fmt.Sprintf("user%d", user.ID),
				Status: model.EmployeeStatusActive,
			}
			if err := tx.Create(employee).Error; err != nil {
				return errors.New("Failed creating employee record")
			}
		default:
			return errors.New("Role doesn't valid")
		}
		return nil
	})
}

func (s *authService) GetProfile(userID uint) (*model.User, error) {
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return user, nil

}
