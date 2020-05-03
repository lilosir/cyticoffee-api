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
	userID, _ := c.Get("reqUserId")
	err := models.CreateOrders(order.Order, userID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, order)
}

// GetMyOrders return my orders
func GetMyOrders(c *gin.Context) {
	id := c.Param("userID")
	orders, err := models.GetMyOrders(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, orders)
}

// GetOrderDetails get one order details
func GetOrderDetails(c *gin.Context) {
	id := c.Param("orderID")
	orderDetail, err := models.GetOrderDetails(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, orderDetail)
}
