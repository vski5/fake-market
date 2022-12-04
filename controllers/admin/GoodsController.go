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
	//获取顶级分类
	goodsCateList := []models.GoodsCate{}
	models.DB.Where("pid = 0").Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}
func (con GoodsController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := models.Int(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "传入参数类型不正确", "/goods/add")
		return
	}
	if err2 != nil {
		con.Error(c, "排序值必须是整数", "/goods/add")
		return
	}
	cateImgDir, _ := models.UploadOneImg(c, "cate_img", "./static/goodsUpload/")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		con.Error(c, "增加数据失败", "/admin/goods/add")
		return
	}
	con.Success(c, "增加数据成功", "/admin/goods/index")
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
