package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsInfoController struct {
	BaseController
}

func (con GoodsInfoController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
}
func (con GoodsInfoController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{})
}
func (con GoodsInfoController) DoAdd(c *gin.Context) {

}
func (con GoodsInfoController) Edit(c *gin.Context) {

}
func (con GoodsInfoController) DoEdit(c *gin.Context) {

}
func (con GoodsInfoController) Delete(c *gin.Context) {

}
