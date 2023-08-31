package defaults

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "front/index/index.html", gin.H{})

}
