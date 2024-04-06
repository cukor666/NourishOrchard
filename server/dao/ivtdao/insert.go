package ivtdao

import (
	"server/dao"
	"server/models"
)

// Insert 添加库存信息
func Insert(inventory models.Inventory) error {
	db := dao.GetMySQL()
	return db.Model(&inventory).Create(&inventory).Error
}
