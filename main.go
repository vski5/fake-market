package main

import (
	"fake-market/routers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1" //注意引用的是指定版本号V1 "gopkg.in/ini.v1"
)

func main() {
	r := gin.Default()

	//加载HTML资源
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态资源   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	//自定义模板函数  注意要把这个函数放在加载模板前
	/* r.SetFuncMap(template.FuncMap{
		"函数名": 赋值给函数名的函数（不加括号）,
	}) */

	//读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	//导入路由组
	routers.AdminRouterInit(r)
	routers.ManagerRouterInit(r)
	routers.FocusRouterInit(r)

	ginPort := config.Section("app").Key("port").String()

	r.Run(ginPort)
}
