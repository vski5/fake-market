package admin

import (
	"fake-market/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsTypeController struct {
	BaseController
}

func (con GoodsTypeController) Index(c *gin.Context) {
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})

}
func (con GoodsTypeController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}

func (con GoodsTypeController) DoAdd(c *gin.Context) {
	c.String(200, "test")
}
func (con GoodsTypeController) Edit(c *gin.Context) {
	c.String(200, "test")
}
func (con GoodsTypeController) DoEdit(c *gin.Context) {
	c.String(200, "test")
}
func (con GoodsTypeController) Delete(c *gin.Context) {
	c.String(200, "test")
}
