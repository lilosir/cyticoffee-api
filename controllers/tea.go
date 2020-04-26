package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// GetAllTea will return the tea
func GetAllTea(c *gin.Context) {
	allTea, err := models.GetAllTea()
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, allTea)
}

// GetTea will return one coffee
func GetTea(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}
	tea, err := models.GetTea(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, tea)
}
