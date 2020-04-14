package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Authenticate will check if token is valid
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("should Authenticate token")
		c.Next()
	}
}
