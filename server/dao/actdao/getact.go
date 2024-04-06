package actdao

import (
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// GetCountByUsername 检验账号存不存在
func GetCountByUsername(username string, promise int) (cnt int64) {
	db := dao.GetMySQL()
	err := db.Model(&models.Account{}).Where("username = ? AND promise = ?", username, promise).Count(&cnt).Error
	if err != nil {
		levellog.Dao("获取数量失败")
		return 0
	}
	return
}
