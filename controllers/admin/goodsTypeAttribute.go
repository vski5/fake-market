package admin

import (
	"fake-market/models"
	"net/http"
	"strings"

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

	cateId, err := models.Int(c.Query("cate_id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType/index")
		return
	}

	//获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"cateId":        cateId,
	})

}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err1 := models.Int(c.PostForm("cate_id"))
	attrType, err2 := models.Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := models.Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(c, "非法请求", "/admin/goodsType/index")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?cate_id="+models.String(cateId))
		return
	}

	if err3 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/add?cate_id="+models.String(cateId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?cate_id="+models.String(cateId))
	} else {
		con.Success(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute/index?id="+models.String(cateId))
	}
}
func (con GoodsTypeAttributeController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	cateId, err2 := models.Int(c.Query("cate_id"))
	if err != nil && err2 != nil {
		con.Error(c, "传入数据错误", "/admin/goodsTypeAttribute/index")
	} else {
		goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
		models.DB.Find(&goodsTypeAttribute)
		c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
			"goodsType": goodsTypeAttribute,
			"cateId":    cateId,
		})
	}
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
	c.String(200, "doedit")
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
