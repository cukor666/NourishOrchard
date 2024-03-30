package fruitdao

import (
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Update 更新水果信息
func Update(fruit models.Fruit) error {
	db := dao.GetMySQL()
	err := db.Model(&models.Fruit{}).Where("id = ?", fruit.ID).Omit("id").Updates(&fruit).Error
	if err != nil {
		levellog.Dao("更新失败")
		return err
	}
	return nil
}
