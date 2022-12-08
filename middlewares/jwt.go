package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtBody struct {
	secret string
	userId int
}

func (JwtBody JwtBody) Jwt() {
	// 存储签名密钥。
	secret := JwtBody.secret

	// 创建一个新的JWT。
	token := jwt.New(jwt.SigningMethodHS256)

	// 为该JWT设置一些声明。
	claims := token.Claims.(jwt.MapClaims) //将JWT的声明转换为一个jwt.MapClaims类型的变量。这样 就可以使用字典语法来设置声明的值，
	claims["userId"] = JwtBody.userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //设置JWT的过期时间,默认是24H

	// 使用签名密钥对JWT进行签名。
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JWT:", tokenString)

	// 使用签名密钥和JWT验证签名。
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		claims := result.Claims.(jwt.MapClaims)
		fmt.Println("userId:", claims["userId"])
		fmt.Println("Expiration Time:", claims["exp"])
	}
}
