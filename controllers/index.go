package controllers

import (
	"github.com/gin-gonic/gin"
)

// Index is index
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
