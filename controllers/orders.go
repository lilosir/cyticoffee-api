package controllers

import "github.com/gin-gonic/gin"

// CreateOrders add a new order
func CreateOrders(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create a new order",
	})
}

// GetOrderDetails get one order details
func GetOrderDetails(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is order details",
	})
}
