package middlewares

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/db/mysql"
	"github.com/lilosir/cyticoffee-api/utils"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// EmailClaims email is needed to claim with
type EmailClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Authenticate will check if token is valid
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header["Authorization"][0]
		newErr := utils.Unauthenticated

		if strings.Split(authHeader, " ")[0] != "Bearer" {
			c.Error(newErr)
			c.Abort()
			return
		}

		tknStr := strings.Split(authHeader, " ")[1]
		token, err := jwt.ParseWithClaims(tknStr, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.Error(newErr)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*EmailClaims); ok && token.Valid {
			// fmt.Printf("%d, %v", claims.ID, claims.Email)
			email, err := checkUser(claims.ID)
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}

			if email != claims.Email {
				c.Error(newErr)
				c.Abort()
				return
			}
			c.Set("reqUserId", claims.ID)
		} else {
			c.Error(newErr)
			c.Abort()
		}
		c.Next()
	}
}

func checkUser(id int64) (string, error) {
	newErr := utils.NewAPIError(http.StatusInternalServerError, "server error", nil)
	stmt, err := mysql.DBConn().Prepare("select email from tbl_user where id=?")
	if err != nil {
		fmt.Println(err)
		return "", newErr
	}

	defer stmt.Close()

	var email string
	err = stmt.QueryRow(id).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			newErr.Code = http.StatusBadRequest
			newErr.Message = "Your request is invalid"
			return "", newErr
		}
		return "", utils.ServerError
	}
	return email, nil
}
