package middlewares

import (
	"encoding/json"
	"fake-market/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context /* , username string */) {
	//判断用户【管理员】是否登录
	pathname := strings.Split(c.Request.URL.String(), "?")[0] //获取 Url 路径去掉 Get 传值
	//获取值【我需要一种在两个函数间传递值的方法】
	/* admincUsername := c.PostForm("username") */
	userinfo := models.CookieRedisStore{}.Get("admin")
	//先类型断言判断是否为string，确定之后才能进行下一步
	//类型断言
	userinfoStr, ok := userinfo.(string)
	if ok {
		if len(userinfoStr) > 0 {
			var u []models.Manager

			json.Unmarshal([]byte(userinfoStr), &u)

			if len(u) > 0 && u[0].Username != "" {

				/* 			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
					   c.Redirect(302, "/admin/login")
				   }
				*/
				fmt.Println("成功跳过验证登录的中间件")
			}
		} else {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}

	} else {
		c.HTML(200, "admin/public/error.html", gin.H{
			"gotourl": "用户名或者密码异常",
			"message": "/admin/login",
		})
	}

}
