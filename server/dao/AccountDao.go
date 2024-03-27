package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"server/common/levellog"
	"server/models"
)

// Insert
// / 添加账户
func (ad *AccountDao) Insert(account models.Account, user models.User) (uint, bool) {
	// 开启事务
	tx := mysqlDB.Begin()
	// 将数据插入到account表中
	err := tx.Create(&account).Error
	if err != nil {
		levellog.Dao("数据添加到account表失败")
		tx.Rollback()
		return 0, false
	}
	// 将数据插入到user表中
	err = tx.Create(&user).Error
	if err != nil {
		levellog.Dao("数据添加到user表失败")
		tx.Rollback()
		return 0, false
	}
	if err = tx.Commit().Error; err != nil {
		levellog.Dao(fmt.Sprintf("err: %v", err))
		return 0, false
	}
	return user.ID, true
}

// GetCountByUsername 检验账号存不存在
func (ad *AccountDao) GetCountByUsername(username string, promise int) (cnt int64) {
	err := mysqlDB.Model(&models.Account{}).Where("username = ? AND promise = ?", username, promise).Count(&cnt).Error
	if err != nil {
		levellog.Dao("获取数量失败")
		return 0
	}
	return
}

// Get 根据账号获取账号信息
func (ad *AccountDao) Get(username string, promise int) (result models.Account, err error) {
	err = mysqlDB.Model(&models.Account{}).
		Where("username = ? AND promise = ?", username, promise).
		First(&result).Error
	return
}

// Exit 退出登录，删除redis中的token
func (ad *AccountDao) Exit(username string) (err error) {
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

	err = redisDB.Del(ctx, tokenKey).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			levellog.Dao(fmt.Sprintf("键%s不存在", tokenKey))
			return err
		} else {
			levellog.Dao(fmt.Sprintf("redis出错, err = %v", err))
			return err
		}
	}
	err = redisDB.Del(ctx, accountKey...).Err()
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
