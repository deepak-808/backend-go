package main

import (
	"example/main/handlers"
	"example/main/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the application")
	collection := GetCollection()
	fmt.Println(collection)
	r := gin.Default()

	// Public routes (do not require authentication)
	publicRoutes := r.Group("/public")
	commonRoutes := r.Group("/")
	{
		publicRoutes.POST("/login", handlers.Login)
		publicRoutes.POST("/register", handlers.Register)
		commonRoutes.GET("/", handlers.Getdata)
	}

	// Protected routes (require authentication)
	protectedRoutes := r.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		// Protected routes here
	}

	r.Run(":8080")
}
