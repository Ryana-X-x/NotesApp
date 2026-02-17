package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/aryansehgal-tech/NotesApp/internal/config"
	"github.com/aryansehgal-tech/NotesApp/internal/database"
	"github.com/aryansehgal-tech/NotesApp/internal/handler"
	"github.com/aryansehgal-tech/NotesApp/internal/repository"
	"github.com/aryansehgal-tech/NotesApp/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	db := database.ConnectDatabase(cfg)

	// Dependency injection
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.POST("/auth/register", authHandler.Register)

	log.Println("Server running on port:", cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
