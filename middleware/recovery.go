package middleware

import (
	"fmt"
	"log/slog"
	"runtime/debug"

	"github.com/Adcols/gin-server-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录堆栈信息
				stackTrace := string(debug.Stack())

				// 记录错误日志
				slog.Error(fmt.Sprintf("[Recovery] panic recovered: %v\n%s", err, stackTrace))

				// 返回服务器错误响应
				response.ServerError(c, fmt.Errorf("%v", err))
				c.Abort()
			}
		}()

		c.Next()
	}
}
