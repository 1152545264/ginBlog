package routers

import (
	"ginBlog/middleware/jwt"
	"ginBlog/pkg/setting"
	"ginBlog/pkg/upload"
	"ginBlog/routers/api"
	v1 "ginBlog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files" // v1.6之后这个包名换了
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// InitRouter 初始化默认的多路复用器
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tags", v1.AddTags)
		//更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)

		//导出标签
		r.POST("/tags/export", v1.ExportTag)
	}
	return r
}
