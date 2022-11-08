package admin

import (
	"fake-market/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCaptchaController struct{ BaseController }

func (con LoginCaptchaController) DoCaptcha(c *gin.Context) {
	//先 生成验证码
	id, b64s, err := models.CaptchaMake()
	if err != nil {
		fmt.Println(err)
	}
	//获取前端传过来的CaptchaId和verifyValue
	captchId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	//排除前端提交的空请求
	if captchId == "" || verifyValue == "" {
		c.JSON(http.StatusOK, "验证码验证失败")
		return
	}
	//调用验证验证码的方法
	if flag := models.CaptchaVerify(captchId, verifyValue); flag {
		//验证通过

		c.JSON(200, gin.H{
			"id":   id,
			"b64s": b64s,
		})
	} else {
		//验证失败
		c.JSON(http.StatusOK, "验证码验证失败")
	}

}
func (con LoginCaptchaController) DoCaptchaMake(c *gin.Context) {
	id, b64s, err := models.CaptchaMake()
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"captchaId":  id,
			"captchaImg": b64s,
		})
	}

}
