package admin

import (
	"fake-market/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

// 返回公共的成功页面
func (con BaseController) Success(c *gin.Context, message string, gotourl string) {
	c.HTML(200, "admin/public/success.html", gin.H{
		"message": message,
		"gotourl": gotourl,
	})
}

// 返回公共的失败页面
func (con BaseController) Error(c *gin.Context, message string, gotourl string) {
	c.HTML(200, "admin/public/error.html", gin.H{
		"message": message,
		"gotourl": gotourl,
	})
}

// 公共修改状态的方法
func (con BaseController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table") //要修改的表名
	field := c.Query("field") //要修改的字段名

	//ABS()取绝对值

	/*
		js：返回"data-table"和"data-field"，用table和field的名字调用路由/admin/changeStatus
		var table=$(this).attr("data-table")
		var field=$(this).attr("data-field")
		$.get("/admin/changeStatus",{id:id,table:table,field:field},function(response){

		HTML：
		data-table="focus" data-field="status"

	*/

	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败 请重试",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改成功",
		})
	}
}
