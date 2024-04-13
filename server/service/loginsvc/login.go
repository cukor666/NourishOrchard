package loginsvc

import (
	"context"
	"errors"
	"fmt"
	"server/common/levellog"
	"server/config"
	"server/dao"
	"server/dao/actdao"
	"server/dao/admindao"
	"server/dao/empdao"
	"server/dao/userdao"
	"server/models"
	mc "server/models/code"
	"server/utils/pwdtool"
	"strconv"
	"time"
)

func Login(account models.Account) (token string, err error) {
	// 首先校验账号是否存在
	//if actdao.GetCountByUsername(account.Username, account.Promise) == 0 {
	//	return "", errors.New("账号不存在")
	//}
	// 如果该用户存在，则继续校验密码是否正确
	// 从数据库中获取该账号
	dbAccount, err := actdao.Get(account.Username, account.Promise)
	if err != nil {
		return "", errors.New("获取账号信息失败")
	}
	// 对密码进行处理
	if !pwdtool.PwdOK(dbAccount.Password, account.Password) {
		levellog.Service("密码错误")
		return "", errors.New("密码错误")
	}
	uid := uint(0)
	switch account.Promise {
	case mc.USER:
		uid, err = userdao.GetUId(dbAccount.Username)
	case mc.EMPLOYEE:
		uid, err = empdao.GetUId(dbAccount.Username)
	case mc.ADMIN:
		uid, err = admindao.GetUId(dbAccount.Username)
	default:
		levellog.Service("权限不正确")
	}
	if err != nil {
		levellog.Service("获取ID失败")
		return "", errors.New("获取ID失败")
	}
	baseKey := "account:" + account.Username
	redisCmd := map[string]string{
		baseKey + ":id":       strconv.Itoa(int(uid)),
		baseKey + ":username": account.Username,
		baseKey + ":password": account.Password,
		baseKey + ":promise":  strconv.Itoa(account.Promise),
	}
	// 将account信息存入redis，然后将redis的键存入JWT然后生成token返回给前端
	ctx := context.Background()
	redisDB := dao.GetRedis()
	for k, v := range redisCmd {
		if err = redisDB.Set(ctx, k, v, 24*time.Hour).Err(); err != nil {
			levellog.Service(fmt.Sprintf("存储到redis失败， err：%v", err))
		}
	}

	// 生成JWT
	jwtToken, err := config.GenerateJWT(account)
	if err != nil {
		levellog.Service("生成JWT失败")
		return "", errors.New("生成JWT失败")
	}
	// 将token保存到redis中
	redisDB.Set(ctx, "token:"+account.Username, jwtToken, 7*24*time.Hour)

	return jwtToken, nil
}
