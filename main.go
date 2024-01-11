package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yqisthi/Golang_Demo_2/controllers"
	"github.com/yqisthi/Golang_Demo_2/models"
)

var db *gorm.DB
var err error

func main() {
	// Connect to PostgreSQL
	db, err = gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=haus_flutter password=12345678 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Auto Migrate
	db.AutoMigrate(&models.User{})

	// Create a new Gin router
	router := gin.Default()

	// Initialize user controller
	userController := controllers.NewUserController(db)

	// Define routes
	// GET ALL DATA
	router.GET("/users", userController.GetUsers)
	// CREATE USER
	router.POST("/users", userController.CreateUser)
	// LOGIN USER
	router.GET("/login", userController.Login)

	// Run the application
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
