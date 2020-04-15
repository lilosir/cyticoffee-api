package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lilosir/cyticoffee-api/utils"
)

func setErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is requred", err.Field())
	case "email":
		return fmt.Sprintf("Invalid email")
	default:
		return fmt.Sprintf("%s is not valid", err.Field())
	}
}

// Errors is a middleware to handle all the errors
func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// Only run if there are some errors to handle
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				switch e.Type {
				case gin.ErrorTypeBind:
					return
				case gin.ErrorTypeRender:
					fmt.Println("gin.ErrorTypeRender")
					return
				case gin.ErrorTypePrivate:
					fmt.Println("gin.ErrorTypePrivate")
					errs, ok := e.Err.(validator.ValidationErrors)
					apiError := utils.NewAPIError(http.StatusBadRequest, "", nil)
					if ok {
						errSlice := make([]string, len(errs))
						for i, err := range errs {
							errSlice[i] = setErrorMessage(err)

						}
						// Make sure we maintain the preset response status
						// status := http.StatusBadRequest
						// if c.Writer.Status() != http.StatusOK {
						// 	status = c.Writer.Status()
						// }
						apiError.Data = errSlice
						apiError.Message = "Validation Error"
						c.JSON(apiError.Code, apiError)
						return
					}
					err, ok := e.Err.(*utils.APIError)
					if ok {
						apiError.Code = err.Code
						apiError.Message = err.Error()
						c.JSON(apiError.Code, apiError)
						return
					}
					c.JSON(http.StatusInternalServerError, utils.ServerError)
					return
				case gin.ErrorTypePublic:
					fmt.Println("gin.ErrorTypePublic")
					return
				case gin.ErrorTypeAny:
					fmt.Println("gin.ErrorTypeAny")
					return
				default:
					if !c.Writer.Written() {
						c.JSON(c.Writer.Status(), utils.ServerError)
					}
				}

			}
		}
	}
}
