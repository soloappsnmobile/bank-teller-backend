package helpers

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, message interface{}, respCode string) {
	c.AbortWithStatusJSON(code, gin.H{"resp_desc": message, "resp_code": respCode})
}

func RespondWithSuccess(c *gin.Context, code int, message interface{}, respCode string, data ...interface{}) {
	response := struct {
		RespCode string      `json:"resp_code"`
		RespDesc interface{} `json:"resp_desc"`
		Data     interface{} `json:"data,omitempty"`
	}{
		RespCode: respCode,
		RespDesc: message,
		Data:     nil,
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	c.JSON(code, response)
}
