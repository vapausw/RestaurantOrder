package constant

const (
	ContextUserIDKey = "user_id"
)

type MyCode int64
type ResponseData struct {
	Code    MyCode      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	CodeSuccess           MyCode = 1000 // 成功
	CodeInvalidAuthFormat MyCode = 1001 // 认证格式有误
	CodeInvalidToken      MyCode = 1002 // 无效的Token
	CodeInvalidParam      MyCode = 1003 // 请求参数错误
	CodeServerBusy        MyCode = 1004 // 服务繁忙
	CodeNotFound          MyCode = 1005 // 未找到
	CodeErrPassword       MyCode = 1006 // 密码错误
	CodeErrCode           MyCode = 1007 // 验证码错误
	CodeExistsCode        MyCode = 1008 // 用户已存在
	CodeUserNotLogin      MyCode = 1009 // 用户未登录
	CodeStockNotEnough    MyCode = 1010 // 库存不足
)
