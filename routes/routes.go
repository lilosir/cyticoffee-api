package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/controllers"
)

// SetupRoutes creates gin engin and routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/concurrency", controllers.Concurrency)

	return r
}
