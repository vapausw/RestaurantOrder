package router

import (
	"RestaurantOrder/controller"
	"RestaurantOrder/log"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(log.GinLogger(), log.GinRecovery(true))
	// 静态文件处理
	r.Static("/static", "static")
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	// 分组，分为用户以及商户两组
	// 用户组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", controller.GETLoginHandler)         // 登录
		userGroup.POST("/login", controller.POSTLoginHandler)       // 登录验证
		userGroup.GET("/register", controller.GETRegisterHandler)   // 注册
		userGroup.POST("/register", controller.POSTRegisterHandler) // 注册流程
		userindex := userGroup.Group("/index")
		{
			userindex.GET("/", controller.GETIndexHandler)
			userindex.GET("/menu", controller.GETMenuHandler)  // 查看菜单
			userindex.POST("menu", controller.POSTMenuHandler) // 点餐
			userindex.GET("/order")                            // 查看订单
		}
	}
	// 商户组
	merchantGroup := r.Group("/merchant")
	{
		merchantGroup.GET("/login")
	}
	return r
}
