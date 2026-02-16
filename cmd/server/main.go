package main

import (
	"log"

	"github.com/aryansehgal-tech/NotesApp/internal/config"
	"github.com/gin-gonic/gin"

)

func main(){
	// load configuration
	cfg:= config.LoadConfig()

	// create Gin router
	router:= gin.Default()

	// Health route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Server running on port:", cfg.AppPort)

	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}