package helpers

import (
	"bank-teller-backend/initializers"
	"bank-teller-backend/models"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func UserExists(email string, phone string) bool {
	var user models.User
	result := initializers.DB.Where("email = ? OR phone = ?", email, phone).First(&user)
	return result.Error == nil
}

func GetRoleByName(name string) (models.Role, error) {
	var role models.Role
	result := initializers.DB.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return models.Role{}, errors.New("role not found")
	}
	return role, nil
}

func GetUserFromToken(token string) (models.User, error) {
	// Parse the token to get the user ID
	userID, err := ParseToken(token)
	if err != nil {
		return models.User{}, err
	}

	// Get the user from the database
	var user models.User
	result := initializers.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return models.User{}, errors.New("ser not found")
	}

	return user, nil
}

func ParseToken(tokenString string) (string, error) {
	// Replace "yourSigningKey" with your actual signing key
	key := os.Getenv("JWT_SECRET")
	signingKey := []byte(key)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid token")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("Invalid token claims")
	}

	return userID, nil
}
