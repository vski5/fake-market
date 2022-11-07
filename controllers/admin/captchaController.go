package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCaptchaController struct{}

func (con LoginCaptchaController) DoCaptcha(c *gin.Context) {
	//获取前端传过来的CaptchaId和verifyValue
	captchId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	//排除前端提交的空请求
	if captchId == "" || verifyValue == "" {
		c.JSON(http.StatusOK, "验证码验证失败")
		return
	}
	//调用验证验证码的方法
	if flag := models.CaptchaVerify(captchId, verifyValue); flag == true {
		//验证通过
		fmt.Println(flag)
		c.JSON(http.StatusOK, "验证码验证成功")
	} else {
		//验证失败
		c.JSON(http.StatusOK, "验证码验证失败")
	}

}
