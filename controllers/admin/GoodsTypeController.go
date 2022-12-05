package admin

import (
	"fake-market/models"
	"net/http"
	"strings"

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
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err1 := models.Int(c.PostForm("status"))

	if err1 != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType/add")
		return
	}

	if title == "" {
		con.Error(c, "标题不能为空", "/admin/goodsType/add")
		return
	}
	goodsType := models.GoodsType{
		Title:       title,
		Description: description,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}

	err := models.DB.Create(&goodsType).Error
	if err != nil {
		con.Error(c, "增加商品类型失败 请重试", "/admin/goodsType/add")
	} else {
		con.Success(c, "增加商品类型成功", "/admin/goodsType/index")
	}
}
func (con GoodsTypeController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Find(&goodsType)
		c.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
			"goodsType": goodsType,
		})
	}
}
func (con GoodsTypeController) DoEdit(c *gin.Context) {
	c.String(200, "test")
}
func (con GoodsTypeController) Delete(c *gin.Context) {
	c.String(200, "test")
}
