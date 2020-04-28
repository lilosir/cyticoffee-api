package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// GetAllSnacks will return all the snacks
func GetAllSnacks(c *gin.Context) {
	allSnacks, err := models.GetAllSnacks()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, allSnacks)
}

// GetSnack will return one snack
func GetSnack(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	snacks, err := models.GetSnacks(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, snacks)
}
