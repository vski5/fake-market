package models

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 设置自带的 store
// var store = base64Captcha.DefaultMemStore
// 配置RedisStore  RedisStore实现base64Captcha.Store接口
var store base64Captcha.Store = RedisStore{}

func CaptchaMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	// 配置验证码信息
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          2,      /*改个简单的验证码，方便测试*/
		Source:          "1234", /* "1234567890qwertyuioplkjhgfdsazxcvbnm" */
		BgColor: &color.RGBA{
			R: 3, G: 102, B: 214, A: 125},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	// ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	// 创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	// Generate 生成随机 id、base64 图像字符串
	id, b64s, err = captcha.Generate()
	return id, b64s, err
}

// 验证 captcha 是否正确
func CaptchaVerify(id string, capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}
}
