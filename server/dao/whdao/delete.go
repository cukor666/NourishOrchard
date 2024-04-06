package whdao

import (
	"server/dao"
	"server/models"
)

// Delete 根据id删除仓库信息
func Delete(id int64) (w models.Warehouse, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.Warehouse{}).Where("id = ?", id).Delete(&w).Error
	if err != nil {
		return models.Warehouse{}, err
	}
	return
}
