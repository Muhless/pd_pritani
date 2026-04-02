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

type UpdateProfileRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Photo    string `json:"photo"`
	Status   string `json:"status"`
}

type AuthService interface {
	Login(username, password string) (string, error)
	Register(username, password, role string) error
	GetProfile(userID uint) (*model.User, error)
	UpdateProfile(userID uint, req UpdateProfileRequest) error
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(id uint, req UpdateProfileRequest) error
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

func (s *authService) UpdateProfile(userID uint, req UpdateProfileRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// get user
		user, err := s.userRepo.FindById(userID)
		if err != nil {
			return errors.New("User not found")
		}

		// update username if found
		if req.Username != "" {
			user.Username = req.Username
		}

		// update password if found
		if req.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				return errors.New("Failed hashing password")
			}
			user.Password = string(hashedPassword)
		}

		// save user
		if err := tx.Save(user).Error; err != nil {
			return errors.New("Update user data failed")
		}

		// update data by role
		switch user.Role {
		case "employee":
			employee, err := s.employeeRepo.FindByUserId(userID)
			if err != nil {
				return errors.New("User employee not found")
			}
			if req.Name != "" {
				employee.Name = req.Name
			}
			if req.Phone != "" {
				employee.Phone = req.Phone
			}
			if req.Address != "" {
				employee.Address = req.Address
			}
			if req.Photo != "" {
				employee.Photo = req.Photo
			}
			if err := tx.Save(employee).Error; err != nil {
				return errors.New("Update employee data failed")
			}
		case "admin":
			admin, err := s.adminRepo.FindByUserID(userID)
			if err != nil {
				return errors.New("data admin not found")
			}
			if req.Name != "" {
				admin.Name = req.Name
			}
			if req.Email != "" {
				admin.Email = req.Email
			}
			if req.Phone != "" {
				admin.Phone = req.Phone
			}
			if req.Address != "" {
				admin.Address = req.Address
			}
			if req.Photo != "" {
				admin.Photo = req.Photo
			}
			if err := tx.Save(admin).Error; err != nil {
				return errors.New("Update admin data failed")
			}
		}
		return nil
	})
}

func (s *authService) GetAllUsers() ([]model.User, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, errors.New("Failed getting user data")
	}
	return users, nil
}

func (s *authService) GetUserByID(id uint) (*model.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

func (s *authService) UpdateUser(id uint, req UpdateProfileRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		user, err := s.userRepo.FindById(id)
		if err != nil {
			return errors.New("User not found")
		}

		if req.Username != "" {
			user.Username = req.Username
		}

		if req.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				return errors.New("Failed hashing password")
			}
			user.Password = string(hashedPassword)
		}

		if err := tx.Save(user).Error; err != nil {
			return errors.New("update user data failed")
		}

		switch user.Role {
		case "employee":
			employee, err := s.employeeRepo.FindByUserId(id)
			if err != nil {
				return errors.New("emploee data not found")
			}
			if req.Name != "" {
				employee.Name = req.Name
			}
			if req.Phone != "" {
				employee.Phone = req.Phone
			}
			if req.Address != "" {
				employee.Address = req.Address
			}
			if req.Photo != "" {
				employee.Photo = req.Photo
			}
			if err := tx.Save(employee).Error; err != nil {
				return errors.New("failed to update employee data")
			}
		case "admin":
			admin, err := s.adminRepo.FindByUserID(id)
			if err != nil {
				return errors.New("admin data not found")
			}
			if req.Name != "" {
				admin.Name = req.Name
			}
			if req.Email != "" {
				admin.Email = req.Email
			}
			if req.Phone != "" {
				admin.Phone = req.Phone
			}
			if req.Address != "" {
				admin.Address = req.Address
			}
			if req.Photo != "" {
				admin.Photo = req.Photo
			}
			if err := tx.Save(admin).Error; err != nil {
				return errors.New("failed to update admin data")
			}
		}
		return nil
	})
}
