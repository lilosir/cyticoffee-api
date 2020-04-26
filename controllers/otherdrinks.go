package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// GetAllOtherDrinks will return the all other drinks
func GetAllOtherDrinks(c *gin.Context) {
	allCoffee, err := models.GetAllOtherDrinks()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, allCoffee)
}

// GetOtherDrinks will return one other drinks
func GetOtherDrinks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	otherDrinks, err := models.GetOtherDrinks(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, otherDrinks)
}
