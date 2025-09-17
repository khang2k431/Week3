package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/khang2k431/Week3/config"
	"github.com/khang2k431/Week3/controllers"
	"github.com/khang2k431/Week3/middlewares"
	"github.com/khang2k431/Week3/models"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()

	if config.DB == nil {
		log.Fatal("Database connection failed")
	}

	// migrate
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected
	auth := r.Group("/api")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			claims, _ := c.Get("claims")
			c.JSON(200, gin.H{"claims": claims})
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running at http://localhost:" + port)
	r.Run(":" + port)
}
