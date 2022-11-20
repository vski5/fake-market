package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessController struct{}

func (con AccessController) Index(c *gin.Context) {

	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})

}
func (con AccessController) Add(c *gin.Context) {
	//获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}
func (con AccessController) DoAdd(c *gin.Context) {
	//先获取form
	module_name := strings.Trim(c.PostForm("module_name"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	action_name := strings.Trim(c.PostForm("action_name"), " ")
	url_name := strings.Trim(c.PostForm("url"), " ")

	type_name := strings.Trim(c.PostForm("type"), " ")
	typeInt, err1 := models.Int(type_name)
	module_id := strings.Trim(c.PostForm("module_id"), " ")
	module_id_Int, err2 := models.Int(module_id)
	sort := strings.Trim(c.PostForm("sort"), " ")
	sortInt, err3 := models.Int(sort)
	status := strings.Trim(c.PostForm("status"), " ")
	statusInt, err4 := models.Int(status)
	//判断form内容是否合规
	//1.module_name不能为空值
	if module_name == "" || module_name == " " {
		BaseController{}.Error(c, "model name 不能为空值", "/admin/access/index")
		return
	}
	//2.判断传入的int是否为纯数字
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		BaseController{}.Error(c, "某些值应该是数字", "/admin/access/index")

		return
	}
	//实例化Access
	accessLise := &models.Access{
		ModuleName:  module_name,   //模块名称
		ActionName:  action_name,   //操作名称
		Type:        typeInt,       //节点类型 :  1、表示模块    2、表示菜单     3、操作
		Url:         url_name,      //路由跳转地址
		ModuleId:    module_id_Int, //此module_id和当前模型的id关联       module_id= 0 表示模块
		Sort:        sortInt,
		Description: description,
		Status:      statusInt,
		AddTime:     models.GetUnix(),
	}
	//使用gorm添加实例化的Access到MySQL里面
	err5 := models.DB.Create(&accessLise).Error
	if err5 != nil {
		BaseController{}.Error(c, "增加数据失败", "/admin/access/add")
		return
	}
	BaseController{}.Success(c, "增加数据成功", "/admin/access/index")

}
func (con AccessController) Edit(c *gin.Context) {
	// 获取id
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		BaseController{}.Error(c, "获取id失败", "/admin/access/index")
		return
	}
	// 用id实例化access结构体，方便查找
	access := &models.Access{
		Id: id,
	}
	models.DB.Find(access)
	//获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id=?", id).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"access":     access,
		"accessList": accessList,
	})

}
func (con AccessController) DoEdit(c *gin.Context) {

	c.String(200, "DoEdit")
}
