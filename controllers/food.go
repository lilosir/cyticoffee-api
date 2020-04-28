package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// GetAllFood will return the food
func GetAllFood(c *gin.Context) {
	allFood, err := models.GetAllFood()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, allFood)
}

// GetFood will return one food
func GetFood(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	food, err := models.GetFood(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, food)
}
