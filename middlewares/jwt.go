package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9" //连接Redis 7（对应go-redis/v9）
	jwt "github.com/golang-jwt/jwt/v4"
)

// 为了演示，我们定义一个简单的用户信息结构体
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"login_name"`
}

// 实现jwt.Claims这个接口里的所有方法，将存储JWT的地方变为Redis
func StoreJWTInRedis(client *redis.Client, token jwt.Token) error {
	// Convert the claims to JSON
	claimsJSON, err := json.Marshal(token.Claims)
	if err != nil {
		return err
	}

	// Store the JSON in Redis
	err = client.Set(token.Raw, claimsJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// 设置JWT
func CreateJWT(c *gin.Context, claims jwt.Claims) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("mysecret"))
	if err != nil {
		http.Error(c.Writer, "Error creating JWT", http.StatusInternalServerError)
		return
	}
	c.Writer.Header().Set("Authorization", "Bearer "+tokenString)
}

// 验证JWT
func VerifyJWT(r *http.Request) (jwt.Claims, error) {
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return nil, errors.New("Missing JWT")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte("mysecret"), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid JWT")
}
