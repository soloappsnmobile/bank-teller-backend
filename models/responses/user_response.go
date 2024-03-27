package responses

import "bank-teller-backend/models"

// createUserResponse creates a response object containing the user's details.
func CreateUserResponse(user models.User, role models.Role) map[string]interface{} {
	return map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"phone":      user.Phone,
		"email":      user.Email,
		"role":       role.Name,
	}
}
