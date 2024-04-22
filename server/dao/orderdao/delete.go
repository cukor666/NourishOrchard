package orderdao

import (
	"server/dao"
	"server/models"
)

// Delete 删除订单
func Delete(id int64) error {
	db := dao.GetMySQL()
	return db.Model(&models.Order{}).Where("id = ?", id).Delete(&models.Order{}).Error
}
