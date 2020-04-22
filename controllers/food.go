package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAllFood will return all the food
func GetAllFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

// GetFood will return one food
func GetFood(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
