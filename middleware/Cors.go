package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//跨域中间件
func Cors() gin.HandlerFunc {
	return func (c *gin.Context){
		method :=c.Request.Method
		//domainName :=setting.DomainName
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method=="OPTIONS"{
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
