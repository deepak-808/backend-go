package handlers

import (
	controller "example/main/db"
	"example/main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProfile retrieves the profile data of the authenticated user
func GetProfile(c *gin.Context) {
	// Extract the user ID from the context set by the AuthMiddleware
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the collection
	collection := controller.GetCollection()

	// Query the database to find the user's profile
	var user models.User
	filter := bson.M{"_id": userID}

	err := collection.Users.FindOne(c, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		}
		return
	}

	// Convert user to UserProfileDTO
	userProfileDTO := user.ToDTO()

	// Return the user profile
	c.JSON(http.StatusOK, userProfileDTO)
}
