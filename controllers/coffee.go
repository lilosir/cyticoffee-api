package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// GetAllCoffee will return the coffee
func GetAllCoffee(c *gin.Context) {
	allCoffee, err := models.GetAllCoffee()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, allCoffee)
}

// GetCoffee will return one coffee
func GetCoffee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	coffee, err := models.GetCoffee(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, coffee)
}
