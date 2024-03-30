package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
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
