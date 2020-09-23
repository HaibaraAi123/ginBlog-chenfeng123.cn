package initialize

import (
	_ "HaibaraAi123/ginBlog-chenfeng123.cn/server/docs"
	"HaibaraAi123/ginBlog-chenfeng123.cn/server/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	var r = gin.New()
	//cannot use ginSwagger.WrapHandler(swaggerFiles.Handler) (type "github.com/gin-gonic/gin".HandlerFunc) as type "HaibaraAi123/chenfeng123-gin/vendor/github.com/gin-gonic/gin".HandlerFunc in argument to r.RouterGroup.GET
	//after delete vendor/gin
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	APIGroup := r.Group("")
	router.InitUserRouter(APIGroup)
	return r
}
