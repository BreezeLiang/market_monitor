package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/JiangInk/market_monitor/extend/utils"
	"github.com/JiangInk/market_monitor/extend/utils/code"
)

// JWTAuth Token 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Authorization token 值
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			// 获取不到 Authorization 报：请求未携带Token,无权访问
			utils.ResponseFormat(c, code.TokenNotFound, nil)
			return
		}

		// 3. 获取到 Token, 解析token信息

		// 4. 未能正常解析 Token，则报：token认证失败

		// 5. 获取缓存中的Token信息

		// 6. 获取不到，则报：用户注销或token失效

		// 7.

	}
}