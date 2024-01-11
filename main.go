package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/your_username/my-gin-postgres-app/controllers"
)

var db *gorm.DB
var err error

func main() {
	// Connect to PostgreSQL
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=your_username dbname=your_database password=your_password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Auto Migrate
	db.AutoMigrate(&controllers.User{})

	// Create a new Gin router
	router := gin.Default()

	// Initialize user controller
	userController := controllers.NewUserController(db)

	// Define routes
	router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.CreateUser)

	// Run the application
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
