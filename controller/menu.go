package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"RestaurantOrder/model"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddMenuHandler(c *gin.Context) {
	// 获取当前商家的id
	merchantId, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		zap.L().Error("AddMenuHandler with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	zap.L().Info("AddMenuHandler", zap.Any("menu", menu))
	menu.ShopID = int(merchantId)
	if err, errCode := logic.AddMenu(&menu); err != nil {
		zap.L().Error("AddMenuHandler with invalid param", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
	}
	responseSuccess(c, nil)
}

func UpdateMenuHandler(c *gin.Context) {

}

func DeleteMenuHandler(c *gin.Context) {

}
