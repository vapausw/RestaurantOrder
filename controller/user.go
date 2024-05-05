package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserLoginHandler(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		zap.L().Error("UserLogin with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	//zap.L().Info("em", zap.Any("em", u.Email))
	if err, errCode := logic.Login(&u); err != nil {
		zap.L().Error("logic.Login() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	//登录成功返回token
	//生成Token
	aToken, rToken, _ := jwt.GenToken(int64(u.UserId))
	responseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
	})
}

// UserSendCodeHandler 发送验证码
func UserSendCodeHandler(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	if err, errCode := logic.SendCode(email); err != nil {

		responseErrorWithMsg(c, errCode, err.Error())
	}
	responseSuccess(c, nil)
}

// UserRegisterHandler 用户注册
func UserRegisterHandler(c *gin.Context) {
	// 获取请求参数
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		zap.L().Error("UserLogin with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	// 业务逻辑
	if err, errCode := logic.Register(&u); err != nil {
		zap.L().Error("logic.Register() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	// 返回响应
	responseSuccess(c, nil)
}

func getCurrentID(c *gin.Context) (userID int64, err error) {
	_userID, ok := c.Get(co.ContextUserIDKey)
	if !ok {
		err = co.ErrUserNotLogin
		return
	}
	userID, ok = _userID.(int64)
	if !ok {
		err = co.ErrServerBusy
		return
	}
	return
}
