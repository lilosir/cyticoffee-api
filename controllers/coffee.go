package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAllCoffee will return the coffee
func GetAllCoffee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

// GetCoffee will return one coffee
func GetCoffee(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
