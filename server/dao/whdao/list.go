package whdao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// List 查询仓库列表，分页查询
func List(p simpletool.Page, w models.Warehouse) (results []models.Warehouse, total int64, err error) {
	addr := w.Address
	w.Address = ""
	db := dao.GetMySQL()
	tx := db.Model(&w).Where(&w).
		Where("address LIKE ?", fmt.Sprintf("%%%s%%", addr)).
		Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&results).Error
	if err != nil {
		levellog.Dao("查询失败")
		return nil, 0, err
	}
	return
}
