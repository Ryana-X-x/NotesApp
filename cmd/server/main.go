package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aryansehgal-tech/NotesApp/internal/handler"
	"github.com/aryansehgal-tech/NotesApp/internal/models"
	"github.com/aryansehgal-tech/NotesApp/internal/repository"
	"github.com/aryansehgal-tech/NotesApp/internal/service"
)

func main() {

	// ==============================
	// DATABASE CONNECTION
	// ==============================

	dsn := "host=localhost user=postgres password=postgres dbname=notesapp port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate only User model
	db.AutoMigrate(&models.User{})

	// ==============================
	// DEPENDENCY INJECTION
	// ==============================

	// Repository
	userRepo := repository.NewUserRepository(db)

	// Service
	authService := service.NewAuthService(userRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)

	// ==============================
	// ROUTER SETUP
	// ==============================

	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// Test protected route (optional)
	// router.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "you are authorized"})
	// })

	// ==============================
	// START SERVER
	// ==============================

	router.Run(":8080")
}
