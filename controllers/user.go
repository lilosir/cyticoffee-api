package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
	"github.com/lilosir/cyticoffee-api/utils"
)

//SignUp handler
func SignUp(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")

	password = utils.CreateSha1([]byte(password))

	user := &models.User{
		Email:     email,
		Password:  password,
		Phone:     phone,
		Firstname: firstname,
		Lastname:  lastname,
	}

	respBody := utils.NewRespMes("", nil)
	id, err := models.UserSignup(user)
	if err != nil {
		respBody.Message = err.Error()
		code := http.StatusInternalServerError
		if respBody.Message == "already exists" {
			code = http.StatusConflict
		}
		c.JSON(code, respBody)
		return
	}
	respBody.Message = "ok"
	user.Id = id
	respBody.Data = user
	c.JSON(http.StatusAccepted, respBody)
}
