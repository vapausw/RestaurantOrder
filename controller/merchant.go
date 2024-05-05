package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MerchantLoginHandler(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBind(&merchant); err != nil {
		zap.L().Error("MerchantLogin with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	if err, errCode := logic.MerchantLogin(&merchant); err != nil {
		responseErrorWithMsg(c, errCode, err.Error())
	}
	//登录成功返回token
	//生成Token
	aToken, rToken, _ := jwt.GenToken(merchant.MerchantId)
	responseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
	})
}

func MerchantSendCodeHandler(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	if err, errCode := logic.SendMerchantCode(email); err != nil {
		responseErrorWithMsg(c, errCode, err.Error())
	}
	responseSuccess(c, nil)
}

func MerchantRegisterHandler(c *gin.Context) {
	var merchant model.Merchant
	if err := c.ShouldBind(&merchant); err != nil {
		zap.L().Error("MerchantLogin with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	// 业务逻辑
	if err, errCode := logic.MerchantRegister(&merchant); err != nil {
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	// 返回响应
	responseSuccess(c, nil)
}

func MerchantInfoHandler(c *gin.Context) {
	merchantId, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	var shop model.Shop
	shop.ShopId = int(merchantId)
	if err := c.ShouldBindJSON(&shop); err != nil {
		zap.L().Error("MerchantInfoHandler with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	if err, errCode := logic.MerchantInfo(&shop); err != nil {
		zap.L().Error("logic.MerchantInfo() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	responseSuccess(c, nil)
}
