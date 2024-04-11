package controller

import (
	"RestaurantOrder/dao"
	"RestaurantOrder/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GETMenuHandler(c *gin.Context) {
	// 从 Redis 获取按销量排序的商品ID
	productIDs, err := dao.Rdb.ZRevRange("product_sales", 0, -1).Result()
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

func POSTMenuHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST order",
		"status":  "success",
	})
}