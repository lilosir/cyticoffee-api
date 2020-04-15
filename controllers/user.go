package controllers

import (
	"database/sql"
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

	apiError := utils.NewAPIError(http.StatusInternalServerError, "server error", "")
	err := models.UserSignup(user)
	if err != nil {
		if err.Error() == "already exists" {
			apiError.Code = http.StatusConflict
			apiError.Message = "user already exists"
			c.Error(apiError)
			return
		}
		c.Error(err)
		return
	}

	data := serializers.UserSerializer(user)
	c.JSON(http.StatusAccepted, data)
}

// type login struct {
// 	email    string
// 	password string
// }

// LogIn handler
func LogIn(c *gin.Context) {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	apiErr := utils.NewAPIError(http.StatusBadRequest, "", nil)
	if err := c.ShouldBind(&login); err != nil {
		c.Error(err)
		return
	}

	result, err := models.UserLogIn(login.Email, login.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			apiErr.Code = http.StatusNotFound
			apiErr.Message = "Your email does not exist"
			c.Error(apiErr)
			return
		}
		if err.Error() == "Email and password do not match" {
			apiErr.Code = http.StatusUnauthorized
			apiErr.Message = err.Error()
			c.Error(apiErr)
			return
		}
		c.Error(err)
		return
	}

	data := serializers.UserSerializer(result)
	c.JSON(http.StatusOK, data)
}
