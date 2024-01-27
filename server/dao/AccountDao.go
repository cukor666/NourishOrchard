package dao

import (
	"server/models"
)

func (a AccountDao) Insert(account models.Account) bool {
	return mysqlDB.Create(&account).Error == nil
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
