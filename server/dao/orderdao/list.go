package orderdao

import (
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// List 查询订单列表
func List(p simpletool.Page, order models.Order) (result []models.Order, total int64, err error) {
	db := dao.GetMySQL()
	title, receiverName, address, remark := order.SetZero()
	tx := db.Model(&models.Order{}).Where(&order).
		Where("title LIKE ?", "%"+title+"%").
		Where("receiver_name LIKE ?", "%"+receiverName+"%").
		Where("address LIKE ?", "%"+address+"%").
		Where("remark LIKE ?", "%"+remark+"%").
		Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	return
}
