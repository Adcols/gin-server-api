package response

import (
	"github.com/Adcols/gin-server-api/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    errors.Success,
		Message: errors.GetMessage(errors.Success),
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, err error) {
	// 处理自定义错误
	if e, ok := err.(*errors.Error); ok {
		c.JSON(http.StatusOK, Response{
			Code:    e.Code,
			Message: e.Message,
			Data:    nil,
		})
		return
	}

	// 处理标准错误
	c.JSON(http.StatusOK, Response{
		Code:    errors.InternalServerError,
		Message: err.Error(),
		Data:    nil,
	})
}

// ParameterError 参数错误响应
func ParameterError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    errors.BadRequest,
		Message: message,
		Data:    nil,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    errors.Unauthorized,
		Message: errors.GetMessage(errors.Unauthorized),
		Data:    nil,
	})
}
