package actdao

import (
	"server/dao"
	"server/models"
)

// Get 根据账号获取账号信息
func Get(username string, promise int) (result models.Account, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.Account{}).
		Where("username = ? AND promise = ?", username, promise).
		First(&result).Error
	return
}
