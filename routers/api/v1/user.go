package v1

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)
///解析token获取用户身份信息
func GetUserInfo(c *gin.Context){
	token := c.Query("token")
	claims, err := util.ParseToken(token)
	code:=e.SUCCESS
	if err != nil{
		code = e.ERRORAUTHCHECKTOKENFAIL
	}
	data := make(map[string]interface{})
	data["name"]=claims.Username
	data["avatar"]="https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	data["roles"]=claims.Audience
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}