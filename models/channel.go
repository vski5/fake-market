package models

// 引用之前先类型断言
func CookieChannel(passwordMD5 interface{}) interface{} {
	passwordMD5str, ok := passwordMD5.(string)
	if ok {
		Passwordch := make(chan string, 1)

		Passwordch <- passwordMD5str
		passwordMD6 := <-Passwordch
		return passwordMD6
	} else {
		return nil
	}

}
