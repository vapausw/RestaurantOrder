package controller

import (
	"RestaurantOrder/models"
	"github.com/gin-gonic/gin"
)

// GETOrderHandler 根据邮箱查看该用户的所有订单
func GETOrderHandler(c *gin.Context) {
	var orders []models.Order
	orders, err := models.GetOrdersByCustomerEmail(c.Query("customer_email"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"orders": orders,
	})
}
