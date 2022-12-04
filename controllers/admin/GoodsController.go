package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)
	fmt.Printf("%#v", goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
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

func (con GoodsController) Delete(c *gin.Context) {
	c.String(200, "Delete")
}
