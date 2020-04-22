package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAllSnacks will return all the snacks
func GetAllSnacks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

// GetSnack will return one snack
func GetSnack(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
