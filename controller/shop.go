package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// GetShopListHandler 获取商店信息
func GetShopListHandler(c *gin.Context) {
	shops, err, code := logic.GetShopList()
	if err != nil {
		responseErrorWithMsg(c, code, err.Error())
	}
	responseSuccess(c, shops)
}

// GetShopHandler 根据shop_id获取商店信息
func GetShopHandler(c *gin.Context) {
	id := c.Param("id")
	shopId, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("invalid shop_id", zap.String("id", id))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
	}
	shop, err, code := logic.GetShop(shopId)
	if err != nil {
		responseErrorWithMsg(c, code, err.Error())
		return
	}
	responseSuccess(c, shop)
}

func GetMenuListHandler(c *gin.Context) {
	id := c.Param("id")
	shopId, err := strconv.Atoi(id)
	//zap.L().Info("shop_id", zap.Any("shop_id", shopId))
	if err != nil {
		zap.L().Error("invalid shop_id", zap.String("id", id))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
	}
	menuList, err, code := logic.GetShopMenuList(shopId)
	if err != nil {
		responseErrorWithMsg(c, code, err.Error())
		return
	}
	responseSuccess(c, menuList)
}

func GetMenuHandler(c *gin.Context) {
	id := c.Param("id")
	menuId := c.Param("menu_id")
	zap.L().Info("menu_id", zap.Any("menu_id", menuId))
	shopId, err := strconv.Atoi(id)
	if err != nil {
		zap.L().Error("invalid shop_id", zap.String("id", id))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
	}
	menuIdInt, err := strconv.Atoi(menuId)
	if err != nil {
		zap.L().Error("invalid menu_id", zap.String("menu_id", menuId))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
	}
	menu, err, code := logic.GetShopMenu(shopId, menuIdInt)
	if err != nil {
		responseErrorWithMsg(c, code, err.Error())
		return
	}
	responseSuccess(c, menu)
}
