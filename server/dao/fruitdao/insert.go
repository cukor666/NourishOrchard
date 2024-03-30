package fruitdao

import (
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Insert 插入水果信息
func Insert(fruit models.Fruit) error {
	db := dao.GetMySQL()
	err := db.Model(&models.Fruit{}).Create(&fruit).Error
	if err != nil {
		levellog.Dao("插入水果信息失败")
		return err
	}
	return nil
}
