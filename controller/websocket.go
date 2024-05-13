package controller

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/jwt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients     = map[int64]*websocket.Conn{}
	clientsLock = sync.RWMutex{}
)

func WebSocketHandler(c *gin.Context) {
	token := c.Query("token")

	var userID int64
	if len(token) != 0 {
		// 通过token获取用户信息
		// 通过用户信息获取用户ID
		// 按空格分割
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			responseErrorWithMsg(c, co.CodeInvalidAuthFormat, "Token格式不对")
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			responseErrorWithMsg(c, co.CodeInvalidToken, "无效的Token")
			return
		}
		userID = mc.UserID
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.L().Error("upgrade to websocket failed", zap.Error(err))
		responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
	}
	// 将用户的连接信息存储到map中
	// 一个连接只能对应一个用户
	if userID != 0 && clients[userID] == nil {
		// 加锁
		clientsLock.Lock()
		if clients[userID] == nil {
			clients[userID] = ws
		}
		// 解锁
		clientsLock.Unlock()
	}
	for {
		// 在此处读取消息，并根据需要处理消息
		msgType, message, err := ws.ReadMessage()
		if err != nil {
			zap.L().Error("read message from websocket failed", zap.Error(err))
			responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
			break
		}

		// 反序列化
		var msg model.UserWebsocket
		err = json.Unmarshal(message, &msg)
		if err != nil {
			zap.L().Error("unmarshal message failed", zap.Error(err))
			responseErrorWithMsg(c, co.CodeInvalidParam, co.ErrInvalidParam.Error())
			break
		}
		// 发送给对应的用户
		clientsLock.RLock()
		client, ok := clients[msg.UserId]
		clientsLock.RUnlock()
		if ok {
			err = client.WriteMessage(msgType, message)
			if err != nil {
				zap.L().Error("write message to websocket failed", zap.Error(err))
				responseErrorWithMsg(c, co.CodeServerBusy, co.ErrServerBusy.Error())
				break
			}
		}
	}
	defer func() {
		if userID != 0 {
			clientsLock.Lock() // 加写锁
			delete(clients, userID)
			clientsLock.Unlock() // 解锁
		}
		ws.Close()
	}()
}
