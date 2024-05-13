package constant

import "errors"

var (
	ErrInvalidParam   = errors.New("请求的参数有误")
	ErrEmailFormat    = errors.New("邮箱格式有误")
	ErrNotFound       = errors.New("未找到数据")
	ErrServerBusy     = errors.New("服务繁忙")
	ErrNotExists      = errors.New("用户不存在")
	ErrPassword       = errors.New("密码错误")
	ErrAuthCode       = errors.New("验证码错误")
	ErrExistsUser     = errors.New("用户已存在")
	ErrUserNotLogin   = errors.New("用户未登录")
	ErrStockNotEnough = errors.New("库存不足")
	ErrExistsData     = errors.New("数据已存在")
	ErrBadRedisData   = errors.New("该数据为防止缓存穿透的坏数据")
	ErrFormat         = errors.New("token格式错误")
)
