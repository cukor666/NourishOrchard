package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"server/common"
	"server/common/statucode"
	"server/dao"
	"strings"
)

func PromiseToString(promise int) string {
	switch promise {
	case common.USER:
		return "user"
	case common.EMPLOYEE:
		return "employee"
	case common.ADMIN:
		return "admin"
	default:
		return "myerr"
	}
}

func PromiseToInt(promise string) int {
	promise = strings.ToLower(promise)
	switch promise {
	case "user":
		return common.USER
	case "employee":
		return common.EMPLOYEE
	case "admin":
		return common.ADMIN
	default:
		return 0
	}
}

//func MD5(v string) string {
//	d := []byte(v)
//	m := md5.New()
//	m.Write(d)
//	return hex.EncodeToString(m.Sum(nil))
//}

// GetPwd 给密码加密
func GetPwd(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func PwdOK(origin, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(origin), []byte(password)) == nil
}

// ValidAuthorization 校验authorization
func ValidAuthorization(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	username := ctx.GetHeader("username")
	token, err := GetToken(authorization)
	if err != nil {
		ctx.AbortWithStatusJSON(statucode.NOTFOUNDTOKEN, gin.H{
			"code": statucode.NOTFOUNDTOKEN,
			"msg":  "获取token失败",
			"data": 0,
		})
		return
	}
	// 判断token里的和redis中的是否一致
	log.Println(token)
	redisDB := dao.GetRedisDB()
	redisToken, err := redisDB.Get(context.Background(), "token:"+username).Result()
	if err != nil {
		log.Println("从redis中获取token失败")
		ctx.AbortWithStatusJSON(statucode.SYSTEMERR, gin.H{
			"code": statucode.SYSTEMERR,
			"msg":  "从redis中获取token失败",
			"data": 0,
		})
		return
	}
	if token != redisToken {
		ctx.AbortWithStatusJSON(statucode.TOKENERROR, gin.H{
			"code": statucode.TOKENERROR,
			"msg":  "前端传递的token与后端数据库不一致",
			"data": 0,
		})
		return
	}
	log.Println("token校验通过")
	ctx.Next()
}

// GetToken 从authorization中获取token
func GetToken(authorization string) (token string, err error) {
	// 获取token
	split := strings.Split(authorization, " ")
	if len(split) != 2 {
		log.Println("获取token失败")
		return "", errors.New(fmt.Sprintf("获取token失败"))
	}
	return split[1], nil
}
