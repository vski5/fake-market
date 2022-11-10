package models

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v9" //连接Redis 7（对应go-redis/v9）
	"gopkg.in/ini.v1"              //注意引用的是指定版本号V1 "gopkg.in/ini.v1"
)

var redisCoretxt = context.Background()
var RedisDb *redis.Client

// 创建 redis 链接
func init() {
	//读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	redisIpPort := config.Section("redis").Key("ip").String()
	redisPassword := config.Section("redis").Key("password").String()

	//连接Redis 7（对应go-redis/v9）
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     redisIpPort,
		Password: redisPassword, // set password
		DB:       0,             // use default DB
	})

	dontNeedString, err := RedisDb.Ping(redisCoretxt).Result()
	if err != nil {
		fmt.Println(err, dontNeedString)
	}

}
