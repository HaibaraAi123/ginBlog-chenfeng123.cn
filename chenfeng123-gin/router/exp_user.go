package router

import (
	v1 "HaibaraAi123/ginBlog-chenfeng123.cn/server/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouterGroup := Router.Group("user")
	//user middlleware here
	//get or post methods
	{
		UserRouterGroup.GET("", v1.GetExaUser)
		UserRouterGroup.POST("", v1.PostExaUser)
	}
}
