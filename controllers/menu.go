package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetMenu will return the menu
func GetMenu(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
