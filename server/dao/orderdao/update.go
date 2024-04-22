package orderdao

import (
	"server/dao"
	"server/models"
)

// Update 更新订单信息
func Update(order models.Order) error {
	db := dao.GetMySQL()
	return db.Model(&models.Order{}).Where("id =?", order.ID).Updates(order).Error
}
