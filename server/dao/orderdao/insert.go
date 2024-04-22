package orderdao

import (
	"server/dao"
	"server/models"
)

// Insert 添加订单信息
func Insert(order models.Order) error {
	order.Status = 1 // 订单状态为待支付
	db := dao.GetMySQL()
	return db.Create(&order).Error
}
