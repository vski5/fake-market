package main

import (
	"fake-market/models"
	"fake-market/routers"
	"fmt"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1" //注意引用的是指定版本号V1 "gopkg.in/ini.v1"
)

func main() {
	// 读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//加载日志库
	zapInfoDir := config.Section("zap").Key("infoDir").String()
	models.InitLogger(zapInfoDir)

	r := gin.Default()
	//自定义模板函数   注意顺序，注册模板函数需要在加载模板上面
	/* r.SetFuncMap(template.FuncMap{
		"函数名": 赋值给函数名的函数（不加括号）,
	}) */

	r.SetFuncMap(template.FuncMap{
		"UnixToToTime": models.UnixToTime,
	})
	//加载HTML资源
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态资源   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	/*写入中间件*/
	/*  为后台验证而生的鉴权按，写在路由比较好，不适合全局配置中间件
	r.Use(middlewares.InitAdminAuthMiddleware(c * gin.Context))
	*/

	//导入路由组
	routers.AdminRouterInit(r)
	routers.ManagerRouterInit(r)
	routers.FocusRouterInit(r)
	routers.RoleRouterInit(r)
	routers.AccessRouterInit(r)

	ginPort := config.Section("app").Key("port").String()
	r.Run(ginPort)
}
