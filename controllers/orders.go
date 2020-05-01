package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilosir/cyticoffee-api/models"
)

// Order struct
type Order struct {
	Order []models.OrderItem
}

// CreateOrders add a new order
func CreateOrders(c *gin.Context) {
	var order Order
	if err := c.ShouldBind(&order); err != nil {
		c.Error(err)
		return
	}
	err := models.CreateOrders(order.Order)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, order)
}

// GetOrderDetails get one order details
func GetOrderDetails(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is order details",
	})
}
