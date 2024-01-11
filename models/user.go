package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Role string `gorm:"column:role"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
    return "user"
}

type UserCreateInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    Name     string `json:"name"`
    Role     string `json:"role"`
}
