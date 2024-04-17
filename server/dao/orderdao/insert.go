package orderdao

import (
	"server/dao"
	"server/models"
)

// Insert 添加订单信息
func Insert(order models.Order) error {
	db := dao.GetMySQL()
	return db.Model(&order).Create(&order).Error
}
