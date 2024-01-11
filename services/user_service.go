package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/yqisthi/Golang_Demo_2/models"
)

// UserService handles business logic for User
type UserService struct {
	DB *gorm.DB
}

// NewUserService creates a new instance of UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

// GetUsers retrieves all users from the database
func (service *UserService) GetUsers() []models.User {
	var users []models.User
	if err := service.DB.Find(&users).Error; err != nil {
		// Log the error for debugging
		fmt.Println("Error retrieving users:", err)
	}
    return users
}

// login user by email and password
func (service *UserService) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
    var user models.User
    err := service.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// CreateUser creates a new user in the database
func (service *UserService) CreateUser(userData *models.UserCreateInput) (*models.User, error) {
    user := models.User{
        Email:    userData.Email,
        Password: userData.Password,
        Name:     userData.Name,
        Role:     userData.Role,
    }

    if err := service.DB.Create(&user).Error; err != nil {
        return nil, err
    }

    return &user, nil
}
