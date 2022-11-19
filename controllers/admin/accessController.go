package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

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

	c.String(200, "DoAdd")
}
