package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/common/levellog"
	"server/config"
	"server/controller/args/header"
	"server/dao"
	"server/response"
)

// ValidAuthorization 校验authorization
func ValidAuthorization(ctx *gin.Context) {
	// 获取请求中的authorization和username
	authorization := ctx.GetHeader(header.Authorization)
	username := ctx.GetHeader(header.Username)

	// 校验token的有效性
	if err := validateToken(authorization, username); err != nil {
		levellog.Controller(fmt.Sprintf("token校验失败: %s", err))
		response.Failed(ctx, err.Error())
		ctx.Abort()
		return
	}

	// 记录日志并继续处理后续请求
	levellog.Controller("token校验通过")
	ctx.Next()
}

// validateToken 校验token的有效性
func validateToken(authorization, username string) error {
	var (
		token      string
		err        error
		redisToken string
	)

	if token, err = GetToken(authorization); err != nil {
		levellog.Controller(fmt.Sprintf("token校验失败，请检查token是否过期: %s", err))
		return err
	}

	// 从redis中获取username对应的token
	if redisToken, err = getRedisToken(username); err != nil {
		levellog.Controller(fmt.Sprintf("从redis中获取token失败: %s", err))
		return err
	}

	// 判断token是否与redis中存储的token一致
	if token != redisToken {
		levellog.Controller("前端传递token与redis中不一致")
		return errors.New("无效的token")
	}

	levellog.Controller("token校验通过")
	return nil
}

// getRedisToken 从Redis中获取token
func getRedisToken(username string) (redisToken string, err error) {
	return dao.GetRedis().Get(context.Background(), "token:"+username).Result()
}

// ValidAuth 有效权限验证函数
func ValidAuth(context *gin.Context) (claims jwt.MapClaims, err error) {
	var token string
	// 解析token
	if token, err = GetToken(context.GetHeader(header.Authorization)); err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	return config.ParseAndVerifyJWT(token)
}
