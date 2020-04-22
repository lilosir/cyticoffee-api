package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAllTea will return the tea
func GetAllTea(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

// GetTea will return one coffee
func GetTea(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
