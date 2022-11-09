package admin

import "github.com/gin-gonic/gin"

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
