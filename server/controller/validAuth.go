package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"server/common/levellog"
	"server/config"
	"server/controller/args/header"
	"server/dao"
	"server/response"
	resc "server/response/code"
)

// ValidAuthorization 校验authorization
func ValidAuthorization(ctx *gin.Context) {
	authorization := ctx.GetHeader(header.Authorization)
	username := ctx.GetHeader(header.Username)
	token, err := GetToken(authorization)
	if err != nil {
		levellog.Controller("token校验失败")
		response.Failed(ctx, "token校验失败，请检查token是否过期")
		ctx.Abort()
		return
	}
	// 判断token里的和redis中的是否一致
	redisDB := dao.GetRedis()
	redisToken, err := redisDB.Get(context.Background(), "token:"+username).Result()
	if err != nil {
		levellog.Controller("从redis中获取token失败")
		response.Failed(ctx, "从数据库中获取token失败")
		ctx.Abort()
		return
	}
	if token != redisToken {
		levellog.Controller("前端传递token与redis中不一致")
		response.FailedWithCode(ctx, resc.TokenError, "无效的token")
		ctx.Abort()
		return
	}
	levellog.Controller("token校验通过")
	ctx.Next()
}

// ValidAuth 权限校验
func ValidAuth(context *gin.Context) (claims jwt.MapClaims, err error) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err = config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return nil, err
	}
	return claims, nil
}
