package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"server/common"
	"server/models"
	"time"
)

type JWTConfig struct {
	SecretKey string `yaml:"secretKey"`
	Issuer    string `yaml:"issuer"`
}

// 权限整数值转为字面量字符串
func promiseToString(promise int) string {
	switch promise {
	case common.USER:
		return "user"
	case common.EMPLOYEE:
		return "employee"
	case common.ADMIN:
		return "admin"
	}
	return ""
}

// GenerateJWT 生成JWT
func GenerateJWT(account models.Account) (string, error) {
	// 设置过期时间为1小时
	expirationTime := time.Now().Add(1 * time.Hour).Unix()
	// 设置claims
	claims := jwt.MapClaims{
		"iss":      conf.JWTConfig.Issuer,
		"username": account.Username,
		"promise":  promiseToString(account.Promise),
		"exp":      expirationTime,
	}

	// 创建Token并签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.JWTConfig.SecretKey))
}

// ParseAndVerifyJWT 解析和验证JWT
func ParseAndVerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.JWTConfig.SecretKey), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if claims["iss"] != conf.JWTConfig.Issuer {
			log.Fatal("Invalid issuer")
		}
		log.Println("Issuer verified: ", claims["iss"])
	} else {
		log.Println(err)
		return nil, err
	}
	return claims, nil
}
