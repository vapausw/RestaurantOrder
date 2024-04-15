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
		userIndex := userGroup.Group("/index")
		{
			userIndex.GET("/", controller.GETIndexHandler)
			userIndex.GET("/menu", controller.GETMenuHandler)   // 查看菜单
			userIndex.POST("menu", controller.CPOSTMenuHandler) // 点餐
			userIndex.GET("/order")                             // 查看订单
		}
	}
	// 商户组
	merchantGroup := r.Group("/merchant")
	{
		merchantGroup.GET("/login", controller.GETLoginHandler)
		merchantGroup.POST("/login", controller.POSTLoginHandler)       // 登录验证
		merchantGroup.GET("/register", controller.GETRegisterHandler)   // 注册
		merchantGroup.POST("/register", controller.POSTRegisterHandler) // 注册流程
		merchantIndex := merchantGroup.Group("/index")
		{
			merchantIndex.GET("/", controller.GETIndexHandler)
			merchantIndex.GET("/menu", controller.GETMenuHandler)            // 查看所有菜品
			merchantIndex.POST("menu", controller.MPOSTMenuHandler)          // 添加菜品
			merchantIndex.PUT("/menu/:id", controller.MPUTMenuHandler)       // 更新菜品
			merchantIndex.DELETE("/menu/:id", controller.MDELETEMenuHandler) // 删除菜品
			merchantIndex.GET("/order")                                      // 查看所有订单信息
		}
	}
	return r
}
