package admin

import (
	"fake-market/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsInfoController struct {
	BaseController
}

// 实现图片上传
func (con GoodsInfoController) ImageUpload(c *gin.Context) {
	savePath, err := models.UploadOneImg(c, "file", "./static/goodsUpload/")
	if err != nil {
		c.JSON(200, gin.H{"link": ""})
	} else {
		// 返回 json 数据 {link: 'path/to/image.jpg'}
		c.JSON(200, gin.H{"link": "/" + savePath})
	}
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
