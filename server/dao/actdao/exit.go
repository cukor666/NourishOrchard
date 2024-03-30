package actdao

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"server/common/levellog"
	"server/dao"
)

// Exit 退出登录，删除redis中的token
func Exit(username string) (err error) {
	var (
		tokenKey   = "token:" + username
		accountKey = []string{
			"account:" + username + ":id",
			"account:" + username + ":username",
			"account:" + username + ":promise",
			"account:" + username + ":password",
		}
		ctx = context.Background()
	)

	rd := dao.GetRedis()

	err = rd.Del(ctx, tokenKey).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			levellog.Dao(fmt.Sprintf("键%s不存在", tokenKey))
			return err
		} else {
			levellog.Dao(fmt.Sprintf("redis出错, err = %v", err))
			return err
		}
	}
	err = rd.Del(ctx, accountKey...).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			levellog.Dao(fmt.Sprintf("键%s不存在", accountKey))
			return err
		} else {
			levellog.Dao(fmt.Sprintf("redis出错，err = %v", err))
			return err
		}
	}
	return err
}
