package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Focus，轮播图管理
type FocusController struct{}

// 后台轮播图管理管理系统的主页
func (a FocusController) Index(c *gin.Context) {
	c.HTML(200, "admin/focus/index.html", gin.H{})
}

// 增加轮播图
func (a FocusController) Add(c *gin.Context) {
	c.HTML(200, "admin/focus/add.html", gin.H{})
}

// 修改轮播图
func (a FocusController) Edit(c *gin.Context) {
	c.HTML(200, "admin/focus/edit.html", gin.H{})
}

// 删除轮播图
func (a FocusController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "test2")
}
