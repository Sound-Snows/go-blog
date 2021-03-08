package jwt

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//JWT 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("authorization")
		if token == "" {
			code = e.INVALIDPARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERRORAUTHCHECKTOKENFAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERRORAUTHCHECKTOKENTIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
