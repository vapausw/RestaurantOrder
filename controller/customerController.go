package controller

import (
	"RestaurantOrder/logic"
	"RestaurantOrder/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GETLoginHandler 唤醒登录页面，此处为GET请求，且一般都是c.HTML,此处只用于后端测试
func GETLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":    "login get",
		"status": "200",
	})
}

// POSTLoginHandler 登录页面，此处为POST请求
func POSTLoginHandler(c *gin.Context) {
	type User struct {
		Username string `json:"email"`
		Password string `json:"password"`
	}
	var u User
	c.BindJSON(&u)
	err := logic.LoginCheck(u.Username, u.Password, "customer")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"error":  err.Error(),
		})
		return
	}
	// 此处调用登录逻辑进行判断
	c.JSON(http.StatusOK, gin.H{
		"msg":      "login success",
		"username": u.Username,
		// 可选：告诉前端重定向到哪里
		"redirectURL": "/user/index",
		"status":      "200",
	})
	// 登录验证成功进行重定向到主页面
}

// GETRegisterHandler 注册的GET请求挂与login下
func GETRegisterHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":    "register get",
		"status": "200",
	})
}

// POSTRegisterHandler 注册的POST请求挂与login下
func POSTRegisterHandler(c *gin.Context) {
	type Register struct {
		Action         string `json:"action"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		PassWord       string `json:"password"`
		RepeatPassWord string `json:"repeatpassword"`
		Captcha        string `json:"captcha"`
	}
	var r Register
	c.BindJSON(&r)
	switch r.Action {
	case "sendCode":
		// 调用logic层的处理函数
		err := logic.RegisterSendCode(r.Email)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "验证码发送成功",
				"status": "200",
				"email":  r.Email,
			})
		}
	case "register":
		var v models.Customer
		v.Email = r.Email
		v.Password = r.PassWord
		v.Name = r.Name
		err := logic.RegisterCheck(r.RepeatPassWord, r.Captcha, v)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":      "注册成功，即将自动登录",
				"status":   "200",
				"Register": r,
			})
		}
	default:
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的action",
		})
	}
}

//是否要添加一个登录的验证中间件，方便后续请求拿到登录的信息

func GETIndexHandler(c *gin.Context) {
	// 唤醒前端页面
	c.HTML(http.StatusOK, "index.html", nil)
}
