package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/logic"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

func UserLoginHandler(c *gin.Context) {
	var u model.User
	if err := c.ShouldBind(&u); err != nil {
		zap.L().Error("UserLogin with invalid param", zap.Error(err))
		responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
		return
	}
	zap.L().Info("em", zap.Any("em", u.Email))
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
	zap.L().Info("em", zap.Any("em", u))
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

func RefreshTokenHandler(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	rt := c.Query("refresh_token")
	zap.L().Info("refresh token", zap.String("rt", rt))
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		responseErrorWithMsg(c, co.CodeInvalidToken, co.ErrFormat.Error())
		c.Abort()
		return
	}
	// 按空格分割
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		responseErrorWithMsg(c, co.CodeInvalidToken, co.ErrFormat.Error())
		c.Abort()
		return
	}
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	fmt.Println(err)
	responseSuccess(c, gin.H{
		"accessToken":  aToken,
		"refreshToken": rToken,
	})
}

func GetUserInfoHandler(c *gin.Context) {
	user_id, err := getCurrentID(c)
	if err != nil {
		if errors.Is(err, co.ErrUserNotLogin) {
			responseErrorWithMsg(c, co.CodeUserNotLogin, co.ErrUserNotLogin.Error())
			return
		}
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	// 获取用户信息
	user, err, errCode := logic.GetUserInfo(user_id)
	if err != nil {
		zap.L().Error("logic.GetUserInfo() failed", zap.Error(err))
		responseErrorWithMsg(c, errCode, err.Error())
		return
	}
	responseSuccess(c, user)
}
