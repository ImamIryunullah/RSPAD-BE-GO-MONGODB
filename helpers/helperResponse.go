package helper

import (
	"github.com/gin-gonic/gin"
)

func ResponseSucces(c *gin.Context, status_code int, Message string, data interface{}) {
	c.JSON(status_code, gin.H{"status_code": status_code, "message": Message, "data": data})
}
func ErrorResponse(c *gin.Context, status_code int, Message string, err string) {
	c.AbortWithStatusJSON(status_code, gin.H{"status_code": status_code, "message": Message, "error": err})
}
