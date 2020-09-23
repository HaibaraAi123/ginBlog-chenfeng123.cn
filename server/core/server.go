package core

import (
	"HaibaraAi123/ginBlog-chenfeng123.cn/server/initialize"
)

func RunWindowServer() {
	Router := initialize.Routers()
	Router.Run(":8080")
}
