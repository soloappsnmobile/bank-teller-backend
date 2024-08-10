package auth

import (
	"bank-teller-backend/helpers"
	"bank-teller-backend/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	_, user, role, tokenString, err := helpers.ProcessLogin(c)
	if err != nil {
		return
	}

	userResponse := responses.CreateUserResponse(user, role)

	helpers.RespondWithSuccess(c, http.StatusOK, "Login successful", "00", gin.H{
		"user":  userResponse,
		"token": tokenString,
	})
}
