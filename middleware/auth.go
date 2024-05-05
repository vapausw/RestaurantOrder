package middleware

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"strings"
	"time"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			responseErrorWithMsg(c, co.CodeInvalidAuthFormat, "请求头缺少Auth Token")
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			responseErrorWithMsg(c, co.CodeInvalidAuthFormat, "Token格式不对")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			responseErrorWithMsg(c, co.CodeInvalidToken, "无效的Token")
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(co.ContextUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("user_id")来获取当前请求的用户信息
	}
}
func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
