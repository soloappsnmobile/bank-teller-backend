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

// func CreateCustomer(c *gin.Context) {
// 	var req map[string]interface{}
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request data", "400")
// 		return
// 	}

// 	// Validate the request
// 	if err := validators.ValidateRequest(req, "Customer"); err != nil {
// 		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "400")
// 		return
// 	}
// 	customerRole := models.Role{Name: "Customer"}
// 	// Hash the password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req["password"].(string)), bcrypt.DefaultCost)
// 	if err != nil {
// 		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password", "500")
// 		return
// 	}
// 	// Create a new user with the teller role
// 	customerUser := models.User{
// 		FirstName: req["first_name"].(string),
// 		LastName:  req["last_name"].(string),
// 		Phone:     req["phone"].(string),
// 		Email:     req["email"].(string),
// 		Password:  string(hashedPassword),
// 		RoleID:    customerRole.ID,
// 	}
// 	initializers.DB.Create(&customerUser)

// 	// Create a new teller with the user ID
// 	customer := models.Teller{
// 		UserID: customerUser.ID,
// 	}
// 	initializers.DB.Create(&customer)

//		helpers.RespondWithSuccess(c, http.StatusOK, "Customer created successfully", "200")
//	}
func CreateCustomer(c *gin.Context) {
	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, "Invalid request data", "400")
		return
	}

	// Validate the request
	if err := validators.ValidateRequest(req, "Customer"); err != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, err.Error(), "400")
		return
	}

	// Get the customer role by name
	customerRole, err := models.GetRoleByName(initializers.DB, "Customer")
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to get role", "500")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to hash password", "500")
		return
	}

	// Create a new user with the customer role
	customerUser := models.User{
		FirstName: req["first_name"].(string),
		LastName:  req["last_name"].(string),
		Phone:     req["phone"].(string),
		Email:     req["email"].(string),
		Password:  string(hashedPassword),
		RoleID:    customerRole.ID,
	}
	// fmt.Printf("Customer User: %+v\n", customerUser)
	if err := initializers.DB.Create(&customerUser).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create user", "500")
		return
	}

	// Create a new teller with the user ID
	customer := models.Customer{
		UserID: customerUser.ID,
	}
	if err := initializers.DB.Create(&customer).Error; err != nil {
		helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create customer", "500")
		return
	}

	// if err := initializers.DB.Create(&customer).Error; err != nil {
	// 	helpers.RespondWithError(c, http.StatusInternalServerError, "Failed to create customer", "500")
	// 	return
	// }

	helpers.RespondWithSuccess(c, http.StatusOK, "Customer created successfully", "200")
}
