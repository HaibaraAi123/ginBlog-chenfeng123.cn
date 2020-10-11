package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Article
// @Summary	获取文章列表
// @Produce  json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /article [get]
func GetArticleList(c *gin.Context) {
	c.String(http.StatusOK, "get article list")
}

// @Tags Article
// @Summary 上传文章
// @Produce  json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /article [post]
func PostArticle(c *gin.Context) {
	c.String(http.StatusOK, "post article")
}

// @Tags Article
// @Summary 获取文章详细内容
// @Produce  json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /article/{id} [get]
func GetArticleDetail(c *gin.Context) {
	c.String(http.StatusOK, "get article detail by article id")
}
