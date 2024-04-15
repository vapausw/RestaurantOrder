package controller

import (
	"RestaurantOrder/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 商家CRUD操作

// MPOSTMenuHandler 添加菜品
func MPOSTMenuHandler(c *gin.Context) {
	email := c.PostForm("email")
	var p models.Product
	c.BindJSON(&p)
	p.MerchantEmail = email
	if err := models.CreateProduct(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "POST menu", "status": "success"})
	}
}

// MPUTMenuHandler 更新菜品
func MPUTMenuHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	p, err := models.GetProductByID(i)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&p)
	if err = models.UpdateProduct(p); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "PUT menu", "status": "success"})
	}
}

// MDELETEMenuHandler 删除菜品
func MDELETEMenuHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	if err := models.DeleteProductByID(i); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "DELETE menu", "status": "success"})
	}
}
