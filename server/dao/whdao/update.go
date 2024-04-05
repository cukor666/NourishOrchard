package whdao

import (
	"server/dao"
	"server/models"
)

// Update 更新仓库信息
func Update(warehouse models.Warehouse) error {
	db := dao.GetMySQL()
	return db.Model(&warehouse).Omit("id").Updates(&warehouse).Error
}
