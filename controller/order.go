package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetOrderListHandler(c *gin.Context) {
	// 获取用户id
	// 获取用户id
	user_id, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	// 调用logic层的函数
	data, err, errCode := logic.GetUserOrderList(user_id)
	if err != nil {
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	// 返回响应
	responseSuccess(c, data)
}
