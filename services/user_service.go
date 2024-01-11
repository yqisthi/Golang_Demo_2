package services

import (
	"github.com/jinzhu/gorm"
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
