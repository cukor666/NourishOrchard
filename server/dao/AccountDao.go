package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"server/models"
)

// Insert
// / 添加账户
func (a AccountDao) Insert(account models.Account, user models.User) (uint, bool) {
	// 开启事务
	tx := mysqlDB.Begin()
	// 将数据插入到account表中
	err := tx.Create(&account).Error
	if err != nil {
		levelLog("数据添加到account表失败")
		tx.Rollback()
		return 0, false
	}
	// 将数据插入到user表中
	err = tx.Create(&user).Error
	if err != nil {
		levelLog("数据添加到user表失败")
		tx.Rollback()
		return 0, false
	}
	if err = tx.Commit().Error; err != nil {
		levelLog(fmt.Sprintf("err: %v", err))
		return 0, false
	}
	return user.ID, true
}

// GetCountByUsername 检验账号存不存在
func (a AccountDao) GetCountByUsername(username string, promise int) (cnt int64) {
	err := mysqlDB.Model(&models.Account{}).Where("username = ? AND promise = ?", username, promise).Count(&cnt).Error
	if err != nil {
		levelLog("获取数量失败")
		return 0
	}
	return
}

// Get 根据账号获取账号信息
func (a AccountDao) Get(username string, promise int) (result models.Account, err error) {
	err = mysqlDB.Model(&models.Account{}).
		Where("username = ? AND promise = ?", username, promise).
		First(&result).Error
	return
}

// Exit 退出登录，删除redis中的token
func (a AccountDao) Exit(username string) (err error) {
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
			levelLog(fmt.Sprintf("键%s不存在", tokenKey))
			return err
		} else {
			levelLog(fmt.Sprintf("redis出错, err = %v", err))
			return err
		}
	}
	err = redisDB.Del(ctx, accountKey...).Err()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			levelLog(fmt.Sprintf("键%s不存在", accountKey))
			return err
		} else {
			levelLog(fmt.Sprintf("redis出错，err = %v", err))
			return err
		}
	}
	return err
}

// GetPassword 根据账号获取数据库中的密码
func (a AccountDao) GetPassword(username string) (password string, err error) {
	err = mysqlDB.Model(&models.Account{}).Select("password").
		Where("username = ?", username).Take(&password).Error
	if err != nil {
		levelLog(fmt.Sprintf("获取账户%s密码失败", username))
		return "", err
	}
	return password, nil
}

// ChangePassword 修改密码
func (a AccountDao) ChangePassword(username, password string) (err error) {
	err = mysqlDB.Model(&models.Account{}).Where("username = ?", username).Update("password", password).Error
	if err != nil {
		levelLog(fmt.Sprintf("用户%s更新密码失败", username))
		return err
	}
	return nil
}
