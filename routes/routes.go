package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/controllers"
	"github.com/lilosir/cyticoffee-api/middlewares"
)

// SetupRoutes creates gin engin and routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Errors())
	r.GET("/", controllers.Index)
	r.POST("/signup", controllers.SignUp)

	return r
}
