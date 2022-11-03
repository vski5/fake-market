package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a AdminController) Login(c *gin.Context) {
	c.HTML(200, "admin/login/login.html", gin.H{})
}
func (a AdminController) Back(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}
