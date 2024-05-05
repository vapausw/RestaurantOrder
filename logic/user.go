package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/cache"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/kafka"
	"RestaurantOrder/pkg/snowflake"
	"RestaurantOrder/pkg/util"
	"RestaurantOrder/setting"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"strings"
)

func Login(u *model.User) (err error, errCode co.MyCode) {
	//zap.L().Info("em", zap.Any("em", u.Email))
	// 验证邮箱格式是否有误
	if !util.ValidateEmail(u.Email) {
		return co.ErrEmailFormat, co.CodeInvalidParam
	}
	// 通过邮箱获取用户加密的密码
	cacheUser, err := cache.GetUserPassword(u.Email)
	if err != nil {
		return co.ErrServerBusy, co.CodeServerBusy
	}
	if cacheUser.Password == co.BadData {
		return co.ErrNotExists, co.CodeNotFound
	}
	if !util.CheckPasswordHash(u.Password, cacheUser.Password) {
		return co.ErrPassword, co.CodeErrPassword
	}
	u.UserId = cacheUser.UserId
	return nil, co.CodeSuccess
}

func SendCode(email string) (err error, errCode co.MyCode) {
	// 验证邮箱格式是否有误
	if !util.ValidateEmail(email) {
		return co.ErrEmailFormat, co.CodeInvalidParam
	}
	// 生成验证码
	code := util.GenValidateCode()
	// 将验证码存入redis
	if err := redis.SetUserCode(email, code); err != nil {
		zap.L().Error("redis.SetUserCode failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	// 发送验证码，使用kafka异步发送
	zap.L().Info("code", zap.String("code", code))
	if err := sendEmail(email, tokenBody(code)); err != nil {
		zap.L().Error("sendEmail token failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}

func Register(u *model.User) (err error, errCode co.MyCode) {
	// 验证邮箱格式是否有误
	if !util.ValidateEmail(u.Email) {
		return co.ErrEmailFormat, co.CodeInvalidParam
	}
	// 验证码密码是否一致
	if u.Password != u.RePassword {
		return co.ErrPassword, co.CodeErrPassword
	}
	// 验证码是否正确
	code, err := redis.GetUserCode(u.Email)
	if err != nil {
		if errors.Is(err, co.ErrNotFound) {
			return co.ErrAuthCode, co.CodeErrCode
		}
		return co.ErrServerBusy, co.CodeServerBusy
	}
	if code != u.Code {
		return co.ErrAuthCode, co.CodeErrCode
	}
	// 将用户存储到mysql数据库中
	u.Password, err = util.HashPassword(u.Password)
	if err != nil {
		zap.L().Error("util.HashPassword failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	u.UserName = util.GenerateRandomNickname()
	u.UserId = int(snowflake.GenID())
	if err := mysql.CreateUser(u); err != nil {
		if errors.Is(err, co.ErrExistsUser) {
			return co.ErrExistsUser, co.CodeExistsCode
		}
		return co.ErrServerBusy, co.CodeServerBusy
	}
	// 异步发送欢迎邮件
	if err := sendEmail(u.Email, welcomeBody()); err != nil {
		zap.L().Error("sendEmail welcome failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}
func emailHeader(email string) string {
	from := setting.Conf.MyEmailConfig.Email
	var header strings.Builder
	header.WriteString(strings.Join([]string{"From: ", from, "\r\n"}, " "))
	header.WriteString(strings.Join([]string{"To: ", email, "\r\n"}, " "))
	header.WriteString("Subject: Verify Your Email\r\n")
	header.WriteString("MIME-Version: 1.0\r\n")
	header.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
	return header.String()
}

func tokenBody(token string) string {
	var body strings.Builder
	// 写入HTML结构
	body.WriteString("<html><body>")
	body.WriteString("<p>Please use the following token to complete your registration:</p>")
	body.WriteString(strings.Join([]string{"<p><strong>", token, "</strong></p>"}, " "))
	body.WriteString("<p>This token will expire in <strong>5 minutes</strong>.</p>")
	body.WriteString("</body></html>")
	return body.String()
}

func welcomeBody() string {
	var body strings.Builder
	// 写入HTML结构
	body.WriteString("<html><body>")
	body.WriteString(strings.Join([]string{"<p><h1>Hello ", ": </h1></p>"}, ""))
	body.WriteString(strings.Join([]string{"<p>Hello, welcome to register for the restaurant, ",
		"please pay attention to abide by the rules of the restaurant, ",
		"I wish you a happy use</p>"}, ""))
	body.WriteString("</body></html>")
	return body.String()
}

func sendEmail(email, body string) error {
	// 设置邮件内容
	// 初始化邮件头部
	header := emailHeader(email)
	// 组合邮件头部和正文
	message := []byte(strings.Join([]string{header, body}, ""))
	//通过kafka异步发送欢迎邮件
	// 发送消息到 Kafka
	// 通过 Kafka 异步发送欢迎邮件
	msg := model.EmailMessage{
		Email:   email,
		Message: string(message),
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		zap.L().Error("json.Marshal failed", zap.Error(err))
		return err
	}
	kafka.StartEmailProducer(msgBytes)
	return nil
}
