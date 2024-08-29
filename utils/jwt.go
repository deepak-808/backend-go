package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

// init loads environment variables and sets up the JWT key
func init() {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		panic("JWT_SECRET_KEY is not set in the environment variables")
	}
	jwtKey = []byte(jwtSecretKey)
}

// GetJwtKey returns the JWT secret key
func GetJwtKey() []byte {
	return jwtKey
}

// GenerateToken generates a JWT token for the given user ID
func GenerateToken(userID string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expiration time
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// HashPassword hashes a plaintext password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a hashed password with a plaintext password
func CheckPassword(providedPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}
