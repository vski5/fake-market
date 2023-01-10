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
	//获取商品分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	//获取所有颜色信息
	goodsColorList := []models.GoodsColor{}
	models.DB.Find(&goodsColorList)

	//获取商品规格包装
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}
func (con GoodsInfoController) DoAdd(c *gin.Context) {

}
func (con GoodsInfoController) Edit(c *gin.Context) {

}
func (con GoodsInfoController) DoEdit(c *gin.Context) {

}
func (con GoodsInfoController) Delete(c *gin.Context) {

}
