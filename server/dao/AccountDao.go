package dao

import (
	"log"
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
		log.Println("数据添加到account表失败")
		tx.Rollback()
		return 0, false
	}
	// 将数据插入到user表中
	err = tx.Create(&user).Error
	if err != nil {
		log.Println("数据添加到user表失败")
		tx.Rollback()
		return 0, false
	}
	if err = tx.Commit().Error; err != nil {
		log.Fatal(err)
		return 0, false
	}
	return user.ID, true
}

func (a AccountDao) GetCountByUsername(username string, promise int) (cnt int64) {
	mysqlDB.Table(models.Account{}.TableName()).
		Where("username = ? AND promise = ?", username, promise).Count(&cnt)
	return
}

// Get 根据账号获取账号信息
func (a AccountDao) Get(username string, promise int) (result models.Account, err error) {
	err = mysqlDB.Table(models.Account{}.TableName()).
		Where("username = ? AND promise = ?", username, promise).
		First(&result).Error
	return
}
