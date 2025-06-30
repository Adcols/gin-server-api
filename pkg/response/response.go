package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 预定义业务状态码
const (
	CodeSuccess       = 200 // 成功
	CodeFailed        = 400 // 失败
	CodeUnauthorized  = 401 // 未授权
	CodeForbidden     = 403 // 禁止访问
	CodeNotFound      = 404 // 资源不存在
	CodeServerError   = 500 // 服务器内部错误
	CodeParameterError = 422 // 参数错误
)

// 预定义响应消息
var CodeMessageMap = map[int]string{
	CodeSuccess:       "操作成功",
	CodeFailed:        "操作失败",
	CodeUnauthorized:  "未授权",
	CodeForbidden:     "禁止访问",
	CodeNotFound:      "资源不存在",
	CodeServerError:   "服务器内部错误",
	CodeParameterError: "参数错误",
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: CodeMessageMap[CodeSuccess],
		Data:    data,
	})
}

// SuccessWithMessage 带自定义消息的成功响应
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, code int, data interface{}) {
	message, ok := CodeMessageMap[code]
	if !ok {
		message = CodeMessageMap[CodeFailed]
	}

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// FailWithMessage 带自定义消息的失败响应
func FailWithMessage(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ParameterError 参数错误响应
func ParameterError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeParameterError,
		Message: CodeMessageMap[CodeParameterError],
		Data:    data,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeUnauthorized,
		Message: CodeMessageMap[CodeUnauthorized],
		Data:    data,
	})
}

// ServerError 服务器错误响应
func ServerError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeServerError,
		Message: CodeMessageMap[CodeServerError],
		Data:    err.Error(),
	})
}