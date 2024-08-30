package handlers

import (
	"context"
	controller "example/main/db"
	"example/main/models"
	"example/main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Function for logging in
func Login(c *gin.Context) {
	collections := controller.GetCollection()

	// Parse JSON request body into user model
	var loginData models.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Find the user by email
	var user models.User
	err := collections.Users.FindOne(context.TODO(), bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Check the provided password against the stored hashed password
	if !utils.CheckPassword(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// Generate a JWT token
	userId := user.ID
	token, err := utils.GenerateToken(userId.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	data := gin.H{
		"token": token,
		"id":    user.ToDTO().ID,
	}

	// Respond with the JWT token and user data
	c.JSON(http.StatusOK, data)

}

// Register function handles new user registration
func Register(c *gin.Context) {
	collections := controller.GetCollection()

	// Parse JSON request body into user model
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Check if the email already exists
	var existingUser models.User
	err := collections.Users.FindOne(context.TODO(), bson.M{"email": newUser.Email}).Decode(&existingUser)
	if err == nil {
		// Email already exists
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	} else if err != mongo.ErrNoDocuments {
		// Some other error occurred
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Hash the user's password before storing it
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	newUser.Password = hashedPassword
	newUser.ID = primitive.NewObjectID()

	// Insert the new user into the database
	_, err = collections.Users.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Example of a protected route that requires authentication
func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}
