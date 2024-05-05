package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/cache"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"errors"
	"go.uber.org/zap"
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
		zap.L().Error("redis.GetUserCode failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	if code != u.Code {
		return co.ErrAuthCode, co.CodeErrCode
	}
	// 将用户存储到mysql数据库中
	u.Password, err = util.HashPassword(u.Password)
	if err != nil {
		zap.L().Error("redis.GetUserCode failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	u.UserName = util.GenerateRandomNickname()
	if err := mysql.CreateUser(u); err != nil {
		if errors.Is(err, co.ErrExistsUser) {
			return co.ErrExistsUser, co.CodeExistsCode
		}
		return co.ErrServerBusy, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}
