package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Focus，轮播图管理
type FocusController struct {
	BaseController
}

// 后台轮播图管理管理系统的主页
func (a FocusController) Index(c *gin.Context) {
	focusList := []models.Focus{}
	models.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

// 增加轮播图
func (a FocusController) Add(c *gin.Context) {
	//获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

// 执行--增加轮播图
func (con FocusController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err2 := models.Int(c.PostForm("sort"))
	status, err3 := models.Int(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.Error(c, "非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/add")
	}
	//上传文件
	focusImgSrc, err4 := models.UploadOneImg(c, "focus_img", "./static/focus/")
	if err4 != nil {
		fmt.Println(err4)
	}

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImgSrc,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}
	err5 := models.DB.Create(&focus).Error
	if err5 != nil {
		con.Error(c, "增加轮播图失败", "/admin/focus/add")
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}

}

// 修改轮播图
func (a FocusController) Edit(c *gin.Context) {
	c.HTML(200, "admin/focus/edit.html", gin.H{})
}

// 执行--修改轮播图
func (a FocusController) DoEdit(c *gin.Context) {
	c.HTML(200, "admin/focus/edit.html", gin.H{})
}

// 删除 轮播图
func (a FocusController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}
