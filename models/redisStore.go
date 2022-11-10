package models

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}
type CookieRedisStore struct {
}

// 实现设置captcha的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := RedisDb.Set(ctx, key, value, time.Minute*2).Err()

	return err
}

// 实现获取captcha的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}

/*Cookie*/
// 设置值
// 全局共用连接var RedisDb *redis.Client 写models.RedisDb
func (r CookieRedisStore) Set(key, value string, c *gin.Context) error {

	// 设置 Session
	err := RedisDb.Set(ctx, key, value, time.Minute*2).Err()
	return err

}

// 获取值
func (r CookieRedisStore) Get(key string, c *gin.Context) string {

	value, err := RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value

}
