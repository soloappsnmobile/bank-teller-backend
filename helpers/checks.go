package helpers

import (
	"bank-teller-backend/initializers"
	"bank-teller-backend/models"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserExists(email string, phone string) bool {
	var user models.User
	result := initializers.DB.Where("email = ? OR phone = ?", email, phone).First(&user)
	return result.Error == nil
}

func ValidateTokenAndGetRole(tokenString string) (string, error) {
	// Replace this with your JWT secret key
	secretKey := os.Getenv("JWT_SECRET")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		// Token parsing failed
		return "", fmt.Errorf("%v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Get the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract claims from token")
	}
	// Get the role from the claims
	role, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("role not found in token")
	}

	return role, err
}

// parseRequest parses the request body and returns it as a map.
func ParseRequest(c *gin.Context) (map[string]interface{}, error) {
	var req map[string]interface{}
	err := c.ShouldBindJSON(&req)
	return req, err
}

// getUser retrieves the user from the database.
func GetUser(email string) (models.User, error) {
	var user models.User
	err := initializers.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// checkPassword compares the provided password with the user's hashed password.
func CheckPassword(user models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// getRole retrieves the user's role from the database.
func GetRole(user models.User) (models.Role, error) {
	var role models.Role
	err := initializers.DB.Where("id = ?", user.RoleID).First(&role).Error
	return role, err
}

// generateToken generates a JWT token for the user.
func GenerateToken(user models.User, role models.Role) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    role.Name,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
