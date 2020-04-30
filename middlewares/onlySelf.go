package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/utils"
)

// OnlySelf not only jwt valid, but also jwt user id equals url request user id
func OnlySelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("userID")
		reqID, ok := c.Get("reqUserId")
		if !ok {
			c.Error(utils.Unauthenticated)
			c.Abort()
			return
		}
		reqid := fmt.Sprintf("%v", reqID)
		if !ok {
			c.Error(utils.ServerError)
			c.Abort()
			return
		}
		if reqid != id {
			c.Error(utils.Unauthenticated)
			c.Abort()
			return
		}

		c.Next()
	}
}
