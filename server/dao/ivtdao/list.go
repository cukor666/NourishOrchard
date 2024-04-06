package ivtdao

import (
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// List 查询库存列表，分页查询
func List(p simpletool.Page, inventory models.Inventory) (result []models.Inventory, total int64, err error) {
	db := dao.GetMySQL()
	tx := db.Model(&inventory).Where(&inventory).Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	return
}
