package ivtdao

import (
	"server/dao"
	"server/models"
)

// Delete 根据id删除库存信息
func Delete(id int64) error {
	db := dao.GetMySQL()
	return db.Delete(&models.Inventory{}, id).Error
}
