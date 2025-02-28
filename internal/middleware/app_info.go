package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// gin.Context提供的setter和getter方法，可以在请求上下文中设置和获取值，在gin中被稱為元數據管理(metadata Management)。
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
