package helpers

import (
	"bank-teller-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessLogin(c *gin.Context) (req map[string]interface{}, user models.User, role models.Role, tokenString string, err error) {
	req, err = ParseRequest(c)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "Invalid request data", "400")
		return
	}

	user, err = GetUser(req["email"].(string))
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, "Invalid email", "401")
		return
	}

	if err = CheckPassword(user, req["password"].(string)); err != nil {
		RespondWithError(c, http.StatusUnauthorized, "Invalid password", "401")
		return
	}

	role, err = GetRole(user)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Could not find role", "500")
		return
	}

	tokenString, err = GenerateToken(user, role)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "Could not generate token", "500")
		return
	}

	return
}
