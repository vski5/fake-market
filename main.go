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
	//读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	//导入路由组
	routers.AdminRouterInit(r)

	ginPort := config.Section("app").Key("port").String()
	r.Run(ginPort)
}
