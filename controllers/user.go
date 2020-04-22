package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
	"github.com/lilosir/cyticoffee-api/serializers"
	"github.com/lilosir/cyticoffee-api/utils"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// EmailClaims email is needed to claim with
type EmailClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateJWT(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &EmailClaims{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 365 * time.Hour).Unix(),
		},
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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

	tokenString, err := generateJWT(user.ID, user.Email)
	if err != nil {
		c.Error(err)
		return
	}
	c.Header("auth", tokenString)

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

	tokenString, err := generateJWT(result.ID, result.Email)
	if err != nil {
		c.Error(err)
		return
	}
	c.Header("auth", tokenString)

	data := serializers.UserSerializer(result)
	c.JSON(http.StatusOK, data)
}
