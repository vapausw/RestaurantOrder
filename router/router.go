package router

import (
	"RestaurantOrder/controller"
	"RestaurantOrder/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(mode string) *gin.Engine {
	if mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	r.POST("/user/login", controller.UserLoginHandler)
	r.POST("/user/send_code", controller.UserSendCodeHandler)
	r.POST("/user/register", controller.UserRegisterHandler)
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/info", controller.GetShopListHandler)          // 获取商店列表
		shopGroup.GET("/:id", controller.GetShopHandler)               // 获取商店详细信息
		shopGroup.GET("/:id/menu", controller.GetMenuListHandler)      // 获取某一个店家菜单列表
		shopGroup.GET("/:id/menu/:menu_id", controller.GetMenuHandler) // 获取某一个店家菜单详情
	}
	userGroup := r.Group("/user").Use(middleware.JWTAuthMiddleware())
	{
		// 思考一下用户购物的逻辑...,首先将其添加到购物车，然后去付款，两个接口
		// 也可以选择直接购买，默认直接添加购物车跳转到购买界面，一个接口
		// 购物车添加逻辑，前端直接发送购物车添加请求，后端直接添加到购物车
		// 购物车数据添加到那？redis，mysql
		// 或者开启一个定时任务，将购物车数据添加到mysql中
		// 添加用户添加商品到购物车的逻辑
		userGroup.POST("/cart/menu", controller.AddCartHandler)
		// 查看购物车的逻辑
		userGroup.GET("/cart/menu", controller.CartInfoHandler)
		// 使用购物车结款
		userGroup.POST("/cart/buy", controller.CartBuyHandler)
	}
	// 商户管理模块，主要就是对menu的控制
	r.POST("/merchant/login", controller.MerchantLoginHandler)
	r.POST("/merchant/send_code", controller.MerchantSendCodeHandler)
	r.POST("/merchant/register", controller.MerchantRegisterHandler)
	merchantGroup := r.Group("/merchant").Use(middleware.JWTAuthMiddleware())
	{
		// 商家信息的完善，不完善将不会展示给用户
		merchantGroup.POST("/info", controller.MerchantInfoHandler)
		// 商户对于自己的menu的CRUD操作
		merchantGroup.POST("/menu", controller.AddMenuHandler)
		merchantGroup.PUT("/menu", controller.UpdateMenuHandler)
		merchantGroup.DELETE("/menu", controller.DeleteMenuHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404 Not Found",
		})
	})
	return r
}
