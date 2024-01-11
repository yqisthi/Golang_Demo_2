package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yqisthi/my-gin-postgres-app/models"
	"github.com/yqisthi/my-gin-postgres-app/services"
)

// UserController handles user-related HTTP requests
type UserController struct {
	UserService *services.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(db *gorm.DB) *UserController {
	userService := services.NewUserService(db)
	return &UserController{
		UserService: userService,
	}
}

// GetUsers handles the GET request to retrieve all users
func (controller *UserController) GetUsers(c *gin.Context) {
	users := controller.UserService.GetUsers()
	c.JSON(http.StatusOK, users)
}

// CreateUser handles the POST request to create a new user
func (controller *UserController) CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	controller.UserService.CreateUser(&user)
	c.JSON(http.StatusOK, user)
}
