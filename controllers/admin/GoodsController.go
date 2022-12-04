package admin

import "github.com/gin-gonic/gin"

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.String(200, "index")
}
func (con GoodsController) Add(c *gin.Context) {
	c.String(200, "add")
}
func (con GoodsController) DoAdd(c *gin.Context) {
	c.String(200, "doadd")
}
func (con GoodsController) Edit(c *gin.Context) {
	c.String(200, "edit")
}

func (con GoodsController) DoEdit(c *gin.Context) {
	c.String(200, "doedit")
}
