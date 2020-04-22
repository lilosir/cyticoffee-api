package controllers

import (
	"github.com/gin-gonic/gin"
)

// GetAllOtherDrinks will return all other drinks
func GetAllOtherDrinks(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}

// GetOtherDrink will return one other drink
func GetOtherDrink(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hi",
	})
}
