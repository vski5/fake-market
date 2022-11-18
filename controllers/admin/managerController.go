package admin

import (
	"fake-market/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ManagerController struct{ BaseController }

// 后台管理系统的主页
func (a ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	c.HTML(200, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

// 增加manager的页面
func (a ManagerController) Add(c *gin.Context) {
	//获取所有manager
	roleList := []models.Role{} //注意，这里是role，不是manager，manager的状态/权限由role表决定
	models.DB.Find(&roleList)

	c.HTML(200, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

// 执行添加manager
func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err1 := models.Int(c.PostForm("role_id"))
	if err1 != nil {
		con.Error(c, "传入数据错误", "/admin/manager/add")
		return
	}
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile, _ := strconv.Atoi(strings.Trim(c.PostForm("mobile"), " "))
	//用户名和密码长度是否合法
	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名或者密码的长度不合法", "/admin/manager/add")
		return
	}

	//判断管理是否存在
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "此管理员已存在", "/admin/manager/add")
		return
	}
	//执行增加管理员
	manager := models.Manager{
		Username: username,
		Password: models.MD5maker(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  models.GetUnix(),
	}
	err2 := models.DB.Create(&manager).Error
	if err2 != nil {
		con.Error(c, "增加管理员失败", "/admin/manager/add")
		return
	}

	con.Success(c, "增加管理员成功", "/admin/manager/index")

}

// 修改 manager 的页面
func (con ManagerController) Edit(c *gin.Context) {

	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager/index")
	} else {
		// 判断管理是否存在,回传数值
		managerList := []models.Manager{}
		//查询对应数值
		models.DB.Where("id=?", id).Find(&managerList)
		c.HTML(200, "admin/manager/edit.html", gin.H{
			"managerList": managerList,
		})
	}

}

// 执行 修改manager 的页面
func (con ManagerController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/manager/index")
	} else {
		// 判断管理是否存在,回传数值
		/* 		managerList := []models.Manager{}
		   		//查询对应数值
		   		models.DB.Where("id=?", id).Find(&managerList)
		   		c.HTML(200, "admin/manager/edit.html", gin.H{
		   			"managerList": managerList,
		   		}) */

		username := strings.Trim(c.PostForm("username"), " ")
		password := strings.Trim(c.PostForm("password"), " ")
		email := strings.Trim(c.PostForm("email"), " ")
		mobile, _ := strconv.Atoi(strings.Trim(c.PostForm("mobile"), " "))

		// 用户名和密码长度是否合法
		if len(username) < 2 || len(password) < 6 {
			con.Error(c, "用户名或者密码的长度不合法", "/admin/manager/index")
			return
		}

		// 判断管理是否存在
		managerList := []models.Manager{}
		models.DB.Where("id=?", id).Find(&managerList)
		if len(managerList) > 0 {
			// 执行增加管理员
			manager := models.Manager{
				Username: username,
				Password: models.MD5maker(password),
				Email:    email,
				Mobile:   mobile,
				Status:   1,
				AddTime:  models.GetUnix(),
			}

			err2 := models.DB.Model(&manager).Select("Username", "Password", "Email", "Mobile", "Status", "AddTime").Updates(manager)
			if err2 != nil {
				con.Error(c, "更新管理员失败", "/admin/manager/index")
				return
			}

			con.Success(c, "更新管理员成功", "/admin/manager/index")
		}

	}
}

// 删除商品
func (a ManagerController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}
