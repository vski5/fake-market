package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct{}

// 后台管理系统的主页
func (a ManagerController) Index(c *gin.Context) {
	c.HTML(200, "admin/manager/index.html", gin.H{})
}

// 增加商品
func (a ManagerController) Add(c *gin.Context) {
	c.HTML(200, "admin/manager/add.html", gin.H{})
}

// 修改商品
func (a ManagerController) Edit(c *gin.Context) {
	c.HTML(200, "admin/manager/edit.html", gin.H{})
}

// 删除商品
func (a ManagerController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}
