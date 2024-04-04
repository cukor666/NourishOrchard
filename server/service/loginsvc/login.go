package loginsvc

import (
	"context"
	"fmt"
	"server/common"
	"server/common/levellog"
	"server/config"
	"server/dao"
	"server/dao/actdao"
	mc "server/models/code"
	"server/request"
	resc "server/response/code"
	"server/service/actsvc"
	"server/utils/promisetool"
	"server/utils/pwdtool"
	"strconv"
	"time"
)

func Login(req request.LoginRequest) (string, error) {
	promise := promisetool.ToInt(req.Promise)
	// 通过账号名找到对应的账号信息，如果没有该用户名则直接返回false
	exists := actsvc.IsExists(req.Username, promise)
	if !exists {
		return "", &common.MyError{
			Code: resc.NotFoundAccount,
			Msg:  "账号不存在, 参数格式不正确: " + req.Username,
		}
	}
	// 如果该用户存在，则继续校验密码是否正确
	// 从数据库中获取该账号
	account, err := actdao.Get(req.Username, promise)
	if err != nil {
		return "", &common.MyError{
			Code: resc.SystemError,
			Msg:  "系统错误",
		}
	}
	// 对密码进行处理
	if !pwdtool.PwdOK(account.Password, req.Password) {
		levellog.Service("密码错误")
		return "", &common.MyError{
			Code: resc.PasswordFailed,
			Msg:  "密码错误",
		}
	}
	uid := uint(0)
	switch promise {
	case mc.USER:
		uid, err = userLogin(account.Username)
	case mc.EMPLOYEE:
		uid, err = employeeLogin(account.Username)
	case mc.ADMIN:
		uid, err = adminLogin(account.Username)
	default:
		levellog.Service("没有定义")
	}
	if err != nil {
		levellog.Service("获取ID失败")
		return "", &common.MyError{
			Code: resc.SystemError,
			Msg:  "账号" + account.Username + "没有对应的ID, promise: " + strconv.Itoa(promise),
		}
	}
	baseKey := "account:" + account.Username
	redisCmd := map[string]string{
		baseKey + ":id":       strconv.Itoa(int(uid)),
		baseKey + ":username": account.Username,
		baseKey + ":password": account.Password,
		baseKey + ":promise":  strconv.Itoa(promise),
	}
	// 将account信息存入redis，然后将redis的键存入JWT然后生成token返回给前端
	ctx := context.Background()
	redisDB := dao.GetRedis()
	for k, v := range redisCmd {
		if err := redisDB.Set(ctx, k, v, 24*time.Hour).Err(); err != nil {
			levellog.Service(fmt.Sprintf("存储到redis失败， err：%v", err))
		}
	}

	// 生成JWT
	jwtToken, err := config.GenerateJWT(account)
	if err != nil {
		levellog.Service("生成JWT失败")
		return "", &common.MyError{
			Code: resc.SystemError,
			Msg:  "系统错误无法生成JWT",
		}
	}
	// 将token保存到redis中
	redisDB.Set(ctx, "token:"+account.Username, jwtToken, 7*24*time.Hour)

	return jwtToken, nil
}