package middleware

import (
	"strings"

	"github.com/Adcols/gin-server-api/pkg/errors"
	"github.com/Adcols/gin-server-api/pkg/response"
	"github.com/Adcols/gin-server-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			response.Error(c, errors.New(errors.Unauthorized, nil))
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, errors.New(errors.Unauthorized, nil))
			c.Abort()
			return
		}

		// 解析JWT令牌
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Error(c, errors.NewWithMessage(errors.Unauthorized, "无效的令牌", err))
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
