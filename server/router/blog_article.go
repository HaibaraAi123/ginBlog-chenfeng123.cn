package router

import (
	v1 "HaibaraAi123/ginBlog-chenfeng123.cn/server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	ArticleRouterGroup := Router.Group("article")
	{
		ArticleRouterGroup.GET("", v1.GetArticleList)
		ArticleRouterGroup.POST("", v1.PostArticle)
		ArticleRouterGroup.GET("/:id", v1.GetArticleDetail)
	}
}
