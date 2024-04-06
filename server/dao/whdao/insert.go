package whdao

import (
	"server/dao"
	"server/models"
)

// Insert 添加仓库信息
func Insert(warehouse models.Warehouse) error {
	db := dao.GetMySQL()
	err := db.Model(&warehouse).Create(&warehouse).Error
	return err
}
