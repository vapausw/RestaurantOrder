package logic

import (
	co "RestaurantOrder/constant"
	"RestaurantOrder/dao/mysql"
	"RestaurantOrder/dao/redis"
	"RestaurantOrder/model"
	"RestaurantOrder/pkg/util"
	"errors"
	"go.uber.org/zap"
)

func MerchantLogin(m *model.Merchant) (error, co.MyCode) {
	// 验证邮箱格式是否有误
	if !util.ValidateEmail(m.Email) {
		return co.ErrEmailFormat, co.CodeInvalidParam
	}
	// 保存原始密码
	OriginalPassword := m.Password
	// 通过邮箱获取商家加密的密码以及商家id
	// 从数据库直接获取
	if err := mysql.GetMerchantPassword(m); err != nil {
		if errors.Is(err, co.ErrNotFound) {
			return co.ErrNotExists, co.CodeNotFound
		}
		return err, co.CodeServerBusy
	}
	if !util.CheckPasswordHash(OriginalPassword, m.Password) {
		return co.ErrPassword, co.CodeErrPassword
	}
	return nil, co.CodeSuccess
}

func SendMerchantCode(email string) (error, co.MyCode) {
	// 验证邮箱格式是否有误
	if !util.ValidateEmail(email) {
		return co.ErrEmailFormat, co.CodeInvalidParam
	}
	// 生成验证码
	code := util.GenValidateCode()
	// 将验证码存入redis
	if err := redis.SetMerchantCode(email, code); err != nil {
		zap.L().Error("redis.SetMerchantCode failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	// 发送验证码，使用kafka异步发送
	zap.L().Info("code", zap.String("code", code))
	return nil, co.CodeSuccess
}

func MerchantRegister(m *model.Merchant) (error, co.MyCode) {
	// 验证邮箱格式是否有误
	if !util.ValidateEmail(m.Email) {
		return co.ErrEmailFormat, co.CodeInvalidParam
	}
	// 验证码密码是否一致
	if m.Password != m.RePassword {
		return co.ErrPassword, co.CodeErrPassword
	}
	// 验证码是否正确
	code, err := redis.GetMerchantCode(m.Email)
	if err != nil {
		zap.L().Error("redis.GetUserCode failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	if code != m.Code {
		return co.ErrAuthCode, co.CodeErrCode
	}
	// 将用户存储到mysql数据库中
	m.Password, err = util.HashPassword(m.Password)
	if err != nil {
		zap.L().Error("redis.GetUserCode failed", zap.Error(err))
		return co.ErrServerBusy, co.CodeServerBusy
	}
	if err := mysql.CreateMerchant(m); err != nil {
		if errors.Is(err, co.ErrExistsUser) {
			return co.ErrExistsUser, co.CodeExistsCode
		}
		return co.ErrServerBusy, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}

func MerchantInfo(m *model.Shop) (error, co.MyCode) {
	// 直接插入到mysql
	if err := mysql.CreateShop(m); err != nil {
		return co.ErrServerBusy, co.CodeServerBusy
	}
	return nil, co.CodeSuccess
}
