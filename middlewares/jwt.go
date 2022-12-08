package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Jwt() {
	// 定义一个字符串变量，用于存储签名密钥。
	var secret = "your-secret-key"

	// 创建一个新的JWT。
	token := jwt.New(jwt.SigningMethodHS256)

	// 为该JWT设置一些声明。
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = 123
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

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
