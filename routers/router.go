package routers

import (
	"gin-blog/middleware"
	"gin-blog/pkg/setting"
	v1 "gin-blog/routers/api/v1"

	"gin-blog/routers/api"

	"gin-blog/middleware/jwt"

	"github.com/gin-gonic/gin"
)

//InitRouter 获取路由
func InitRouter() *gin.Engine {
	r := gin.New()
    r.Use(middleware.Cors())
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//
	r.POST("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/user/info",v1.GetUserInfo)
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTag)

		apiv1.PUT("/tags", v1.EditTag)

		apiv1.DELETE("/tags", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("articles/:id", v1.DeleteArticle)
	}

	return r
}
