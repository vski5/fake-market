package admin

import (
	"fake-market/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(c *gin.Context) {

	cateId, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	//获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	models.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList)
	//获取商品类型属性对应的类型

	goodsType := models.GoodsType{}
	models.DB.Where("id=?", cateId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})

}
func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	//获取当前商品类型属性对应的类型id

	//获取所有的商品类型

}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {

}
func (con GoodsTypeAttributeController) Edit(c *gin.Context) {
	/* id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Find(&goodsType)
		c.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
			"goodsType": goodsType,
		})
	} */
}
func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	/* id, err1 := models.Int(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err2 := models.Int(c.PostForm("status"))
	if err1 != nil || err2 != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType")
		return
	}

	if title == "" {
		con.Error(c, "商品类型的标题不能为空", "/admin/goodsType/edit?id="+models.String(id))
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err3 := models.DB.Save(&goodsType).Error
	if err3 != nil {
		con.Error(c, "修改数据失败", "/admin/goodsType/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/goodsType/index")
	} */
}
func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	/* id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/goodsType/index")
	} else {
		goodsType := models.GoodsType{Id: id}
		models.DB.Delete(&goodsType)
		con.Success(c, "删除数据成功", "/admin/goodsType/index")
	} */
}
