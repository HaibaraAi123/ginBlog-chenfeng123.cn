package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetExaUser(c *gin.Context) {
	c.String(http.StatusOK, "hello")
}

func PostExaUser(c *gin.Context) {

}
