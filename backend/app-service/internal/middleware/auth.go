package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Cookie获取token
		token, err := c.Cookie("auth_token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
			c.Abort()
			return
		}

		// TODO: 验证token
		// 这里简化处理，实际应该使用JWT验证

		c.Set("userId", uint(1))
		c.Next()
	}
}
