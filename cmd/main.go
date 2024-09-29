package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"testGorm/internal/application/handler"
	"testGorm/internal/application/services"
	"testGorm/internal/application/validation"
	"testGorm/internal/infrastructure/auth"
	"testGorm/internal/infrastructure/db"
	"testGorm/internal/infrastructure/persistence"
	"testGorm/internal/interfaces/http"
	"testGorm/internal/interfaces/rest"
)

var (
	userController *rest.UserController
	router         *gin.Engine
)

func init() {
	// Database setup in the infrastructure layer
	dsn := "root:expecto-patronum1379@tcp(127.0.0.1:3306)/test_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := db.SetupDatabase(dsn)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	// Dependency injection for repository and services
	userRepository := persistence.NewMySQLUserRepository(database)
	validationService := validation.NewValidationService(userRepository)
	jwtService := auth.NewJWTService("secret", "nadali")
	createUserHandler := handler.NewCreateUserCommandHandler(userRepository, validationService)
	loginUserHandler := handler.NewLoginUserCommandHandler(userRepository, validationService, jwtService)
	userService := services.NewUserService(userRepository, createUserHandler, validationService, loginUserHandler)

	// HTTP Controller setup
	userController = rest.NewUserController(userService)

	// Setup HTTP routes
	router = http.SetupRoutes(userController, jwtService)
}

func main() {
	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
