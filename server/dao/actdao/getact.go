package actdao

import (
	"server/dao"
	"server/models"
)

// GetCountByUsername 检验账号存不存在
func GetCountByUsername(username string, promise int) (cnt int64) {
	dao.GetMySQL().Model(&models.Account{}).Where("username = ? AND promise = ?", username, promise).Count(&cnt)
	return
}
