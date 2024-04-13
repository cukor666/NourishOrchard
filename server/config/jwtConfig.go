package config

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"server/common/levellog"
	"server/models"
	"server/utils/promisetool"
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
		"promise":  promisetool.ToString(account.Promise),
		"exp":      expirationTime,
	}

	// 创建Token并签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.JWTConfig.SecretKey))
}

// parseToken 解析Token
func parseToken(token *jwt.Token) (interface{}, error) {
	// 验证签名方法
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(conf.JWTConfig.SecretKey), nil
}

// ParseAndVerifyJWT 解析和验证JWT
func ParseAndVerifyJWT(tokenString string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	// 解析Token并验证签名
	if token, err = jwt.Parse(tokenString, parseToken); err != nil {
		levellog.Config(fmt.Sprintf("%v", err))
		return nil, err
	}
	// 验证发行人
	claims, ok := token.Claims.(jwt.MapClaims)
	// 验证Token是否有效
	if !ok || !token.Valid {
		err = errors.New("JWT claims验证失败")
		levellog.Config(fmt.Sprintf("JWT验证失败: %v", err))
		return nil, err
	}
	// 验证发行人
	if claims["iss"] != conf.JWTConfig.Issuer {
		levellog.Config("Invalid issuer")
		err = errors.New("invalid issuer")
		return nil, err
	}
	levellog.Config(fmt.Sprintf("发行人验证成功: %v", claims["iss"]))
	return claims, nil
}
