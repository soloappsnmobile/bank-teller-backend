package admin

import (
	"bank-teller-backend/helpers"
	"bank-teller-backend/initializers"
	"bank-teller-backend/models"
	"bank-teller-backend/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateTeller(c *gin.Context) {
	// Get the user's role from the context
	role, exists := c.Get("role")
	if !exists || role != "Admin" {
		helpers.RespondWithError(c, http.StatusUnauthorized, "Only admins can create tellers", "401")
		return
	}

	// Parse the request body
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request data", "400")
		return
	}

	// Validate the request
	if err := validators.ValidateRequest(req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "400")
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := initializers.DB.Where("email = ?", req["email"].(string)).First(&existingUser).Error; err == nil {
		helpers.RespondWithError(c, http.StatusConflict, "User with this email already exists", "409")
		return
	}

	tellerRole := models.Role{Name: "Teller"}
	// Check if the teller role already exists
	if err := initializers.DB.Where("name = ?", "Teller").First(&tellerRole).Error; err != nil {
		// If not, create it
		tellerRole = models.Role{Name: "Teller"}
		initializers.DB.Create(&tellerRole)
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password", "500")
		return
	}

	// Create a new user with the teller role
	tellerUser := models.User{
		FirstName: req["first_name"].(string),
		LastName:  req["last_name"].(string),
		Phone:     req["phone"].(string),
		Email:     req["email"].(string),
		Password:  string(hashedPassword),
		RoleID:    tellerRole.ID,
	}
	initializers.DB.Create(&tellerUser)

	// Create a new teller with the user ID
	teller := models.Teller{
		UserID: tellerUser.ID,
	}
	initializers.DB.Create(&teller)

	helpers.RespondWithSuccess(c, http.StatusOK, "Teller created successfully", "200")
}
