package controller

import (
	"RestaurantOrder/dao"
	"RestaurantOrder/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GETMenuHandler 获取某个商家的所有所有菜品
func GETMenuHandler(c *gin.Context) {
	name := c.PostForm("name")
	// 从 Redis 获取按销量排序的商品ID
	productIDs, err := dao.Rdb.ZRevRange(fmt.Sprintf("%s_ids", name), 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product IDs from Redis"})
		return
	}

	// 根据商品ID从 MySQL 获取商品详细信息
	var products []models.Product
	if err := dao.DB.Where("id IN ?", productIDs).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products from DB"})
		return
	}

	// 根据 Redis 中的顺序重新排序 products
	sortedProducts := make([]models.Product, len(productIDs))
	idToProduct := make(map[string]models.Product)
	for _, product := range products {
		idToProduct[fmt.Sprint(product.ID)] = product
	}
	for i, id := range productIDs {
		sortedProducts[i] = idToProduct[id]
	}
	// 返回给前端
	c.JSON(http.StatusOK, gin.H{
		"message": "GET order",
		"status":  "success",
		"data":    sortedProducts,
	})
}

// CPOSTMenuHandler 客户点餐控制
func CPOSTMenuHandler(c *gin.Context) {
	var order []models.OrderItem
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 生成订单
	// 支付
	// 将订单信息存储
	c.JSON(http.StatusOK, gin.H{
		"message": "POST order",
		"status":  "success",
	})
}
