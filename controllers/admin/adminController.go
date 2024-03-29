package admin

import (
	"fake-market/models"
	"gopkg.in/ini.v1" //注意引用的是指定版本号V1 "gopkg.in/ini.v1"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
)

type AdminController struct{ BaseController }

func (a AdminController) Login(c *gin.Context) {

	c.HTML(200, "admin/login/login.html", gin.H{})
}

// 此处是验证
func (con AdminController) Dolog(c *gin.Context) {
	/* 先写验证码的验证 */
	//先 生成验证码
	/* 	id, b64s, err := models.CaptchaMake()
	   	if err != nil {
	   		fmt.Println(err, id, b64s)
	   	} */
	//获取前端传过来的CaptchaId和verifyValue
	captchId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	/* 	//排除前端提交的空请求
	   	if captchId == "" || verifyValue == "" {
	   		c.JSON(http.StatusOK, "验证码验证失败")
	   		return
	   	} */
	//调用验证验证码的方法
	if flag := models.CaptchaVerify(captchId, verifyValue); flag {
		//获取前端传过来的username和password
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 获取用户名密码 对密码进行加密
		//后期还要写 加盐
		passwordMD5 := models.MD5maker(password)

		managerinfo := []models.Manager{}

		/*此处有拼接SQL语句漏洞，试了一下可能没问题，但是感觉不对，只要是拼接就很有可能有问题，这里传的是MD5之后的值，所以不怕拼接*/
		models.DB.Where("username=? AND password =?", username, passwordMD5).First(&managerinfo)

		if len(managerinfo) > 0 {

			//设置一个cookie方便验证
			//设置cookie，第5个参数表示所有20011111.xyz的二级域名（比如a.20011111.xyz b.20011111.xyz）都共享一个cookie
			//第4个表示路径，/admin 路径的共享这个cookie，如果设置为/，就是所有都共享。
			//第一个是cookie的名字
			config1, err := ini.Load("./conf/app.ini")
			if err != nil {
				fmt.Printf("Fail to read file: %v", err)
				os.Exit(1)
			}
			cookieWeb := config1.Section("webcookie").Key("cookieinfoDir").String() //cookieWeb是.20011111.xyz
			c.SetCookie("admin_cookie", username, 3600, "/admin", cookieWeb, false, false)

			//设置一个session，保持登录状态
			/* 			managerinfoSlice, _ := json.Marshal(managerinfo)
			   			string(managerinfoSlice) */
			errCookie := models.CookieRedisStore{}.Set(username, passwordMD5)
			if errCookie != nil {
				return
			}

			//使用BaseController中的返回公共的成功页面
			/* con := &BaseController{} */
			con.Success(c, "登录成功", "/admin/manager/index")
		} else {
			con.Error(c, "用户名或者密码错误", "/admin/login")
		}

	} else {
		//验证失败
		/* con := &BaseController{} */
		con.Error(c, "验证失败,返回登录界面", "login")

	}

}
