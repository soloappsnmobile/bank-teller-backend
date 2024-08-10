package admin

import (
	"bank-teller-backend/helpers"
	"bank-teller-backend/initializers"
	"bank-teller-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTellers(c *gin.Context) {
	var tellers []models.Teller
	result := initializers.DB.Find(&tellers)

	if result.Error != nil {
		helpers.RespondWithError(c, http.StatusBadRequest, result.Error, "01")
		return
	}

	// If there's no error, respond with a success message, a status code of 200 (OK),
	// and the addressBook slice containing the retrieved records.
	helpers.RespondWithSuccess(c, http.StatusOK, "Tellers retrieved successfully", "00", tellers)
}
