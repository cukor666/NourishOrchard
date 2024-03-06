package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/common/statucode"
	"server/dao"
	"server/response"
	"server/service"
	"strings"
)

type AccountController struct{}
type LoginController struct{}
type RegisterController struct{}

var (
	loginService    service.LoginService
	registerService service.RegisterService
	userService     service.UserService
)

// ValidAuthorization 校验authorization
func ValidAuthorization(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	username := ctx.GetHeader("username")
	token, err := GetToken(authorization)
	if err != nil {
		levelLog("token校验失败")
		response.Failed(ctx, "token校验失败，请检查token是否过期")
		return
	}
	// 判断token里的和redis中的是否一致
	levelLog("token: " + token)
	redisDB := dao.GetRedisDB()
	redisToken, err := redisDB.Get(context.Background(), "token:"+username).Result()
	if err != nil {
		levelLog("从redis中获取token失败")
		response.Failed(ctx, "从数据库中获取token失败")
		return
	}
	if token != redisToken {
		levelLog("token已过期")
		response.FailedWithCode(ctx, statucode.TOKENERROR, "token已过期")
		return
	}
	levelLog("token校验通过")
	ctx.Next()
}

// GetToken 从authorization中获取token
func GetToken(authorization string) (token string, err error) {
	// 获取token
	split := strings.Split(authorization, " ")
	if len(split) != 2 {
		levelLog("获取token失败")
		return "", errors.New(fmt.Sprintf("获取token失败"))
	}
	return split[1], nil
}

func levelLog(w string) {
	log.Println("controller层：", w)
}
