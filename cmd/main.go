package main

import (
	controller "example/main/db"
	"example/main/handlers"
	"example/main/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Starting the application")
	collection := controller.GetCollection()
	fmt.Println(collection)
	r := gin.Default()

	// Public routes (do not require authentication)
	publicRoutes := r.Group("/api")
	// commonRoutes := r.Group("/")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
		// commonRoutes.GET("/", handlers.Getdata)
	}

	// Protected routes (require authentication)
	protectedRoutes := publicRoutes.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.POST("/profile", handlers.GetProfile)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port
	}
	r.Run(":" + port)
}
