package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a AdminController) Add(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{})
}
func (a AdminController) Back(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}
