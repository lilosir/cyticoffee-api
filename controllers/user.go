package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
	"github.com/lilosir/cyticoffee-api/serializers"
	"github.com/lilosir/cyticoffee-api/utils"
)

//SignUp handler
func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); err != nil {
		c.Error(err)
		return
	}
	user.Password = utils.CreateSha1([]byte(user.Password))

	err := models.UserSignup(user)
	if err != nil {
		c.Error(err)
		return
	}

	data := serializers.UserSerializer(user)
	c.JSON(http.StatusAccepted, data)
}

// LogIn handler
func LogIn(c *gin.Context) {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&login); err != nil {
		c.Error(err)
		return
	}

	result, err := models.UserLogIn(login.Email, login.Password)
	if err != nil {
		c.Error(err)
		return
	}

	data := serializers.UserSerializer(result)
	c.JSON(http.StatusOK, data)
}
