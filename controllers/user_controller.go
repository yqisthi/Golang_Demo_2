package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yqisthi/Golang_Demo_2/models"
	"github.com/yqisthi/Golang_Demo_2/services"
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
    var userInput models.UserCreateInput
    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := controller.UserService.CreateUser(&userInput)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// Login handles the POST request for user login
func (controller *UserController) Login(c *gin.Context) {
    var loginData struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := controller.UserService.GetUserByEmailAndPassword(loginData.Email, loginData.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, user)
}
