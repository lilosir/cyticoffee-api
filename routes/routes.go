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
	r.POST("/login", controllers.LogIn)

	r.GET("/type", middlewares.Authenticate(), controllers.GetType)
	r.GET("/coffee", middlewares.Authenticate(), controllers.GetAllCoffee)
	r.GET("/coffee/:id", middlewares.Authenticate(), controllers.GetCoffee)
	r.GET("/tea", middlewares.Authenticate(), controllers.GetAllTea)
	r.GET("/tea/:id", middlewares.Authenticate(), controllers.GetTea)
	r.GET("/otherdrinks", middlewares.Authenticate(), controllers.GetAllOtherDrinks)
	r.GET("/otherdrinks/:id", middlewares.Authenticate(), controllers.GetOtherDrink)
	r.GET("/snacks", middlewares.Authenticate(), controllers.GetAllSnacks)
	r.GET("/snacks/:id", middlewares.Authenticate(), controllers.GetSnack)
	r.GET("/food", middlewares.Authenticate(), controllers.GetAllFood)
	r.GET("/food/:id", middlewares.Authenticate(), controllers.GetFood)

	return r
}
