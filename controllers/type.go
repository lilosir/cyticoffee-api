package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// GetType will return the menu
func GetType(c *gin.Context) {
	types, err := models.GetType()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, types)
}
