package models

import (
	"context"
	"fmt"
	"os"

	"github.com/rueian/rueidis"
	"gopkg.in/ini.v1" //注意引用的是指定版本号V1 "gopkg.in/ini.v1"
)

var Rueidisctx context.Context
var RedisStore rueidis.Client

func RedisInit() {
	//读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	redisIpPort := config.Section("redis").Key("ip").String()
	redisPassword := config.Section("redis").Key("password").String()

	//连接redis
	var err2 error
	RedisStore, err2 = rueidis.NewClient(rueidis.ClientOption{
		//写链接
		InitAddress: []string{redisIpPort},
		//写密码
		Password: redisPassword,
	})

	//检错
	if err2 != nil {
		panic(err2)
	}
	//结束时的关闭
	//defer redisStore.Close()

	//redis共享的上下文
	Rueidisctx = context.Background()

	// SET key val NX，设置key/value
	//redisStore.Do(rueidisctx, redisStore.B().Set().Key("key").Value("val").Nx().Build()).Error()
	// GET key
	//redisStore.Do(rueidisctx, redisStore.B().Get().Key("key").Build()).ToString()

}
