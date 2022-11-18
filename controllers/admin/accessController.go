package admin

import "github.com/gin-gonic/gin"

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {

	c.HTML(200, "admin/access/index.html", gin.H{})
}
func (con AccessController) Add(c *gin.Context) {

	c.HTML(200, "admin/access/add.html", gin.H{})
}
func (con AccessController) DoAdd(c *gin.Context) {

	c.String(200, "DoAdd")
}
