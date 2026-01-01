package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/mini-blog/internal/pkg/response"
	"github.com/nanfeng/mini-blog/internal/pkg/util"
)

// AuthMiddleware 通用token校验中间间
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从请求头中获取 auth 信息
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			response.BadRequest(c, "请求头中缺少Authorization")
			c.Abort()
			return
		}

		// 2.校验请求头中携带的token格式
		parts := strings.Split(auth, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			response.BadRequest(c, "请求头格式错误")
			c.Abort()
			return
		}

		// 3.解析token， 判断token是否有效
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			response.BadRequest(c, err.Error())
		}

		// 4.将用户id放入 context 中
		c.Set("user_id", claims.Subject)
		c.Next()
	}
}

// AdminAuthMmiddleware 管理员token校验
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.从请求头中获取 auth 信息
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			response.BadRequest(c, "请求头中缺少Authorization")
			c.Abort()
			return
		}

		// 2.校验请求头中携带的token格式
		parts := strings.Split(auth, " ")
		if parts[0] != "Bearer" || len(parts) != 2 {
			response.BadRequest(c, "请求头格式错误")
			c.Abort()
			return
		}

		// 3.解析token， 判断token是否有效
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			response.BadRequest(c, err.Error())
		}

		// 4.将用户id放入 context 中
		c.Set("user_id", claims.Subject)
		c.Next()
	}
}
