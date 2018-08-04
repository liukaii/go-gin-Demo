package mdw

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

//全局中间件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

//鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		//http.StatusUnauthorized = 401
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}