package middlewares

import (
	"bank-teller-backend/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			helpers.RespondWithError(c, http.StatusUnauthorized, "API token required", "401")
			c.Abort()
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			helpers.RespondWithError(c, http.StatusUnauthorized, "Authorization header format must be Bearer {token}", "401")
			c.Abort()
			return
		}

		// Remove "Bearer " from the token
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token and get the role
		role, err := helpers.ValidateTokenAndGetRole(token)
		if err != nil || role != "Admin" {
			helpers.RespondWithError(c, http.StatusUnauthorized, err.Error(), "401")
			c.Abort()
			return
		}

		// Save the role in the context
		c.Set("role", role)

		c.Next()
	}
}
