package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"server/models"
	"server/utils"
	"time"
)

type JWTConfig struct {
	SecretKey string `yaml:"secretKey"`
	Issuer    string `yaml:"issuer"`
}

// GenerateJWT 生成JWT
func GenerateJWT(account models.Account) (string, error) {
	// 设置过期时间为1小时
	expirationTime := time.Now().Add(7 * 24 * time.Hour).Unix()
	// 设置claims
	claims := jwt.MapClaims{
		"iss":      conf.JWTConfig.Issuer,
		"username": account.Username,
		"promise":  utils.PromiseToString(account.Promise),
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
		levelLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if claims["iss"] != conf.JWTConfig.Issuer {
			levelLog("Invalid issuer")
			return nil, err
		}
		levelLog(fmt.Sprintf("发行人验证成功: %v", claims["iss"]))
	} else {
		levelLog(fmt.Sprintf("%v", err))
		return nil, err
	}
	return claims, nil
}
