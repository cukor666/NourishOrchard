package config

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTConfig struct {
	SecretKey string `yaml:"secretKey"`
	Issuer    string `yaml:"issuer"`
}

// MyClaims 定义了自定义的JWT声明
type MyClaims struct {
	AccountUsername string `json:"accountUsername"`
	jwt.StandardClaims
}

// GenerateJWT 生成JWT
func GenerateJWT(accountUsername string) (string, error) {
	// 设置过期时间为1小时
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建CustomClaims结构体实例
	claims := MyClaims{
		AccountUsername: accountUsername,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			Issuer:    conf.JWTConfig.Issuer, // 签发人
		},
	}

	// 创建Token并签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(conf.JWTConfig.SecretKey))
}

// ParseAndVerifyJWT 解析和验证JWT
func ParseAndVerifyJWT(tokenString string) (*MyClaims, error) {
	// 解析Token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return conf.JWTConfig.SecretKey, nil
		})

	// 验证Token的有效性
	if err != nil {
		return nil, err
	}

	// 获取Claims
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
