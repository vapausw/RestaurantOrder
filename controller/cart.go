package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"RestaurantOrder/model"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func AddCartHandler(c *gin.Context) {
	// 获取用户id
	user_id, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	// 获取所要添加商品的信息
	var info model.CartInfo
	if err := c.ShouldBind(&info); err != nil {
		zap.L().Error("AddCartHandler with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	zap.L().Info("AddCartHandler", zap.Any("info", info))
	// 调用logic层的添加购物车的函数
	if err, errCode := logic.AddCart(user_id, &info); err != nil {
		zap.L().Error("logic.AddCart() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	// 返回响应
	responseSuccess(c, nil)
}

func CartInfoHandler(c *gin.Context) {
	// 根据用户的id去查询购物车的数据
	// 获取用户id
	user_id, err := getCurrentID(c)
	zap.L().Info("user_id", zap.Any("user_id", user_id))
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	data, err, errCode := logic.GetCartInfo(user_id)
	if err != nil {
		zap.L().Error("logic.GetCartInfo() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	responseSuccess(c, data)
}

func CartBuyHandler(c *gin.Context) {
	// 获取用户id
	user_id, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	// 获取前端发送的购物车的信息
	// 获取前端发送的购物车的信息
	var info []*model.CartInfo                      // 使用指针的切片来接收数组
	if err := c.ShouldBindJSON(&info); err != nil { // 注意这里是 &infos，传递的是切片的指针
		zap.L().Error("AddCartHandler with invalid param", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	zap.L().Info("CartBuyHandler", zap.Any("info", info))
	// 将用户id添加到购物车信息中
	for i := range info {
		info[i].UserID = user_id
	}
	// 调用logic层的购买函数
	if err, errCode := logic.CartBuy(info); err != nil {
		zap.L().Error("logic.CartBuy() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	// 返回响应
	responseSuccess(c, nil)
}

func UpdateCartHandler(c *gin.Context) {
	// 获取用户id
	user_id, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	// 获取所要添加商品的信息
	var info model.CartInfo
	if err := c.ShouldBind(&info); err != nil {
		zap.L().Error("AddCartHandler with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	info.UserID = user_id
	if err, errCode := logic.UpdateCart(info); err != nil {
		responseErrorWithMsg(c, errCode, err.Error())
	}
	responseSuccess(c, nil)
}

func DeleteCartHandler(c *gin.Context) {
	// 获取用户id
	user_id, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	menu_id := c.Param("menu_id")
	menuId, _ := strconv.Atoi(menu_id)
	if err, errCode := logic.DeleteCart(user_id, int64(menuId)); err != nil {
		responseErrorWithMsg(c, errCode, err.Error())
	}
	responseSuccess(c, nil)
}
