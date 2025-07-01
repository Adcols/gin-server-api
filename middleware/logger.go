package middleware

import (
	"fmt"
	"time"

	"github.com/Adcols/gin-server-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latency := endTime.Sub(startTime)

		// 请求方法
		method := c.Request.Method

		// 请求路由
		uri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 客户端IP
		clientIP := c.ClientIP()

		// 日志格式
		logString := fmt.Sprintf("%s | %3d | %13v | %15s | %s",
			method,
			statusCode,
			latency,
			clientIP,
			uri,
		)

		// 根据状态码记录日志
		if statusCode >= 500 {
			logger.Error("%s", logString)
		} else if statusCode >= 400 {
			logger.Warn("%s", logString)
		} else {
			logger.Info("%s", logString)
		}
	}
}
