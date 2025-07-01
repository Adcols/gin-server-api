package errors

// 错误码定义
const (
	// 成功
	Success = 200

	// 客户端错误 4xx
	BadRequest       = 400 // 请求参数错误
	Unauthorized     = 401 // 未授权
	Forbidden        = 403 // 禁止访问
	NotFound         = 404 // 资源不存在
	MethodNotAllowed = 405 // 方法不允许
	Conflict         = 409 // 资源冲突
	TooManyRequests  = 429 // 请求过多

	// 服务端错误 5xx
	InternalServerError = 500 // 服务器内部错误
	NotImplemented      = 501 // 未实现
	ServiceUnavailable  = 503 // 服务不可用

	// 业务错误码 1xxx
	UserNotFound        = 1001 // 用户不存在
	UserAlreadyExists   = 1002 // 用户已存在
	PasswordIncorrect   = 1003 // 密码错误
	UserDisabled        = 1004 // 用户已禁用
	EmailAlreadyExists  = 1005 // 邮箱已存在
	PhoneAlreadyExists  = 1006 // 手机号已存在
	UsernameExists      = 1007 // 用户名已存在
	PasswordEncryptFail = 1008 // 密码加密失败
)

// 错误码消息映射
var codeMessages = map[int]string{
	Success:             "成功",
	BadRequest:          "请求参数错误",
	Unauthorized:        "未授权",
	Forbidden:           "禁止访问",
	NotFound:            "资源不存在",
	MethodNotAllowed:    "方法不允许",
	Conflict:            "资源冲突",
	TooManyRequests:     "请求过多",
	InternalServerError: "服务器内部错误",
	NotImplemented:      "未实现",
	ServiceUnavailable:  "服务不可用",
	UserNotFound:        "用户不存在",
	UserAlreadyExists:   "用户已存在",
	PasswordIncorrect:   "密码错误",
	UserDisabled:        "用户已禁用",
	EmailAlreadyExists:  "邮箱已存在",
	PhoneAlreadyExists:  "手机号已存在",
	UsernameExists:      "用户名已存在",
	PasswordEncryptFail: "密码加密失败",
}

// GetMessage 获取错误码对应的消息
func GetMessage(code int) string {
	if msg, ok := codeMessages[code]; ok {
		return msg
	}
	return "未知错误"
}
