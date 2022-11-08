package admin

import "github.com/gin-gonic/gin"

type BaseController struct {
}

// 返回公共的成功页面
func (con BaseController) Success(c *gin.Context) {
	c.String(200, "format string")
}

// 返回公共的失败页面
func (con BaseController) Error(c *gin.Context) {

}
