package helpers

import "github.com/gin-gonic/gin"

func RespondWithError(c *gin.Context, code int, message interface{}, respCode string) {
	c.AbortWithStatusJSON(code, gin.H{"resp_desc": message, "resp_code": respCode})
}

func RespondWithSuccess(c *gin.Context, code int, message interface{}, respCode string, data ...interface{}) {
	if len(data) > 0 {
		c.JSON(code, gin.H{"resp_desc": message, "resp_code": respCode, "data": data[0]})
		return
	}
	c.JSON(code, gin.H{"resp_desc": message, "resp_code": respCode})
}
