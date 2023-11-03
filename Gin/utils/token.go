package utils

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// UserClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
type UserClaims struct {
	// 可根据需要自行添加字段
	Name                 string `json:"name"`
	Password             string `json:"password"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// 用于签名的字符串
var mySigningKey = []byte("cukor.cn")

// GenToken 创建jwt
func GenToken(name, password string) (string, error) {
	// 创建 Claims
	claims := UserClaims{
		name,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 设置过期时间，24小时后过期
			Issuer:    "cukor",
		}, //签发人
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成前面字符串
	return token.SignedString(mySigningKey)
}

// ParseToken 解析jwt
func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		// token, err := jwt.Parse(tokenString, func(t * jwt.Token) (interface{}, error) {...})
		return mySigningKey, nil
	})
	if err != nil {
		log.Printf("parse failed, err: %v\n", err)
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式：1. 放在请求头 2. 放在请求体 3. 放在URI
		// 这里将Token放到请求头上，Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "请求头中的auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割，预期[0]是Bearer，[1]是token内容，没有其他的了
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "请求头中的auth格式错误",
			})
			c.Abort()
			return
		}
		// parts[1]是token内容，解析token
		uc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的用户名和密码信息保存到请求的上下文c上
		c.Set("name", uc.Name)
		c.Set("password", uc.Password)
		c.Next() // 后续的处理函数可以用过c.Get("name")来获取当前请求的用户信息
	}
}
