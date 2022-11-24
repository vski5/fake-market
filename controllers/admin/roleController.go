package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

// 改用户权限的 主页路由
func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	fmt.Println(roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})

}

// 改用户权限的 增加用户路由
func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

// 改用户权限的 增加这一步操作
func (con RoleController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/add")
		return
	}
	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(time.Now().Unix())

	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "增加角色失败 请重试", "/admin/role/add")
	} else {
		con.Success(c, "增加角色成功", "/admin/role")
	}

}
func (con RoleController) Edit(c *gin.Context) {

	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Find(&role)
		c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
			"role": role,
		})
	}

}
func (con RoleController) DoEdit(c *gin.Context) {

	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/edit")
	}

	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description

	err2 := models.DB.Save(&role).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role/edit?id="+models.String(id))
	}

	//查询要修改的数据 然后 修改

	// c.String(http.StatusOK, "-执行修改")
}
func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(c, "删除数据成功", "/admin/role")
	}
}

// 权限管理界面
func (con AccessController) Auth(c *gin.Context) {}

// 执行权限的管理
func (con AccessController) DoAuth(c *gin.Context) {}
